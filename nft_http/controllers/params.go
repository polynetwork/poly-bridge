package controllers

import (
	"poly-bridge/basedef"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
)

type PolyBridgeInfoResp struct {
	Version  string
	URL      string
	Entrance []*ContractEntrance
}

type ContractEntrance struct {
	ChainId         uint64
	ChainName       string
	WrapperContract []string
}

type ErrorRsp struct {
	Message string
}

func MakeErrorRsp(messgae string) *ErrorRsp {
	errorRsp := &ErrorRsp{
		Message: messgae,
	}
	return errorRsp
}

type AssetItems struct {
	Asset   *AssetRsp
	Items   []*Item
	HasMore bool
}

type Item struct {
	AssetName string
	TokenId   string
	Name      string
	Url       string
	Image     string
	Desc      string
	Meta      string
}

func (i *Item) instance(assetName string, tokenId string, profile *models.NFTProfile) *Item {
	i.AssetName = assetName
	i.TokenId = tokenId
	if profile == nil {
		return i
	}
	i.Name = profile.Name
	i.Url = profile.Url
	i.Image = profile.Image
	i.Desc = profile.Description

	// todo(fuk): useless field
	//i.Meta = profile.Text
	return i
}

type AssetRsp struct {
	Hash    string
	ChainId uint64
	Name    string
	BaseUri string
}

func (s *AssetRsp) instance(t *models.Token) *AssetRsp {
	s.Hash = t.Hash
	s.ChainId = t.ChainId
	s.Name = t.Name
	if t.TokenBasic != nil {
		s.BaseUri = t.TokenBasic.Meta
	}
	return s
}

type HomeReq struct {
	ChainId  uint64
	PageSize int
	PageNo   int
}

type HomeRsp struct {
	TotalCount uint64
	Assets     []*AssetItems
}

func (s *HomeRsp) instance(totalCnt int, list []*AssetItems) *HomeRsp {
	s.Assets = list
	s.TotalCount = uint64(totalCnt)
	return s
}

type AssetsReq struct {
	ChainId uint64
}

type AssetsRsp struct {
	TotalCount uint64
	Assets     []*AssetMap
}

type AssetMap struct {
	AssetRsp
	DstAssets []*AssetRsp
}

func (s *AssetMap) instance(t *models.Token) *AssetMap {
	s.AssetRsp = *(s.AssetRsp.instance(t))
	s.DstAssets = make([]*AssetRsp, 0)
	if t.TokenMaps == nil {
		return s
	}
	for _, v := range t.TokenMaps {
		if v.DstToken != nil {
			s.DstAssets = append(s.DstAssets, new(AssetRsp).instance(v.DstToken))
		}
	}
	return s
}

func (s *AssetsRsp) instance(assets []*models.Token) *AssetsRsp {
	s.TotalCount = uint64(len(assets))
	if s.Assets == nil {
		s.Assets = make([]*AssetMap, 0)
	}
	for _, asset := range assets {
		s.Assets = append(s.Assets, new(AssetMap).instance(asset))
	}
	return s
}

type AssetReq struct {
	ChainId uint64
	Hash    string
}

//
//type NFTAssetRsp struct {
//	Hash       string
//	ChainId    uint64
//	Name       string
//	Disable    int64
//	BaseUri    string
//	AssetBasic *NFTAssetBasicRsp
//	AssetMaps  []*NFTAssetMapRsp
//}
//
//func MakeNFTAssetRsp(asset *models.Asset) *NFTAssetRsp {
//	rsp := &NFTAssetRsp{
//		Hash:    asset.Hash,
//		ChainId: asset.ChainId,
//		Name:    asset.Name,
//		Disable: property2Disable(asset.Property),
//	}
//	if asset.TokenBasic != nil {
//		rsp.BaseUri = asset.TokenBasic.Meta
//	}
//	if asset.TokenBasic != nil {
//		rsp.AssetBasic = MakeAssetBasicRsp(asset.TokenBasic)
//	}
//	if asset.TokenMaps != nil {
//		for _, m := range asset.TokenMaps {
//			rsp.AssetMaps = append(rsp.AssetMaps, MakeNFTAssetMapRsp(m))
//		}
//	}
//	return rsp
//}

type ItemsOfAddressReq struct {
	ChainId  uint64
	Asset    string
	Address  string
	TokenId  string
	PageSize int
	PageNo   int
}

type ItemsOfAddressRsp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	Items      []*Item
}

