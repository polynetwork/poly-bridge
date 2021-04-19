package controllers

import (
	"github.com/astaxie/beego/logs"
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

func (c *ExplorerController) TransactionDetail() {
	var req TransactionDetailReq
	if !input(&c.Controller, &req) {
		return
	}

	relation := new(TransactionDetailRelation)
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash").
		Where("src_transactions.hash = ?", req.Hash).
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Order("src_transactions.time desc").
		Find(relation)

	if res.RowsAffected == 0 {
		output(&c.Controller, nil)
		return
	}

	data := new(TransactionDetailRsp).instance(relation)
	fillMetaInfo(data)

	output(&c.Controller, data)
}

func fillMetaInfo(data *TransactionDetailRsp) {
	if data.Transaction == nil {
		return
	}

	chainId := data.Transaction.SrcChainId
	sdk := selectNode(chainId)
	wrapper := selectWrapper(chainId)
	if wrapper == emptyAddr {
		return
	}
	if data.SrcTransaction == nil {
		return
	}

	asset := selectNFTAsset(data.SrcTransaction.AssetHash)
	tokenId, ok := string2Big(data.Transaction.TokenId)
	if !ok {
		return
	}

	item, err := getSingleItem(sdk, wrapper, asset, tokenId, "")
	if err != nil {
		logs.Error("fillMetaInfo err: %v", err)
		return
	}

	data.Meta = item
}
