package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/fee"
	"strings"
)

const (
	SKIP     models.CheckFeeStatus = -2 // Skip since not our tx
	NOT_PAID models.CheckFeeStatus = -1 // Not paid or paid too low
	MISSING  models.CheckFeeStatus = 0  // Tx not received yet
	PAID     models.CheckFeeStatus = 1  // Paid and enough pass
)

func (c *FeeController) NewCheckFee() {
	logs.Debug("new check fee request: %s", string(c.Ctx.Input.RequestBody))
	var mapCheckFeesReq map[string]*models.CheckFeeRequest
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &mapCheckFeesReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	srcHashs := make([]string, 0)
	for k, v := range mapCheckFeesReq {
		srcTransaction, err := checkFeeSrcTransaction(v.ChainId, v.TxId)
		if err != nil {
			//has not listen src_transaction
			v.Status = MISSING
			logs.Info("check fee poly_hash %s MISSING,hasn't src_Transaction %s", k, err)
			continue
		}
		if len(conf.PolyProxy) > 0 {
			if _, in := conf.PolyProxy[strings.ToUpper(srcTransaction.Contract)]; !in {
				//is not poly proxy
				v.Status = SKIP
				logs.Info("check fee poly_hash %s SKIP,is not poly proxy", k)
				continue
			}
		}
		v.SrcTransaction = srcTransaction
		srcHashs = append(srcHashs, srcTransaction.Hash)
	}
	checkFeewrapperTransaction(srcHashs, mapCheckFeesReq)
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	for k, v := range mapCheckFeesReq {
		if v.SrcTransaction != nil {
			exists, _ := cacheRedis.Redis.Exists(cacheRedis.MarkTxAsPaidPrefix + v.SrcTransaction.Hash)
			if exists {
				logs.Info("check fee poly_hash %s marked as paid", k)
				v.Status = PAID
				continue
			}
		}

		if v.WrapperTransactionWithToken == nil {
			if v.SrcTransaction != nil {
				//has src_transaction but not wrapper_transaction
				if v.SrcTransaction.ChainId == basedef.NEO_CROSSCHAIN_ID ||
					v.SrcTransaction.DstChainId == basedef.NEO_CROSSCHAIN_ID ||
					v.SrcTransaction.ChainId == basedef.NEO3_CROSSCHAIN_ID ||
					v.SrcTransaction.DstChainId == basedef.NEO3_CROSSCHAIN_ID {
					v.Status = SKIP
					logs.Info("check fee poly_hash %s SKIP, because it is a NEO/NEO3 tx with no wrapper_transactions", k)
				}

				v.Status = NOT_PAID
				logs.Info("check fee poly_hash %s NOT_PAID,src_transaction but not wrapper_transaction", k)
				continue
			}
		} else {
			chainFee, ok := chain2Fees[v.WrapperTransactionWithToken.DstChainId]
			if !ok {
				v.Status = NOT_PAID
				logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't DstChainId's fee", k)
				continue
			}
			x := new(big.Int).Mul(&v.WrapperTransactionWithToken.FeeAmount.Int, big.NewInt(v.WrapperTransactionWithToken.FeeToken.TokenBasic.Price))
			feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(v.WrapperTransactionWithToken.FeeToken.Precision))))
			feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
			x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
			feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
			feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
			feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))

			// get optimistic L1 fee on ethereum
			if chainFee.ChainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
				ethChainFee, ok := chain2Fees[basedef.ETHEREUM_CROSSCHAIN_ID]
				if !ok {
					v.Status = NOT_PAID
					logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't ethereum fee", k)
					continue
				}

				L1MinFee, _, err := fee.GetL1Fee(ethChainFee, chainFee.ChainId)
				if err != nil {
					v.Status = NOT_PAID
					logs.Info("check fee poly_hash %s NOT_PAID, get L1 fee failed. err=%v", k, err)
					continue
				}
				feeMin = new(big.Float).Add(feeMin, L1MinFee)
			}

			FluctuatingFeeMin := feeMin
			excludeChainIds := map[uint64]interface{}{basedef.BSC_CROSSCHAIN_ID: nil, basedef.ARBITRUM_CROSSCHAIN_ID: nil, basedef.ETHEREUM_CROSSCHAIN_ID: nil, basedef.OPTIMISTIC_CROSSCHAIN_ID: nil}
			polyTx := new(models.PolyTransaction)
			res := db.Model(&models.PolyTransaction{}).
				Where("hash =?", v.PolyHash).First(polyTx)
			if res.Error == nil {
				if _, ok := excludeChainIds[polyTx.DstChainId]; !ok {
					FluctuatingFeeMin = new(big.Float).Mul(FluctuatingFeeMin, new(big.Float).SetFloat64(0.9))
				}
			}

			v.Paid, _ = feePay.Float64()
			v.Min, _ = FluctuatingFeeMin.Float64()
			if feePay.Cmp(feeMin) >= 0 {
				v.Status = PAID
				logs.Info("check fee poly_hash %s PAID,feePay %v >= feeMin %v", k, v.Paid, v.Min)
			} else if feePay.Cmp(FluctuatingFeeMin) >= 0 {
				v.Status = PAID
				logs.Info("check fee poly_hash %s PAID,feePay %v >= FluctuatingFeeMin %v", k, v.Paid, v.Min)
			} else {
				v.Status = NOT_PAID
				logs.Info("check fee poly_hash %s NOT_PAID,feePay %v < FluctuatingFeeMin %v", k, v.Paid, v.Min)
			}
		}
	}
	c.Data["json"] = mapCheckFeesReq
	c.ServeJSON()
	return
}

func checkFeeSrcTransaction(chainId uint64, txId string) (*models.SrcTransaction, error) {
	transaction := new(models.SrcTransaction)
	if strings.Contains(txId, "00000000") {
		res := db.Model(&models.SrcTransaction{}).
			Where("chain_id=? and `key` =?", chainId, txId).
			First(transaction)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		res := db.Model(&models.SrcTransaction{}).
			Where("chain_id=? and `hash` =?", chainId, txId).
			First(transaction)
		if res.Error != nil {
			res := db.Model(&models.SrcTransaction{}).
				Where("chain_id=? and `hash` =?", chainId, basedef.HexStringReverse(txId)).First(transaction)
			if res.Error != nil {
				return nil, res.Error
			}
		}
	}
	if chainId != basedef.O3_CROSSCHAIN_ID {
		return transaction, nil
	}

	srcTransaction := new(models.SrcTransaction)
	res := db.Debug().Table("src_transactions").
		Joins("inner join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("inner join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Where("dst_transactions.hash = ?", transaction.Hash).
		First(srcTransaction)
	if res.Error != nil {
		return nil, res.Error
	}
	return srcTransaction, nil
}

func checkFeewrapperTransaction(srcHashs []string, mapCheckFeesReq map[string]*models.CheckFeeRequest) {
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	db.Table("wrapper_transactions").Where("hash in ?", srcHashs).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
	for _, v := range mapCheckFeesReq {
		for _, wrapper := range wrapperTransactionWithTokens {
			if v.SrcTransaction != nil && v.SrcTransaction.Hash == wrapper.Hash {
				v.WrapperTransactionWithToken = wrapper
				break
			}
		}
	}
}
