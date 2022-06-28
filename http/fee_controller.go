/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package http

import (
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/fee"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type FeeController struct {
	web.Controller
}

var (
	riskyCoinRankThreshold int
	riskyCoinRisingRate    *big.Float
	proxyFeeRatioMap       map[uint64]int64
)

var dstLockProxyMap = make(map[string]string, 0)

func (c *FeeController) GetFee() {
	var getFeeReq models.GetFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &getFeeReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	token := new(models.Token)
	res := db.Where("hash = ? and chain_id = ?", getFeeReq.Hash, getFeeReq.SrcChainId).Preload("TokenBasic").First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have token: %s", getFeeReq.SrcChainId, getFeeReq.Hash))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	feeTokenPrecison := token.Precision
	if token.TokenBasic.Price == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("token: %v price is 0", token.TokenBasic.Name))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	chainFee := new(models.ChainFee)
	res = db.Where("chain_id = ?", getFeeReq.DstChainId).Preload("TokenBasic").First(chainFee)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.DstChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	chainFeeToken := new(models.Token)
	res = db.Where("chain_id = ? and token_basic_name = ?", chainFee.ChainId, chainFee.TokenBasicName).
		First(chainFeeToken)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.DstChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	//check if rank of src token is risky, if so, change the proxyFee value
	proxyFee := new(big.Float).SetInt(&chainFee.ProxyFee.Int)
	if token.TokenBasic.Rank > riskyCoinRankThreshold {
		proxyFeeRatio := proxyFeeRatioMap[getFeeReq.SrcChainId]
		proxyFee.Quo(proxyFee, big.NewFloat(float64(proxyFeeRatio)))
		proxyFee.Mul(proxyFee, riskyCoinRisingRate)
	}
	//check if any coin marked as dying in redis
	if exists, _ := cacheRedis.Redis.Exists(cacheRedis.MarkTokenAsDying + token.TokenBasicName); exists {
		logs.Info("this token is dying", token.TokenBasicName)
		if val, err := cacheRedis.Redis.Get(cacheRedis.MarkTokenAsDying + token.TokenBasicName); err == nil {
			proxyFeeRatio := proxyFeeRatioMap[getFeeReq.SrcChainId]
			proxyFee.Quo(proxyFee, big.NewFloat(float64(proxyFeeRatio)))
			manualRatio, ok := big.NewFloat(0.0).SetString(val)
			if ok {
				proxyFee.Mul(proxyFee, manualRatio)
			} else {
				logs.Error("get dying token manualRatio fail, tokenbasicname: %s", token.TokenBasicName)
			}
		}
	}
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.FEE_PRECISION))
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFeeToken.Precision))))
	usdtFee := new(big.Float).Mul(proxyFee, new(big.Float).SetInt64(chainFee.TokenBasic.Price))
	usdtFee = new(big.Float).Quo(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee := new(big.Float).Mul(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee = new(big.Float).Quo(tokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
	tokenFeeWithPrecision := new(big.Float).Mul(tokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))

	isNative := false
	nativeTokenAmount := new(big.Float).SetInt64(0)
	// get optimistic L1 fee on ethereum
	if basedef.OPTIMISTIC_CROSSCHAIN_ID == getFeeReq.DstChainId {
		ethChainFee := new(models.ChainFee)
		res = db.Where("chain_id = ?", basedef.ETHEREUM_CROSSCHAIN_ID).Preload("TokenBasic").First(ethChainFee)
		if res.RowsAffected == 0 {
			c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", basedef.ETHEREUM_CROSSCHAIN_ID))
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		_, l1UsdtFee, _, err := fee.GetL1Fee(ethChainFee, getFeeReq.DstChainId)
		if err != nil {
			c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("get ethereum L1 fee failed. err=%v", err))
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		l1TokenFee := new(big.Float).Mul(l1UsdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		l1TokenFee = new(big.Float).Quo(l1TokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
		l1TokenFeeWithPrecision := new(big.Float).Mul(l1TokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
		tokenFee = new(big.Float).Add(tokenFee, l1TokenFee)
		tokenFeeWithPrecision = new(big.Float).Add(tokenFeeWithPrecision, l1TokenFeeWithPrecision)
	}

	{
		chainFeeJson, _ := json.Marshal(chainFee)
		logs.Error("chain fee: %s", string(chainFeeJson))
	}
	{
		tokenJson, _ := json.Marshal(token)
		logs.Error("token: %s", string(tokenJson))
	}

	if getFeeReq.SwapTokenHash != "" {
		//check cross native token
		nativeChainFee := new(models.ChainFee)
		res = db.Where("chain_id = ?", getFeeReq.SrcChainId).Preload("TokenBasic").
			First(nativeChainFee)
		if res.RowsAffected == 0 {
			c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.SrcChainId))
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		preloadTokens := make([]*models.Token, 0)
		res = db.Where("token_basic_name = ?", nativeChainFee.TokenBasicName).
			Find(&preloadTokens)
		nativeChainFee.TokenBasic.Tokens = preloadTokens
		if nativeChainFee.TokenBasic != nil && nativeChainFee.TokenBasic.Tokens != nil {
			for _, v := range nativeChainFee.TokenBasic.Tokens {
				if v.ChainId == getFeeReq.SrcChainId && strings.EqualFold(v.Hash, getFeeReq.SwapTokenHash) {
					isNative = true
					nativeFeeAmount := new(big.Float).SetInt(&nativeChainFee.MaxFee.Int)
					nativeFeeAmount = new(big.Float).Quo(nativeFeeAmount, new(big.Float).SetInt64(basedef.FEE_PRECISION))
					nativeFeeAmount = new(big.Float).Quo(nativeFeeAmount, new(big.Float).SetInt64(basedef.Int64FromFigure(int(nativeChainFee.TokenBasic.Precision))))
					if getFeeReq.SrcChainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
						ethChainFee := new(models.ChainFee)
						res = db.Where("chain_id = ?", basedef.ETHEREUM_CROSSCHAIN_ID).Preload("TokenBasic").First(ethChainFee)
						if res.RowsAffected == 0 {
							c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", basedef.ETHEREUM_CROSSCHAIN_ID))
							c.Ctx.ResponseWriter.WriteHeader(400)
							c.ServeJSON()
							return
						}
						_, _, l1FeeAmount, err := fee.GetL1Fee(ethChainFee, getFeeReq.SrcChainId)
						if err != nil {
							c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("get ethereum L1 fee failed. err=%v", err))
							c.Ctx.ResponseWriter.WriteHeader(400)
							c.ServeJSON()
							return
						}
						nativeFeeAmount = new(big.Float).Add(nativeFeeAmount, l1FeeAmount)
					}
					nativeTokenAmount = nativeFeeAmount
					break
				}
			}
		}

		tokenMap := new(models.TokenMap)
		res := db.Where("src_token_hash = ? and src_chain_id = ? and dst_chain_id = ?", getFeeReq.SwapTokenHash, getFeeReq.SrcChainId, getFeeReq.DstChainId).Preload("DstToken").First(tokenMap)
		if res.RowsAffected == 0 {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		if tokenMap.DstChainId != getFeeReq.DstChainId || tokenMap.DstToken == nil {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		tokenBalance, _ := new(big.Int).SetString("100000000000000000000000000000", 10)
		if tokenMap.DstChainId != basedef.PLT_CROSSCHAIN_ID && tokenMap.DstChainId != basedef.BCSPALETTE_CROSSCHAIN_ID && tokenMap.DstChainId != basedef.BCSPALETTE2_CROSSCHAIN_ID {
			tokenBalance, err = cacheRedis.Redis.GetTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash)
			if err != nil {
				ethChains := make(map[uint64]struct{})
				for _, chainId := range basedef.ETH_CHAINS {
					ethChains[chainId] = struct{}{}
				}
				if _, ok := ethChains[tokenMap.DstChainId]; ok {
					var dstLockProxies []string
					for _, cfg := range conf.GlobalConfig.ChainListenConfig {
						if cfg.ChainId == tokenMap.DstChainId {
							dstLockProxies = cfg.ProxyContract
							break
						}
					}
					var dstLockProxy string
					lockProxyKey := fmt.Sprintf("%d-%d-%s", tokenMap.SrcChainId, tokenMap.DstChainId, strings.ToLower(tokenMap.SrcTokenHash))
					dstLockProxy, ok := dstLockProxyMap[lockProxyKey]
					if !ok || len(dstLockProxy) == 0 {
						dstLockProxy, err = common.GetBoundLockProxy(dstLockProxies, tokenMap.SrcTokenHash, tokenMap.DstTokenHash, tokenMap.SrcChainId, tokenMap.DstChainId)
						logs.Info("GetBoundLockProxy srcChain=%d, srcTokenHash=%s, dstTokenHash=%s dstLockProxy=%s, err=%s", tokenMap.SrcChainId, tokenMap.SrcTokenHash, tokenMap.DstTokenHash, dstLockProxy, err)
						if err == nil {
							dstLockProxyMap[lockProxyKey] = dstLockProxy
						}
					}
					logs.Info("lockProxyKey=%s, dstLockProxy=%s", lockProxyKey, dstLockProxy)
					tokenBalance, err = common.GetProxyBalance(tokenMap.DstChainId, tokenMap.DstTokenHash, dstLockProxy)
				} else {
					tokenBalance, err = common.GetBalance(tokenMap.DstChainId, tokenMap.DstTokenHash)
				}

				if err != nil {
					tokenBalance, err = cacheRedis.Redis.GetLongTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash)
					if err != nil {
						c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
							getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
						c.ServeJSON()
						return
					}
				}
				setErr := cacheRedis.Redis.SetTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash, tokenBalance)
				if setErr != nil {
					logs.Error("qweasdredis SetTokenBalance err", setErr)
				}
				setErr1 := cacheRedis.Redis.SetLongTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash, tokenBalance)
				if setErr1 != nil {
					logs.Error("qweasdredis SetLongTokenBalance err", setErr1)
				}
			}
		}
		balance, result := new(big.Float).SetString(tokenBalance.String())
		if !result {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		tokenBalanceWithoutPrecision := new(big.Float).Quo(balance, new(big.Float).SetInt64(basedef.Int64FromFigure(int(tokenMap.DstToken.Precision))))
		c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
			getFeeReq.SwapTokenHash, balance, tokenBalanceWithoutPrecision, isNative, nativeTokenAmount, feeTokenPrecison)
		c.ServeJSON()
	} else {
		c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
			getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
		c.ServeJSON()
	}
}

