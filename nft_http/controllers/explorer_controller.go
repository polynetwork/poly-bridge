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
