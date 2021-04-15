package controllers

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
)

type PolyBridgeInfoResp struct {
	Version string
	URL     string
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
	Asset *AssetRsp
	Items []*Item
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

func (i *Item) instance(assetName string, tokenId *big.Int, profile *models.NFTProfile) *Item {
	i.AssetName = assetName
	i.TokenId = tokenId.String()
	if profile == nil {
		return i
	}
	i.Name = profile.Name
	i.Url = profile.Url
	i.Image = profile.Image
	i.Desc = profile.Description
	i.Meta = profile.Text
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
	ChainId uint64
	Size    int
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
	s.Items = make([]*Item, 0)
	if items == nil {
		return s
	}
	for _, v := range items {
		s.Items = append(s.Items, v)
	}
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

type WrapperTransactionsReq struct {
	PageSize int
	PageNo   int
}

type WrapperTransactionsRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*WrapperTransactionRsp
}

func (s *WrapperTransactionsRsp) instance(
	pageSize, pageNo, totalPage, totalCount int,
	transactions []*models.WrapperTransaction,
) *WrapperTransactionsRsp {

	s.PageSize = pageSize
	s.PageNo = pageNo
	s.TotalCount = totalCount
	s.TotalPage = totalPage
	s.Transactions = make([]*WrapperTransactionRsp, 0)

	if transactions == nil {
		return s
	}

	for _, v := range transactions {
		tx := new(WrapperTransactionRsp).instance(v)
		s.Transactions = append(s.Transactions, tx)
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
	if transaction.SrcTransaction != nil {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    transaction.SrcTransaction.Hash,
			ChainId: transaction.SrcTransaction.ChainId,
			Blocks:  transaction.SrcTransaction.Height,
			Time:    transaction.SrcTransaction.Time,
		})
	} else {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: transaction.WrapperTransaction.SrcChainId,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction.PolyTransaction != nil {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    transaction.PolyTransaction.Hash,
			ChainId: transaction.PolyTransaction.ChainId,
			Blocks:  transaction.PolyTransaction.Height,
			Time:    transaction.PolyTransaction.Time,
		})
	} else {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: 0,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction.DstTransaction != nil {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    transaction.DstTransaction.Hash,
			ChainId: transaction.DstTransaction.ChainId,
			Blocks:  transaction.DstTransaction.Height,
			Time:    transaction.DstTransaction.Time,
		})
	} else {
		s.TransactionState = append(s.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: transaction.WrapperTransaction.DstChainId,
			Blocks:  0,
			Time:    0,
		})
	}
	for _, state := range s.TransactionState {
		chain, ok := chainsMap[state.ChainId]
		if ok {
			if state.ChainId == transaction.WrapperTransaction.DstChainId {
				state.NeedBlocks = 1
			} else {
				state.NeedBlocks = chain.BackwardBlockNumber
			}
			if state.Blocks <= 1 {
				continue
			}
			state.Blocks = chain.Height - state.Blocks
			if state.Blocks > state.NeedBlocks {
				state.Blocks = state.NeedBlocks
			}
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
