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
	"poly-bridge/utils/decimal"
	"poly-bridge/utils/fee"
	"strings"
)

const (
	SKIP        models.CheckFeeStatus = -2 // Skip since not our tx
	NOT_PAID    models.CheckFeeStatus = -1 // Not paid or paid too low
	MISSING     models.CheckFeeStatus = 0  // Tx not received yet
	PAID        models.CheckFeeStatus = 1  // Paid and enough pass
	EstimatePay models.CheckFeeStatus = 2  // Paid but need EstimateGas
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
		if v.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID || v.ChainId == basedef.ZILLIQA_CROSSCHAIN_ID {
			//switcheo || zilliqa
			v.Status = SKIP
			logs.Info("check fee poly_hash %s SKIP,is switcheo or zilliqa", k)
			continue
		}
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
	//get chain fee
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	for k, v := range mapCheckFeesReq {
		//check fee from cache（special case）
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
					continue
				} else if v.SrcTransaction.ChainId == basedef.RIPPLE_CROSSCHAIN_ID {
					v.Status = MISSING
					logs.Info("check fee poly_hash %s MISSING, chain is ripple, src_transaction but not wrapper_transaction", k)
					continue
				} else {
					v.Status = NOT_PAID
					logs.Info("check fee poly_hash %s NOT_PAID,src_transaction but not wrapper_transaction", k)
					continue
				}
			}
		} else {
			//check db
			if v.WrapperTransactionWithToken.IsPaid == true {
				logs.Info("check fee poly_hash %s marked as paid in db", k)
				v.Status = PAID
				continue
			}
			chainFee, ok := chain2Fees[v.WrapperTransactionWithToken.DstChainId]
			if !ok {
				v.Status = NOT_PAID
				logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't DstChainId's fee", k)
				continue
			}
			//money paid in wrapper
			feePay, feeMin, gasPay := fee.CheckFeeCal(chainFee, v.WrapperTransactionWithToken.FeeToken, v.WrapperTransactionWithToken.FeeAmount)

			// get optimistic L1 fee on ethereum
			if chainFee.ChainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
				ethChainFee, ok := chain2Fees[basedef.ETHEREUM_CROSSCHAIN_ID]
				if !ok {
					v.Status = NOT_PAID
					logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't ethereum fee", k)
					continue
				}

				L1MinFee, _, _, l1FeeWei, err := fee.GetL1Fee(ethChainFee, chainFee.ChainId)
				if err != nil {
					v.Status = NOT_PAID
					logs.Info("check fee poly_hash %s NOT_PAID, get L1 fee failed. err=%v", k, err)
					continue
				}
				feeMin = new(big.Float).Add(feeMin, L1MinFee)
				gasPay = new(big.Float).Sub(gasPay, l1FeeWei)
			}

			v.Paid, _ = feePay.Float64()
			v.Min, _ = feeMin.Float64()

			if _, in := conf.EstimateProxy[strings.ToUpper(v.SrcTransaction.Contract)]; in {
				//is estimateGas proxy
				if gasPay.Cmp(new(big.Float).SetInt64(0)) <= 0 {
					v.Status = NOT_PAID
					continue
				}
				v.Status = EstimatePay
				if minFee, in := conf.EstimateFeeMin[v.WrapperTransactionWithToken.DstChainId]; in {
					if minFee > 0 && minFee < 100 {
						gasPay = new(big.Float).Mul(gasPay, new(big.Float).SetInt64(100))
						gasPay = new(big.Float).Quo(gasPay, new(big.Float).SetInt64(minFee))
					}
				}
				//compare current result with gasPay in db
				v.PaidGas, _ = gasPay.Float64()
				if v.WrapperTransactionWithToken.PaidGas != nil {
					PaidGasDecimal := decimal.NewFromBigInt(&v.WrapperTransactionWithToken.PaidGas.Int, 0).Div(decimal.NewFromInt(100))
					PaidGasFloat64, _ := PaidGasDecimal.Float64()
					PaidGasDb := big.NewFloat(PaidGasFloat64)
					if PaidGasDb.Cmp(gasPay) > 0 {
						logs.Info("paid gas in db larger", "gasdb", PaidGasDb, "gascal", gasPay)
						v.PaidGas = PaidGasFloat64
					} else {
						logs.Info("paid gas in db smaller", "gasdb", PaidGasDb, "gascal", gasPay)
					}
				}
				logs.Info("check fee poly_hash %s is EstimateProxy,PaidGas %v", k, v.PaidGas)
				continue
			}

			if feePay.Cmp(feeMin) >= 0 {
				v.Status = PAID
				logs.Info("check fee poly_hash %s PAID,feePay %v >= feeMin %v", k, v.Paid, v.Min)
			} else {
				v.Status = NOT_PAID
				logs.Info("check fee poly_hash %s NOT_PAID,feePay %v < feeMin %v", k, v.Paid, v.Min)
			}
		}
	}
	c.Data["json"] = mapCheckFeesReq
	c.ServeJSON()
	return
}

//checkFeeSrcTransaction fetch src transaction record from db by chain ID and txId
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

//checkFeewrapperTransaction fetch wrapper transaction record from db
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
