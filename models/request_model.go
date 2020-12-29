package models

type PolySwapResp struct {
	Version string
	URL     string
}

type TokenBasicReq struct {
	Name string
}

type TokenBasicRsp struct {
	Name     string
	CmcName  string
	CmcPrice int64
	CmcInd   uint64
	BinName  string
	BinPrice int64
	BinInd   uint64
	AvgPrice int64
	AvgInd   uint64
	Time     uint64
	Tokens   []*TokenRsp
}

func MakeTokenBasicRsp(tokenBasic *TokenBasic) *TokenBasicRsp {
	tokenBasicRsp := &TokenBasicRsp{
		Name:     tokenBasic.Name,
		CmcName:  tokenBasic.CmcName,
		CmcPrice: tokenBasic.CmcPrice,
		CmcInd:   tokenBasic.CmcInd,
		BinName:  tokenBasic.BinName,
		BinPrice: tokenBasic.BinPrice,
		BinInd:   tokenBasic.BinInd,
		AvgPrice: tokenBasic.AvgPrice,
		AvgInd:   tokenBasic.AvgInd,
		Time:     tokenBasic.Time,
		Tokens:   nil,
	}
	if tokenBasic.Tokens != nil {
		for _, token := range tokenBasic.Tokens {
			tokenBasicRsp.Tokens = append(tokenBasicRsp.Tokens, MakeTokenRsp(token))
		}
	}
	return tokenBasicRsp
}

type TokenReq struct {
	Hash string
}

type TokenRsp struct {
	Hash           string
	ChainId        uint64
	Name           string
	TokenBasicName string
	TokenBasic     *TokenBasicRsp
	TokenMaps      []*TokenMapRsp
}

func MakeTokenRsp(token *Token) *TokenRsp {
	tokenRsp := &TokenRsp{
		Hash:           token.Hash,
		ChainId:        token.ChainId,
		Name:           token.Name,
		TokenBasicName: token.TokenBasicName,
	}
	if token.TokenBasic != nil {
		tokenRsp.TokenBasic = MakeTokenBasicRsp(token.TokenBasic)
	}
	if token.TokenMaps != nil {
		for _, tokenmap := range token.TokenMaps {
			tokenRsp.TokenMaps = append(tokenRsp.TokenMaps, MakeTokenMapRsp(tokenmap))
		}
	}
	return tokenRsp
}

type TokensReq struct {
	ChainId uint64
}

type TokensRsp struct {
	TotalCount uint64
	Tokens     []*TokenRsp
}

func MakeTokensRsp(tokens []*Token) *TokensRsp {
	tokensRsp := &TokensRsp{
		TotalCount: uint64(len(tokens)),
	}
	for _, token := range tokens {
		tokensRsp.Tokens = append(tokensRsp.Tokens, MakeTokenRsp(token))
	}
	return tokensRsp
}

type TokenMapReq struct {
	Hash string
}

type TokenMapRsp struct {
	SrcTokenHash string
	SrcToken     *TokenRsp
	DstTokenHash string
	DstToken     *TokenRsp
}

func MakeTokenMapRsp(tokenMap *TokenMap) *TokenMapRsp {
	tokenMapRsp := &TokenMapRsp{
		SrcTokenHash: tokenMap.SrcTokenHash,
		DstTokenHash: tokenMap.DstTokenHash,
	}
	if tokenMap.SrcToken != nil {
		tokenMapRsp.SrcToken = MakeTokenRsp(tokenMap.SrcToken)
	}
	if tokenMap.DstToken != nil {
		tokenMapRsp.DstToken = MakeTokenRsp(tokenMap.DstToken)
	}
	return tokenMapRsp
}

type GetFeeReq struct {
	ChainId uint64
	Hash    string
}

type GetFeeRsp struct {
	ChainId uint64
	Hash    string
	Amount  float64
}

func MakeGetFeeRsp(chainId uint64, hash string, amount float64) *GetFeeRsp {
	getFeeRsp := &GetFeeRsp{
		ChainId: chainId,
		Hash:    hash,
		Amount:  amount,
	}
	return getFeeRsp
}

type CheckFeeReq struct {
	Hash string
}

type CheckFeeRsp struct {
	HasPay bool
	Amount float64
}

func MakeCheckFeeRsp(hashPay bool, amount float64) *CheckFeeRsp {
	checkFeeRsp := &CheckFeeRsp{
		HasPay: hashPay,
		Amount: amount,
	}
	return checkFeeRsp
}

type TransactionReq struct {
	Hash string
}

type TransactionRsp struct {
	Hash           string
	User           string
	SrcChainId        uint64
	BlockHeight           uint64
	Time uint64
	DstChainId     uint64
	FeeTokenHash      string
	FeeAmount   uint64
}

func MakeTransactionRsp(transaction *WrapperTransaction) *TransactionRsp {
	transactionRsp := &TransactionRsp{
		Hash:           transaction.Hash,
		User:        transaction.User,
		SrcChainId:           transaction.SrcChainId,
		BlockHeight:        transaction.BlockHeight,
		Time:           transaction.Time,
		DstChainId:        transaction.DstChainId,
		FeeTokenHash:           transaction.FeeTokenHash,
		FeeAmount:           transaction.FeeAmount.Uint64(),
	}
	return transactionRsp
}

type TransactionsReq struct {
	PageSize int
	PageNo int
}

type TransactionsRsp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	Transactions     []*TransactionRsp
}

func MakeTransactionsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *TransactionsRsp {
	transactionsRsp := &TransactionsRsp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeTransactionRsp(transaction))
	}
	return transactionsRsp
}

type TransactionsOfUserReq struct {
	User string
	PageSize int
	PageNo int
}

type TransactionsOfUserRsp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	Transactions     []*TransactionRsp
}

func MakeTransactionsOfUserRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *TransactionsRsp {
	transactionsRsp := &TransactionsRsp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeTransactionRsp(transaction))
	}
	return transactionsRsp
}





type AddressReq struct {
	ChainId     uint64
	AddressHash string
}

type AddressRsp struct {
	AddressHash      string
	Address          string
	ChainId        uint64
}

func MakeAddressRsp(addressHash string, chainId uint64, address string) *AddressRsp {
	addressRsp := &AddressRsp{
		AddressHash:          addressHash,
		Address:        address,
		ChainId:           chainId,
	}
	return addressRsp
}