func (s *ItemsOfAddressRsp) instance(pageSize, pageNo, totalPage, totalCnt int, items []*Item) *ItemsOfAddressRsp {
	s.PageSize = pageSize
	s.PageNo = pageNo
	s.TotalCount = totalCnt
	s.TotalPage = totalPage
	if items == nil {
		s.Items = make([]*Item, 0)
		return s
	}
	s.Items = items
	return s
}

type WrapperTransactionReq struct {
	Hash string
}

type WrapperTransactionRsp struct {
	Hash         string
	User         string
	SrcChainId   uint64
	BlockHeight  uint64
	Time         uint64
	DstChainId   uint64
	DstUser      string
	ServerId     uint64
	FeeTokenHash string
	FeeAmount    string
	State        uint64
}

func (s *WrapperTransactionRsp) instance(transaction *models.WrapperTransaction) *WrapperTransactionRsp {

	s.Hash = transaction.Hash
	s.User = transaction.User
	s.SrcChainId = transaction.SrcChainId
	s.BlockHeight = transaction.BlockHeight
	s.Time = transaction.Time
	s.DstChainId = transaction.DstChainId
	s.DstUser = transaction.DstUser
	s.ServerId = transaction.ServerId
	s.FeeTokenHash = transaction.FeeTokenHash
	s.FeeAmount = transaction.FeeAmount.String()
	s.State = transaction.Status

	return s
}

type TransactionBriefsReq struct {
	PageSize int
	PageNo   int
}

type TransactionBriefsOfAddressReq struct {
	PageSize  int
	PageNo    int
	Addresses []string
	ChainId   uint64
}

type TransactionBriefRelation struct {
	models.WrapperTransaction
	SrcAsset string
	TokenId  string
}

type TransactionBriefRsp struct {
	Hash             string
	Status           uint64
	BlockHeight      uint64
	SrcChainId       uint64
	SrcChainName     string `json:"srcchainname"`
	SrcChainExplorer string `json:"srcchainexplorer"`
	SrcChainLogo     string `json:"srcchainlogo"`
	DstChainId       uint64
	DstChainName     string `json:"dstchainname"`
	DstChainExplorer string `json:"dstchainexplorer"`
	DstChainLogo     string `json:"dstchainlogo"`
	Time             uint64
	TokenId          string
	AssetName        string
	From             string
	To               string
	NftImage         string `json:"image"`
}

func (s *TransactionBriefRsp) instance(assetName string, r *TransactionBriefRelation) *TransactionBriefRsp {
	s.AssetName = assetName
	s.Hash = r.Hash
	s.Status = r.Status
	s.BlockHeight = r.BlockHeight
	s.SrcChainId = r.SrcChainId
	s.SrcChainName = models.ChainId2Name(r.SrcChainId)
	s.SrcChainLogo = models.ChainId2ChainCache(r.SrcChainId).ChainLogo
	s.DstChainId = r.DstChainId
	s.DstChainName = models.ChainId2Name(r.DstChainId)
	s.DstChainLogo = models.ChainId2ChainCache(r.DstChainId).ChainLogo
	s.Time = r.Time
	s.TokenId = r.TokenId
	s.From = r.User
	s.To = r.DstUser
	return s
}

type TransactionBriefsRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionBriefRsp
}

func (s *TransactionBriefsRsp) instance(
	pageSize, pageNo, totalPage, totalCount int,
	txs []*TransactionBriefRsp,
) *TransactionBriefsRsp {

	s.PageSize = pageSize
	s.PageNo = pageNo
	s.TotalPage = totalPage
	s.TotalCount = totalCount
	s.Transactions = txs
	return s
}