func (c *FeeController) OldGetFee() {
	var getFeeReq models.GetFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &getFeeReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	token := new(models.Token)
	res := db.Where("hash = ? and chain_id = ?", getFeeReq.Hash, getFeeReq.SrcChainId).Preload("TokenBasic").First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have token: %s", getFeeReq.SrcChainId, getFeeReq.Hash))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	feeTokenPrecison := token.Precision
	chainFee := new(models.ChainFee)
	res = db.Where("chain_id = ?", getFeeReq.DstChainId).Preload("TokenBasic").First(chainFee)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.DstChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	proxyFee := new(big.Float).SetInt(&chainFee.ProxyFee.Int)
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.FEE_PRECISION))
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	usdtFee := new(big.Float).Mul(proxyFee, new(big.Float).SetInt64(chainFee.TokenBasic.Price))
	usdtFee = new(big.Float).Quo(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee := new(big.Float).Mul(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee = new(big.Float).Quo(tokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
	tokenFeeWithPrecision := new(big.Float).Mul(tokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
	isNative := false
	nativeTokenAmount := new(big.Float).SetInt64(0)
	// get optimistic L1 fee on ethereum
	if basedef.OPTIMISTIC_CROSSCHAIN_ID == getFeeReq.DstChainId {
		ethChainFee := new(models.ChainFee)
		res = db.Where("chain_id = ?", basedef.ETHEREUM_CROSSCHAIN_ID).Preload("TokenBasic").First(ethChainFee)
		if res.RowsAffected == 0 {
			c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", basedef.ETHEREUM_CROSSCHAIN_ID))
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		_, l1UsdtFee, _, err := fee.GetL1Fee(ethChainFee, getFeeReq.DstChainId)
		if err != nil {
			c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("get ethereum L1 fee failed. err=%v", err))
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		l1TokenFee := new(big.Float).Mul(l1UsdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		l1TokenFee = new(big.Float).Quo(l1TokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
		l1TokenFeeWithPrecision := new(big.Float).Mul(l1TokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
		tokenFee = new(big.Float).Add(tokenFee, l1TokenFee)
		tokenFeeWithPrecision = new(big.Float).Add(tokenFeeWithPrecision, l1TokenFeeWithPrecision)
	}

	{
		chainFeeJson, _ := json.Marshal(chainFee)
		logs.Error("chain fee: %s", string(chainFeeJson))
	}
	{
		tokenJson, _ := json.Marshal(token)
		logs.Error("token: %s", string(tokenJson))
	}

	if getFeeReq.SwapTokenHash != "" {
		tokenMap := new(models.TokenMap)
		res := db.Where("src_token_hash = ? and src_chain_id = ? and dst_chain_id = ?", getFeeReq.SwapTokenHash, getFeeReq.SrcChainId, getFeeReq.DstChainId).Preload("DstToken").First(tokenMap)
		if res.RowsAffected == 0 {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		if tokenMap.DstChainId != getFeeReq.DstChainId || tokenMap.DstToken == nil {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		tokenBalance, _ := new(big.Int).SetString("100000000000000000000000000000", 10)
		if tokenMap.DstChainId != basedef.PLT_CROSSCHAIN_ID {
			tokenBalance, err = cacheRedis.Redis.GetTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash)
			if err != nil {
				if tokenMap.SrcChainId == basedef.METIS_CROSSCHAIN_ID && (strings.EqualFold(tokenMap.SrcTokenHash, "deaddeaddeaddeaddeaddeaddeaddeaddead0000") || strings.EqualFold(tokenMap.SrcTokenHash, "F3eCc2FF57DF74aE638551b060864717EFE493d2")) && tokenMap.DstChainId == basedef.BSC_CROSSCHAIN_ID {
					lockproxy := "960Ff3132b72E3F0b1B9F588e7122d78BB5C4946"
					if basedef.ENV == basedef.TESTNET {
						lockproxy = "e6E89cde11B89D940D25c35eaec7aCB489D29820"
					}
					tokenBalance, err = common.GetProxyBalance(basedef.BSC_CROSSCHAIN_ID, tokenMap.DstTokenHash, lockproxy)
				} else if tokenMap.SrcChainId == basedef.BSC_CROSSCHAIN_ID && (strings.EqualFold(tokenMap.DstTokenHash, "deaddeaddeaddeaddeaddeaddeaddeaddead0000") || strings.EqualFold(tokenMap.DstTokenHash, "F3eCc2FF57DF74aE638551b060864717EFE493d2")) && tokenMap.DstChainId == basedef.METIS_CROSSCHAIN_ID {
					lockproxy := "bE46E4c47958A79E7F789ea94C5D8071a0DeE31e"
					if basedef.ENV == basedef.TESTNET {
						lockproxy = "B4004B93f1ce1E63131413cA201D35D1F3f40e5D"
					}
					tokenBalance, err = common.GetProxyBalance(basedef.METIS_CROSSCHAIN_ID, tokenMap.DstTokenHash, lockproxy)
				} else if tokenMap.SrcChainId == basedef.METIS_CROSSCHAIN_ID && tokenMap.DstChainId == basedef.BSC_CROSSCHAIN_ID {
					lockproxy := "fB571d4dd7039f96D34bB41E695AdC92dF4A332f"
					tokenBalance, err = common.GetProxyBalance(basedef.BSC_CROSSCHAIN_ID, tokenMap.DstTokenHash, lockproxy)
				} else if tokenMap.SrcChainId == basedef.BSC_CROSSCHAIN_ID && tokenMap.DstChainId == basedef.METIS_CROSSCHAIN_ID {
					lockproxy := "eFB5a01Ed9f3E94B646233FB68537C5Cb45e301D"
					tokenBalance, err = common.GetProxyBalance(basedef.METIS_CROSSCHAIN_ID, tokenMap.DstTokenHash, lockproxy)
				} else {
					tokenBalance, err = common.GetBalance(tokenMap.DstChainId, tokenMap.DstTokenHash)
				}
				if err != nil {
					tokenBalance, err = cacheRedis.Redis.GetLongTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash)
					if err != nil {
						c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
							getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
						c.ServeJSON()
						return
					}
				}
				setErr := cacheRedis.Redis.SetTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash, tokenBalance)
				if setErr != nil {
					logs.Error("qweasdredis SetTokenBalance err", setErr)
				}
				setErr1 := cacheRedis.Redis.SetLongTokenBalance(tokenMap.SrcChainId, tokenMap.DstChainId, tokenMap.DstTokenHash, tokenBalance)
				if setErr1 != nil {
					logs.Error("qweasdredis SetLongTokenBalance err", setErr1)
				}
			}
		}
		balance, result := new(big.Float).SetString(tokenBalance.String())
		if !result {
			c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
				getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
			c.ServeJSON()
			return
		}
		tokenBalanceWithoutPrecision := new(big.Float).Quo(balance, new(big.Float).SetInt64(basedef.Int64FromFigure(int(tokenMap.DstToken.Precision))))
		c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
			getFeeReq.SwapTokenHash, balance, tokenBalanceWithoutPrecision, isNative, nativeTokenAmount, feeTokenPrecison)
		c.ServeJSON()
	} else {
		c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision,
			getFeeReq.SwapTokenHash, new(big.Float).SetUint64(0), new(big.Float).SetUint64(0), isNative, nativeTokenAmount, feeTokenPrecison)
		c.ServeJSON()
	}
}

func (c *FeeController) CheckFee() {
	logs.Debug("check fee request: %s", string(c.Ctx.Input.RequestBody))
	var checkFeesReq models.CheckFeesReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeesReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	checkFeesReq4Nomal := make([]*models.CheckFeeReq, 0)
	checkFeesReq4O3 := make([]*models.CheckFeeReq, 0)
	for _, v := range checkFeesReq.Checks {
		if v.ChainId == basedef.O3_CROSSCHAIN_ID {
			checkFeesReq4O3 = append(checkFeesReq4O3, v)
		} else {
			checkFeesReq4Nomal = append(checkFeesReq4Nomal, v)
		}
	}
	checkFees4Normal := make([]*models.CheckFee, 0)
	checkFees4O3 := make([]*models.CheckFee, 0)
	if len(checkFeesReq4O3) > 0 {
		checkFees4O3 = c.CheckSwapFee(checkFeesReq4O3)
	}
	if len(checkFeesReq4Nomal) > 0 {
		checkFees4Normal = c.checkFee(checkFeesReq4Nomal)
	}
	checkFees := make([]*models.CheckFee, 0)
	checkFees = append(checkFees, checkFees4Normal...)
	checkFees = append(checkFees, checkFees4O3...)
	c.Data["json"] = models.MakeCheckFeesRsp(checkFees)
	c.ServeJSON()
}

func (c *FeeController) checkFee(Checks []*models.CheckFeeReq) []*models.CheckFee {
	hash2ChainId := make(map[string]uint64, 0)
	requestHashs := make([]string, 0)
	for _, check := range Checks {
		hash2ChainId[check.Hash] = check.ChainId
		requestHashs = append(requestHashs, check.Hash)
		requestHashs = append(requestHashs, basedef.HexStringReverse(check.Hash))
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	db.Model(&models.SrcTransaction{}).Where("(`key` in ? or `hash` in ?)", requestHashs, requestHashs).Find(&srcTransactions)
	key2Txhash := make(map[string]string, 0)
	isPolyProxy := make(map[string]bool, 0)

	for _, srcTransaction := range srcTransactions {
		prefix := srcTransaction.Key[0:8]
		if _, in := conf.PolyProxy[strings.ToUpper(srcTransaction.Contract)]; in {
			isPolyProxy[srcTransaction.Key] = true
			isPolyProxy[srcTransaction.Hash] = true
			isPolyProxy[basedef.HexStringReverse(srcTransaction.Hash)] = true
		}
		if prefix == "00000000" {
			chainId, ok := hash2ChainId[srcTransaction.Key]
			if ok && chainId == srcTransaction.ChainId {
				key2Txhash[srcTransaction.Key] = srcTransaction.Hash
			}
		} else {
			key2Txhash[srcTransaction.Hash] = srcTransaction.Hash
			key2Txhash[basedef.HexStringReverse(srcTransaction.Hash)] = srcTransaction.Hash
		}
	}
	checkHashes := make([]string, 0)
	for _, check := range Checks {
		newHash, ok := key2Txhash[check.Hash]
		if ok {
			checkHashes = append(checkHashes, newHash)
		}
	}
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	db.Table("wrapper_transactions").Where("hash in ?", checkHashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
	txHash2WrapperTransaction := make(map[string]*models.WrapperTransactionWithToken, 0)
	for _, wrapperTransactionWithToken := range wrapperTransactionWithTokens {
		txHash2WrapperTransaction[wrapperTransactionWithToken.Hash] = wrapperTransactionWithToken
	}
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	checkFees := make([]*models.CheckFee, 0)
	for _, check := range Checks {
		checkFee := &models.CheckFee{}
		checkFee.Hash = check.Hash
		checkFee.ChainId = check.ChainId
		checkFee.Amount = new(big.Float).SetInt64(0)
		checkFee.MinProxyFee = new(big.Float).SetInt64(0)
		_, ok := isPolyProxy[check.Hash]
		if !ok {
			checkFee.PayState = -2
			checkFees = append(checkFees, checkFee)
			continue
		}
		_, ok = chain2Fees[check.ChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		newHash, ok := key2Txhash[check.Hash]
		if !ok {
			checkFee.PayState = 0
			checkFees = append(checkFees, checkFee)
			continue
		}
		wrapperTransactionWithToken, ok := txHash2WrapperTransaction[newHash]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		chainFee, ok := chain2Fees[wrapperTransactionWithToken.DstChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		x := new(big.Int).Mul(&wrapperTransactionWithToken.FeeAmount.Int, big.NewInt(wrapperTransactionWithToken.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(wrapperTransactionWithToken.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		if feePay.Cmp(feeMin) >= 0 {
			checkFee.PayState = 1
		} else {
			checkFee.PayState = -1
			logs.Info("check fee PayState = -1 ChainId:%v Hash:%v feePay:%v < feeMin:%v", check.ChainId, check.Hash, feePay, feeMin)
		}
		checkFee.Amount = feePay
		checkFee.MinProxyFee = feeMin
		checkFees = append(checkFees, checkFee)
	}
	return checkFees
}

func (c *FeeController) getSwapSrcTransactions(o3Hashs []string) (map[string]string, error) {
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	res := db.Table("dst_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").
		Where("dst_transactions.hash in ?", o3Hashs).
		Joins("inner join poly_transactions on dst_transactions.poly_hash = poly_transactions.hash").
		Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.hash").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		return nil, res.Error
	}
	checkHashes := make(map[string]string, 0)
	for _, srcPolyDstRelation := range srcPolyDstRelations {
		checkHashes[srcPolyDstRelation.DstHash] = srcPolyDstRelation.SrcHash
	}
	return checkHashes, nil
}

func (c *FeeController) CheckSwapFee(Checks []*models.CheckFeeReq) []*models.CheckFee {
	hash2ChainId := make(map[string]uint64, 0)
	requestHashs := make([]string, 0)
	for _, check := range Checks {
		hash2ChainId[check.Hash] = check.ChainId
		requestHashs = append(requestHashs, check.Hash)
		requestHashs = append(requestHashs, basedef.HexStringReverse(check.Hash))
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	db.Model(&models.SrcTransaction{}).Where("(`key` in ? or `hash` in ?)", requestHashs, requestHashs).Find(&srcTransactions)
	key2Txhash := make(map[string]string, 0)
	o3Hashs := make([]string, 0)
	for _, srcTransaction := range srcTransactions {
		prefix := srcTransaction.Key[0:8]
		if prefix == "00000000" {
			chainId, ok := hash2ChainId[srcTransaction.Key]
			if ok && chainId == srcTransaction.ChainId {
				key2Txhash[srcTransaction.Key] = srcTransaction.Hash
			}
		} else {
			key2Txhash[srcTransaction.Hash] = srcTransaction.Hash
			key2Txhash[basedef.HexStringReverse(srcTransaction.Hash)] = srcTransaction.Hash
		}
		o3Hashs = append(o3Hashs, srcTransaction.Hash)
	}
	srcHashs, err := c.getSwapSrcTransactions(o3Hashs)
	if err != nil {
		return nil
	}

	checkHashes := make([]string, 0)
	for _, check := range Checks {
		newHash1, ok1 := key2Txhash[check.Hash]
		if ok1 {
			newHash2, ok2 := srcHashs[newHash1]
			if ok2 {
				checkHashes = append(checkHashes, newHash2)
			}
		}
	}
	//
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	db.Table("wrapper_transactions").Where("hash in ?", checkHashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
	txHash2WrapperTransaction := make(map[string]*models.WrapperTransactionWithToken, 0)
	for _, wrapperTransactionWithToken := range wrapperTransactionWithTokens {
		txHash2WrapperTransaction[wrapperTransactionWithToken.Hash] = wrapperTransactionWithToken
	}
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	checkFees := make([]*models.CheckFee, 0)
	for _, check := range Checks {
		checkFee := &models.CheckFee{}
		checkFee.Hash = check.Hash
		checkFee.ChainId = check.ChainId
		checkFee.Amount = new(big.Float).SetInt64(0)
		checkFee.MinProxyFee = new(big.Float).SetInt64(0)
		_, ok := chain2Fees[check.ChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		newHash, ok := key2Txhash[check.Hash]
		if !ok {
			checkFee.PayState = 0
			checkFees = append(checkFees, checkFee)
			continue
		}
		newHash, ok = srcHashs[newHash]
		if !ok {
			checkFee.PayState = 0
			checkFees = append(checkFees, checkFee)
			continue
		}
		wrapperTransactionWithToken, ok := txHash2WrapperTransaction[newHash]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		chainFee, ok := chain2Fees[wrapperTransactionWithToken.DstChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		x := new(big.Int).Mul(&wrapperTransactionWithToken.FeeAmount.Int, big.NewInt(wrapperTransactionWithToken.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(wrapperTransactionWithToken.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		if feePay.Cmp(feeMin) >= 0 {
			checkFee.PayState = 1
		} else {
			checkFee.PayState = -1
		}
		checkFee.Amount = feePay
		checkFee.MinProxyFee = feeMin
		checkFees = append(checkFees, checkFee)
	}
	return checkFees
}

func SetCoinRankFilterInfo(RiskyCoinHandleConfig *conf.RiskyCoinHandleConfig) {
	if RiskyCoinHandleConfig == nil {
		riskyCoinRisingRate = big.NewFloat(250)
		riskyCoinRankThreshold = 100
	}
	riskyCoinRankThreshold = RiskyCoinHandleConfig.RiskyCoinRankThreshold
	riskyCoinRisingRate = big.NewFloat(float64(RiskyCoinHandleConfig.RiskyCoinRisingRate))
}
func SetProxyFeeRatioMap(config *conf.Config) {
	proxyFeeRatioMap = make(map[uint64]int64, 0)
	for _, v := range config.FeeListenConfig {
		proxyFeeRatioMap[v.ChainId] = v.ProxyFee
	}
}
