package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"math/big"
	"poly-swap/models"
)

type FeeController struct {
	beego.Controller
}

func (c *FeeController) GetFee() {
	var getFeeReq models.GetFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &getFeeReq); err != nil {
		panic(err)
	}
	db := newDB()
	token := new(models.Token)
	db.Where("hash = ?", getFeeReq.Hash).Preload("TokenBasic").First(token)
	chainFee := new(models.ChainFee)
	db.Where("chain_id = ?", getFeeReq.ChainId).Preload("TokenBasic").First(chainFee)
	//proxyFee := &chainFee.ProxyFee.Int
	x := new(big.Int).Mul(&chainFee.ProxyFee.Int, big.NewInt(chainFee.TokenBasic.AvgPrice))
	y := new(big.Int).Div(x, big.NewInt(token.TokenBasic.AvgPrice))
	c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.ChainId, getFeeReq.Hash, float64(y.Int64()))
	c.ServeJSON()
}

func (c *FeeController) CheckFee() {
	var checkFeeReq models.CheckFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeeReq); err != nil {
		panic(err)
	}
	db := newDB()
	transaction := new(models.WrapperTransaction)
	db.Where("hash = ?", checkFeeReq.Hash).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(transaction)
	chainFee := new(models.ChainFee)
	db.Where("chain_id = ?", transaction.DstChainId).Preload("TokenBasic").First(chainFee)
	hasPay := false
	feePay := new(big.Int).Mul(&transaction.FeeAmount.Int, big.NewInt(transaction.FeeToken.TokenBasic.AvgPrice))
	feeMin := new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.AvgPrice))
	if feePay.Cmp(feeMin) >= 0 {
		hasPay = true
	}
	c.Data["json"] = models.MakeCheckFeeRsp(hasPay, float64(feePay.Int64()))
	c.ServeJSON()
}
