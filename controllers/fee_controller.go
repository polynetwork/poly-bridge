package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
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
	fee := chainFee.ProxyFee * chainFee.TokenBasic.AvgPrice / token.TokenBasic.AvgPrice
	c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.ChainId, getFeeReq.Hash, float64(fee))
	c.ServeJSON()
}

func (c *FeeController) CheckFee() {
	var checkFeeReq models.CheckFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeeReq); err != nil {
		panic(err)
	}
	db := newDB()
	transaction := new(models.Transaction)
	db.Where("hash = ?", checkFeeReq.Hash).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(transaction)
	chainFee := new(models.ChainFee)
	db.Where("chain_id = ?", transaction.DstChainId).Preload("TokenBasic").First(chainFee)
	hasPay := false
	feePay := transaction.FeeAmount * transaction.FeeToken.TokenBasic.AvgPrice
	feeMin := chainFee.MinFee * chainFee.TokenBasic.AvgPrice
	if feePay >= feeMin {
		hasPay = true
	}
	c.Data["json"] = models.MakeCheckFeeRsp(hasPay, float64(feePay))
	c.ServeJSON()
}