type TransactionDetailRelation struct {
	SrcHash            string
	WrapperTransaction *models.WrapperTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	SrcTransaction     *models.SrcTransaction     `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash           string
	PolyTransaction    *models.PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
	DstHash            string
	DstTransaction     *models.DstTransaction `gorm:"foreignKey:DstHash;references:Hash"`
}

type SideChainRsp struct {
	Hash             string
	ChainId          uint64
	ChainName        string `json:"chainname"`
	ChainLogo        string `json:"chainlogo"`
	ChainExplorerUrl string `json:"chainexplorerurl"`
	Asset            string
	AssetHash        string
	TokenType        string
	BlockHeight      uint64
	Time             uint64
	Fee              string
	FeeName          string `json:"feename"`
	FeeLogo          string `json:"feelogo"`
	Status           uint64
	From             string
	To               string
}

type PolyChainRsp struct {
	Hash        string
	Time        uint64
	BlockHeight uint64
	Status      uint64
	From        string
	To          string
}

type TransactionDetailReq struct {
	Hash string
}

type TransactionDetailRsp struct {
	Transaction     *TransactionBriefRsp
	SrcTransaction  *SideChainRsp
	DstTransaction  *SideChainRsp
	PolyTransaction *PolyChainRsp
	Meta            *Item
}

func (s *TransactionDetailRsp) instance(r *TransactionDetailRelation) *TransactionDetailRsp {
	if r == nil {
		return nil
	}

	s.Transaction = new(TransactionBriefRsp)
	s.SrcTransaction = new(SideChainRsp)
	s.DstTransaction = new(SideChainRsp)
	s.PolyTransaction = new(PolyChainRsp)

	s.Transaction.Hash = r.WrapperTransaction.Hash
	s.Transaction.Status = r.WrapperTransaction.Status
	s.Transaction.BlockHeight = r.WrapperTransaction.BlockHeight
	s.Transaction.SrcChainId = r.WrapperTransaction.SrcChainId
	s.Transaction.SrcChainExplorer = models.ChainId2ChainCache(r.WrapperTransaction.SrcChainId).ChainExplorerUrl
	s.Transaction.SrcChainName = models.ChainId2Name(r.WrapperTransaction.SrcChainId)
	s.Transaction.SrcChainLogo = models.ChainId2ChainCache(r.WrapperTransaction.SrcChainId).ChainLogo

	s.Transaction.DstChainId = r.WrapperTransaction.DstChainId
	s.Transaction.DstChainName = models.ChainId2Name(r.WrapperTransaction.DstChainId)
	s.Transaction.DstChainExplorer = models.ChainId2ChainCache(r.WrapperTransaction.DstChainId).ChainExplorerUrl
	s.Transaction.DstChainLogo = models.ChainId2ChainCache(r.WrapperTransaction.DstChainId).ChainLogo
	s.Transaction.Time = r.WrapperTransaction.Time
	s.Transaction.From = r.WrapperTransaction.User
	s.Transaction.To = r.WrapperTransaction.DstUser

	if r.SrcTransaction != nil {

		if r.SrcTransaction.SrcTransfer != nil {
			s.SrcTransaction.From = r.SrcTransaction.SrcTransfer.From
			s.SrcTransaction.To = r.SrcTransaction.SrcTransfer.To

			s.Transaction.TokenId = r.SrcTransaction.SrcTransfer.Amount.String()
			token := selectNFTAsset(r.SrcTransaction.SrcTransfer.Asset)
			if token != nil {
				s.SrcTransaction.Asset = token.TokenBasicName
				s.SrcTransaction.AssetHash = token.Hash
				s.SrcTransaction.TokenType = models.GetTokenType(r.SrcTransaction.ChainId, token.Standard)

				s.Transaction.AssetName = token.TokenBasicName
				s.Transaction.NftImage = token.TokenBasic.Meta
			}
		}

		s.SrcTransaction.Hash = r.SrcTransaction.Hash
		s.SrcTransaction.BlockHeight = r.SrcTransaction.Height
		s.SrcTransaction.Time = r.SrcTransaction.Time
		s.SrcTransaction.Status = r.SrcTransaction.State
		s.SrcTransaction.ChainId = r.SrcTransaction.ChainId
		s.SrcTransaction.ChainName = models.ChainId2Name(r.SrcTransaction.ChainId)
		s.SrcTransaction.ChainLogo = models.ChainId2ChainCache(r.SrcTransaction.ChainId).ChainLogo
		s.SrcTransaction.ChainExplorerUrl = models.ChainId2ChainCache(r.SrcTransaction.ChainId).ChainExplorerUrl
		s.SrcTransaction.FeeName = models.ChainId2ChainCache(r.SrcTransaction.ChainId).ChainFeeName
		s.SrcTransaction.FeeLogo = models.ChainId2ChainCache(r.SrcTransaction.ChainId).ChainFeeLogo

		s.SrcTransaction.Fee = models.FormatFee(r.SrcTransaction.ChainId, r.SrcTransaction.Fee)
	}

	if r.DstTransaction != nil {
		if r.DstTransaction.DstTransfer != nil {

			s.DstTransaction.From = r.DstTransaction.DstTransfer.From
			s.DstTransaction.To = r.DstTransaction.DstTransfer.To

			token := selectNFTAsset(r.DstTransaction.DstTransfer.Asset)
			if token != nil {
				s.DstTransaction.Asset = token.TokenBasicName
				s.DstTransaction.AssetHash = token.Hash
				s.DstTransaction.TokenType = models.GetTokenType(r.DstTransaction.ChainId, token.Standard)

				s.Transaction.AssetName = token.TokenBasicName
				s.Transaction.NftImage = token.TokenBasic.Meta
			}
		}

		s.DstTransaction.Hash = r.DstTransaction.Hash
		s.DstTransaction.BlockHeight = r.DstTransaction.Height
		s.DstTransaction.Time = r.DstTransaction.Time
		s.DstTransaction.Status = r.DstTransaction.State
		s.DstTransaction.ChainId = r.DstTransaction.ChainId
		s.DstTransaction.ChainName = models.ChainId2Name(r.DstTransaction.ChainId)
		s.DstTransaction.ChainLogo = models.ChainId2ChainCache(r.DstTransaction.ChainId).ChainLogo
		s.DstTransaction.ChainExplorerUrl = models.ChainId2ChainCache(r.DstTransaction.ChainId).ChainExplorerUrl
		s.DstTransaction.FeeName = models.ChainId2ChainCache(r.DstTransaction.ChainId).ChainFeeName
		s.DstTransaction.FeeLogo = models.ChainId2ChainCache(r.DstTransaction.ChainId).ChainFeeLogo

		s.DstTransaction.Fee = models.FormatFee(r.DstTransaction.ChainId, r.DstTransaction.Fee)
	}

	if r.PolyTransaction != nil {
		s.PolyTransaction.Hash = r.PolyTransaction.Hash
		s.PolyTransaction.BlockHeight = r.PolyTransaction.Height
		s.PolyTransaction.Time = r.PolyTransaction.Time
		s.PolyTransaction.Status = r.PolyTransaction.State
		s.PolyTransaction.From = s.Transaction.From
		s.PolyTransaction.To = s.Transaction.To
	}
	return s
}

type TransactionStateRsp struct {
	Hash       string
	ChainId    uint64
	Blocks     uint64
	NeedBlocks uint64
	Time       uint64
}

type TransactionRsp struct {
	Hash             string
	User             string
	SrcChainId       uint64
	BlockHeight      uint64
	Time             uint64
	DstChainId       uint64
	DstUser          string
	TokenId          string
	ServerId         uint64
	FeeToken         *models.TokenRsp
	FeeAmount        string
	State            uint64
	SrcAsset         *AssetRsp
	DstAsset         *AssetRsp
	TransactionState []*TransactionStateRsp
}

type SrcPolyDstRelation struct {
	SrcHash            string
	WrapperTransaction *models.WrapperTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	SrcTransaction     *models.SrcTransaction     `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash           string
	PolyTransaction    *models.PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
	DstHash            string
	DstTransaction     *models.DstTransaction `gorm:"foreignKey:DstHash;references:Hash"`
	ChainId            uint64                 `gorm:"type:bigint(20);not null"`
	SrcAssetHash       string                 `gorm:"type:varchar(66);not null"`
	SrcAsset           *models.Token          `gorm:"foreignKey:SrcAssetHash,ChainId;references:Hash,ChainId"`
	DstAssetHash       string                 `gorm:"type:varchar(66);not null"`
	DstAsset           *models.Token          `gorm:"foreignKey:DstAssetHash,ChainId;references:Hash,ChainId"`
	FeeTokenHash       string                 `gorm:"size:66;not null"`
	FeeToken           *models.Token          `gorm:"foreignKey:FeeTokenHash,ChainId;references:Hash,ChainId"`
}

