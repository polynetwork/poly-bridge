package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"poly-swap/models"
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Transactions() {
	var transactionsReq models.TransactionsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil {
		panic(err)
	}
	db := newDB()
	transactions := make([]*models.WrapperTransaction, 0)
	db.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsRsp(transactionsReq.PageSize, transactionsReq.PageNo, (int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize,
		int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactoinsOfUser() {
	var transactionsOfUserReq models.TransactionsOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfUserReq); err != nil {
		panic(err)
	}
	db := newDB()
	transactions := make([]*models.WrapperTransaction, 0)
	db.Where("user = ?", transactionsOfUserReq.User).Limit(transactionsOfUserReq.PageSize).Offset(transactionsOfUserReq.PageSize * transactionsOfUserReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Where("user = ?", transactionsOfUserReq.User).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsOfUserRsp(transactionsOfUserReq.PageSize, transactionsOfUserReq.PageNo,
		(int(transactionNum)+transactionsOfUserReq.PageSize-1)/transactionsOfUserReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}
