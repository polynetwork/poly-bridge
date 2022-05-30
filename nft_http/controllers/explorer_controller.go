package controllers

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"poly-bridge/basedef"
	"poly-bridge/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type ExplorerController struct {
	web.Controller
}

func (c *ExplorerController) Transactions() {
	var req TransactionBriefsReq
	if !input(&c.Controller, &req) {
		return
	}
	if !checkPageSize(&c.Controller, req.PageSize) {
		return
	}

	relations := make([]*TransactionBriefRelation, 0)
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	limit := req.PageSize
	offset := req.PageSize * (req.PageNo - 1)
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
	if !checkPageSize(&c.Controller, req.PageSize) {
		return
	}

	relations := make([]*TransactionBriefRelation, 0)
	limit := req.PageSize
	offset := req.PageSize * req.PageNo
	var transactionNum int64

	if req.ChainId == basedef.POLY_CROSSCHAIN_ID {
		db.Raw("select wp.*, tr.amount as token_id, tr.asset as src_asset "+
			"from wrapper_transactions wp "+
			"left join src_transfers as tr on wp.hash=tr.tx_hash "+
			"where wp.standard=? and (wp.user in ? or wp.dst_user in ?) "+
			"order by wp.time desc "+
			"limit ? offset ?",
			models.TokenTypeErc721, req.Addresses, req.Addresses, limit, offset).
			Find(&relations)
		db.Model(&models.SrcTransfer{}).
			Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
			Where("src_transfers.standard = ? and (`from` in ? or src_transfers.dst_user in ?)", models.TokenTypeErc721, req.Addresses, req.Addresses).
			Count(&transactionNum)
	} else {
		db.Raw("select wp.*, tr.amount as token_id, tr.asset as src_asset "+
			"from wrapper_transactions wp "+
			"left join src_transfers as tr on wp.hash=tr.tx_hash "+
			"where wp.standard=? and (wp.user in ? or wp.dst_user in ?) and (wp.src_chain_id = ? or wp.dst_chain_id = ?)"+
			"order by wp.time desc "+
			"limit ? offset ?",
			models.TokenTypeErc721, req.Addresses, req.Addresses, req.ChainId, req.ChainId, limit, offset).
			Find(&relations)
		db.Model(&models.SrcTransfer{}).
			Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
			Where("src_transfers.standard = ? and (`from` in ? or src_transfers.dst_user in ?) and (wrapper_transactions.src_chain_id = ? or wrapper_transactions.dst_chain_id = ?)",
				models.TokenTypeErc721, req.Addresses, req.Addresses, req.ChainId, req.ChainId).
			Count(&transactionNum)
	}

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

	startTime := time.Now().UnixNano()
	relations := make([]*TransactionDetailRelation, 0)
	tx := &struct{ Hash string }{}
	err := db.Raw(`select hash from src_transactions where hash=?
		UNION select s.hash from src_transactions s left join poly_transactions p on p.src_hash=s.hash where p.hash=?
		UNION select s.hash from src_transactions s left join poly_transactions p on p.src_hash=s.hash
		left join dst_transactions d on d.poly_hash=p.hash where d.hash=?`,
		req.Hash, req.Hash, req.Hash).First(tx).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || tx.Hash == "" {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("relations does not exist"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash").
		Where("src_transactions.hash = ? ", tx.Hash).
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
		Find(&relations)

	endTime := time.Now().UnixNano()
	logs.Info("mysql spent time %d", (endTime-startTime)/int64(time.Millisecond))

	if res.RowsAffected == 0 {
		output(&c.Controller, nil)
		return
	}

	data := new(TransactionDetailRsp).instance(relations[0])
	fillMetaInfo(data)

	output(&c.Controller, data)
}

func fillMetaInfo(data *TransactionDetailRsp) {
	if data.Transaction == nil {
		return
	}

	chainId := data.Transaction.SrcChainId
	sdk, inquirer, _, err := selectNodeAndWrapper(chainId)
	if err != nil {
		return
	}
	if data.SrcTransaction == nil {
		return
	}

	tokenId, ok := string2Big(data.Transaction.TokenId)
	if !ok {
		return
	}
	asset := selectNFTAsset(data.SrcTransaction.AssetHash)
	if asset != nil {
		item, err := getSingleItem(sdk, inquirer, asset, tokenId, "")
		if err != nil {
			logs.Error("fillMetaInfo err: %v", err)
			return
		}
		data.Meta = item
	}
}