func (s *TransactionRsp) instance(
	transaction *SrcPolyDstRelation,
	chainsMap map[uint64]*models.Chain,
) *TransactionRsp {

	s.Hash = transaction.WrapperTransaction.Hash
	s.User = transaction.WrapperTransaction.User
	s.SrcChainId = transaction.WrapperTransaction.SrcChainId
	s.BlockHeight = transaction.WrapperTransaction.BlockHeight
	s.Time = transaction.WrapperTransaction.Time
	s.DstChainId = transaction.WrapperTransaction.DstChainId
	s.ServerId = transaction.WrapperTransaction.ServerId
	s.FeeAmount = transaction.WrapperTransaction.FeeAmount.String()
	s.TokenId = transaction.SrcTransaction.SrcTransfer.Amount.String()
	s.DstUser = transaction.SrcTransaction.SrcTransfer.DstUser
	s.State = transaction.WrapperTransaction.Status

	if transaction.SrcAsset != nil {
		s.SrcAsset = new(AssetRsp).instance(transaction.SrcAsset)
	}
	if transaction.DstAsset != nil {
		s.DstAsset = new(AssetRsp).instance(transaction.DstAsset)
	}
	if transaction.FeeToken != nil {
		s.FeeToken = models.MakeTokenRsp(transaction.FeeToken)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.FeeToken.TokenBasic.Precision)))
		{
			bbb := decimal.NewFromBigInt(&transaction.WrapperTransaction.FeeAmount.Int, 0)
			feeAmount := bbb.Div(precision)
			s.FeeAmount = feeAmount.String()
		}
	}

	srcTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.WrapperTransaction.SrcChainId,
		Blocks:  0,
		Time:    0,
	}
	polyTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: 0,
		Blocks:  0,
		Time:    0,
	}
	dstTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.WrapperTransaction.DstChainId,
		Blocks:  0,
		Time:    0,
	}

	s.TransactionState = append(s.TransactionState, srcTransactionState)
	s.TransactionState = append(s.TransactionState, polyTransactionState)
	s.TransactionState = append(s.TransactionState, dstTransactionState)

	if transaction.SrcTransaction != nil {
		height := transaction.SrcTransaction.Height
		srcTransactionState.Hash = transaction.SrcTransaction.Hash
		srcTransactionState.ChainId = transaction.SrcTransaction.ChainId
		srcTransactionState.Time = transaction.SrcTransaction.Time

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcChain.Height - height
			if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
				srcTransactionState.Blocks = srcTransactionState.NeedBlocks
			}
		}
	}
	if transaction.PolyTransaction != nil {
		polyTransactionState.Hash = transaction.PolyTransaction.Hash
		polyTransactionState.ChainId = transaction.PolyTransaction.ChainId
		polyTransactionState.Time = transaction.PolyTransaction.Time

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyChain.Height - transaction.PolyTransaction.Height
			if polyTransactionState.Blocks > polyTransactionState.NeedBlocks {
				polyTransactionState.Blocks = polyTransactionState.NeedBlocks
			}
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}
	}

	if transaction.DstTransaction != nil {
		dstTransactionState.Hash = transaction.DstTransaction.Hash
		dstTransactionState.ChainId = transaction.DstTransaction.ChainId
		dstTransactionState.Time = transaction.DstTransaction.Time
		dstTransactionState.NeedBlocks = 1

		dstTransactionState.Blocks = transaction.DstTransaction.Height
		dstChain, ok := chainsMap[dstTransactionState.ChainId]
		if ok {
			dstTransactionState.Blocks = dstChain.Height - transaction.DstTransaction.Height
		}
		if dstTransactionState.Blocks > dstTransactionState.NeedBlocks {
			dstTransactionState.Blocks = dstTransactionState.NeedBlocks
		}

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyTransactionState.NeedBlocks
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}
	}
	return s
}

