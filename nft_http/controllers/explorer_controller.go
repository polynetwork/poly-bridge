package controllers

import (
	"poly-bridge/models"

	"github.com/astaxie/beego"
)

type ExplorerController struct {
	beego.Controller
}

func (c *ExplorerController) Transactions() {
	var req TransactionBriefsReq
	if !input(&c.Controller, &req) {
		return
	}

	relations := make([]*TransactionBriefRelation, 0)
	limit := req.PageSize
	offset := req.PageSize * req.PageNo
	db.Raw("select wp.*, tr.amount as token_id, tr.asset as src_asset "+
		"from wrapper_transactions wp "+
		"left join src_transfers as tr on wp.hash=tr.tx_hash "+
		"where wp.standard=? "+
		"order by wp.time desc "+
		"limit ? offset ?", models.TokenTypeErc721, limit, offset).
		Find(&relations)

	transactionNum := txCounter.Number()
	totalPage := (int(transactionNum) + req.PageSize - 1) / req.PageSize
	totalCnt := int(transactionNum)

	list := make([]*TransactionBriefRsp, 0)
	for _, v := range relations {
		tk := selectNFTAsset(v.SrcAsset)
		data := new(TransactionBriefRsp).instance(tk.TokenBasicName, v)
		list = append(list, data)
	}

	resp := new(TransactionBriefsRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, list)
	output(&c.Controller, resp)
}

func (c *ExplorerController) TransactionsOfAddress() {
	var req TransactionBriefsOfAddressReq
	if !input(&c.Controller, &req) {
		return
	}

	relations := make([]*TransactionBriefRelation, 0)
	limit := req.PageSize
	offset := req.PageSize * req.PageNo
	db.Raw("select wp.*, tr.amount as token_id, tr.asset as src_asset "+
		"from wrapper_transactions wp "+
		"left join src_transfers as tr on wp.hash=tr.tx_hash "+
		"where wp.standard=? and (wp.user in ? or wp.dst_user in ?) "+
		"order by wp.time desc "+
		"limit ? offset ?",
		models.TokenTypeErc721, req.Addresses, req.Addresses, limit, offset).
		Find(&relations)

	var transactionNum int64
	db.Model(&models.SrcTransfer{}).
		Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
		Where("src_transfers.standard = ? and (`from` in ? or src_transfers.dst_user in ?)", models.TokenTypeErc721, req.Addresses, req.Addresses).
		Count(&transactionNum)
	totalPage := (int(transactionNum) + req.PageSize - 1) / req.PageSize
	totalCnt := int(transactionNum)

	list := make([]*TransactionBriefRsp, 0)
	for _, v := range relations {
		tk := selectNFTAsset(v.SrcAsset)
		data := new(TransactionBriefRsp).instance(tk.TokenBasicName, v)
		list = append(list, data)
	}

	resp := new(TransactionBriefsRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, list)
	output(&c.Controller, resp)
}