type TransactionsOfAddressReq struct {
	Addresses []string
	PageSize  int
	PageNo    int
}

type TransactionsOfAddressRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionRsp
}

func (s *TransactionsOfAddressRsp) instance(
	pageSize int, pageNo int, totalPage int, totalCount int,
	transactions []*SrcPolyDstRelation,
	chainsMap map[uint64]*models.Chain,
) *TransactionsOfAddressRsp {

	s.PageSize = pageSize
	s.PageNo = pageNo
	s.TotalPage = totalPage
	s.TotalCount = totalCount
	s.Transactions = make([]*TransactionRsp, 0)
	if transactions == nil {
		return s
	}
	for _, transaction := range transactions {
		rsp := new(TransactionRsp).instance(transaction, chainsMap)
		s.Transactions = append(s.Transactions, rsp)
	}
	return s
}

type TransactionOfHashReq struct {
	Hash string
}

//
//type TransactionsOfStateReq struct {
//	State    uint64
//	PageSize int
//	PageNo   int
//}
//
//type TransactionsOfStateRsp struct {
//	PageSize     int
//	PageNo       int
//	TotalPage    int
//	TotalCount   int
//	Transactions []*WrapperTransactionRsp
//}
//
//func MakeTransactionsOfStateRsp(pageSize int, pageNo int, totalPage int, totalCount int,
//	transactions []*models.WrapperTransaction) *WrapperTransactionsRsp {
//
//	transactionsRsp := &WrapperTransactionsRsp{
//		PageSize:   pageSize,
//		PageNo:     pageNo,
//		TotalPage:  totalPage,
//		TotalCount: totalCount,
//	}
//	for _, transaction := range transactions {
//		tx := new(WrapperTransactionRsp).instance(transaction)
//		transactionsRsp.Transactions = append(transactionsRsp.Transactions, tx)
//	}
//	return transactionsRsp
//}

func property2Disable(property int64) int64 {
	if property == 1 {
		return 0
	} else {
		return 1
	}
}
