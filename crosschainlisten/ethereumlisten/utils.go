package ethereumlisten

import (
	"encoding/hex"
	"strings"

	"poly-bridge/basedef"
	"poly-bridge/go_abi/eccm_abi"
	nftlp "poly-bridge/go_abi/nft_lock_proxy_abi"
	nftwp "poly-bridge/go_abi/nft_wrap_abi"
	"poly-bridge/models"
)

func assembleSrcTransaction(
	eccmLockEvent *models.ECCMLockEvent,
	proxyLockEvents []*models.ProxyLockEvent,
	chainID uint64,
	tt uint64,
) *models.SrcTransaction {
	srcTransaction := &models.SrcTransaction{}
	srcTransaction.ChainId = chainID
	srcTransaction.Hash = eccmLockEvent.TxHash
	srcTransaction.State = 1
	srcTransaction.Fee = models.NewBigIntFromInt(int64(eccmLockEvent.Fee))
	if tt > 0 {
		srcTransaction.Time = tt
	}
	srcTransaction.Height = eccmLockEvent.Height
	srcTransaction.User = eccmLockEvent.User
	srcTransaction.DstChainId = uint64(eccmLockEvent.Tchain)
	srcTransaction.Contract = eccmLockEvent.Contract
	srcTransaction.Key = eccmLockEvent.Txid
	srcTransaction.Param = hex.EncodeToString(eccmLockEvent.Value)
	for _, v := range proxyLockEvents {
		if v.TxHash == eccmLockEvent.TxHash {
			toAssetHash := v.ToAssetHash
			srcTransfer := &models.SrcTransfer{}
			if tt > 0 {
				srcTransfer.Time = tt
			}
			srcTransfer.ChainId = chainID
			srcTransfer.TxHash = eccmLockEvent.TxHash
			srcTransfer.From = eccmLockEvent.User
			srcTransfer.To = eccmLockEvent.Contract
			srcTransfer.Asset = v.FromAssetHash
			srcTransfer.Amount = models.NewBigInt(v.Amount)
			srcTransfer.DstChainId = uint64(v.ToChainId)
			srcTransfer.DstAsset = toAssetHash
			srcTransfer.DstUser = v.ToAddress
			srcTransaction.SrcTransfer = srcTransfer
			break
		}
	}
	return srcTransaction
}

func assembleDstTransaction(
	eccmUnlockEvent *models.ECCMUnlockEvent,
	proxyUnlockEvents []*models.ProxyUnlockEvent,
	chainID uint64,
	tt uint64,
) *models.DstTransaction {

	dstTransaction := &models.DstTransaction{}
	dstTransaction.ChainId = chainID
	dstTransaction.Hash = eccmUnlockEvent.TxHash
	dstTransaction.State = 1
	dstTransaction.Fee = models.NewBigIntFromInt(int64(eccmUnlockEvent.Fee))
	if tt > 0 {
		dstTransaction.Time = tt
	}
	dstTransaction.Height = eccmUnlockEvent.Height
	dstTransaction.SrcChainId = uint64(eccmUnlockEvent.FChainId)
	dstTransaction.Contract = eccmUnlockEvent.Contract
	dstTransaction.PolyHash = eccmUnlockEvent.RTxHash
	for _, v := range proxyUnlockEvents {
		if v.TxHash == eccmUnlockEvent.TxHash {
			dstTransfer := &models.DstTransfer{}
			dstTransfer.TxHash = eccmUnlockEvent.TxHash
			if tt > 0 {
				dstTransfer.Time = tt
			}
			dstTransfer.ChainId = chainID
			dstTransfer.From = eccmUnlockEvent.Contract
			dstTransfer.To = v.ToAddress
			dstTransfer.Asset = v.ToAssetHash
			dstTransfer.Amount = models.NewBigInt(v.Amount)
			dstTransaction.DstTransfer = dstTransfer
			break
		}
	}

	return dstTransaction
}

func wrapLockEvent2WrapTx(evt *nftwp.PolyNFTWrapperPolyWrapperLock) *models.WrapperTransaction {
	return &models.WrapperTransaction{
		Hash:         evt.Raw.TxHash.String()[2:],
		User:         strings.ToLower(evt.Sender.String()[2:]),
		DstChainId:   evt.ToChainId,
		DstUser:      strings.ToLower(evt.ToAddress.String()[2:]),
		FeeTokenHash: strings.ToLower(evt.FeeToken.String()[2:]),
		FeeAmount:    models.NewBigInt(evt.Fee),
		ServerId:     evt.Id.Uint64(),
		BlockHeight:  evt.Raw.BlockNumber,
		Standard:     models.TokenTypeErc721,
	}
}

func wrapSpeedUpEvent2WrapTx(evt *nftwp.PolyNFTWrapperPolyWrapperSpeedUp) *models.WrapperTransaction {
	return &models.WrapperTransaction{
		Hash:         evt.TxHash.String(),
		User:         evt.Sender.String(),
		FeeTokenHash: evt.FeeToken.String(),
		FeeAmount:    models.NewBigInt(evt.Efee),
		Standard:     models.TokenTypeErc721,
	}
}

func crossChainEvent2ProxyLockEvent(
	evt *eccm_abi.EthCrossChainManagerImplemetationCrossChainEvent,
	fee uint64,
) *models.ECCMLockEvent {

	return &models.ECCMLockEvent{
		Method:   _eth_crosschainlock,
		Txid:     hex.EncodeToString(evt.TxId),
		TxHash:   evt.Raw.TxHash.String()[2:],
		User:     strings.ToLower(evt.Sender.String()[2:]),
		Tchain:   uint32(evt.ToChainId),
		Contract: strings.ToLower(evt.ProxyOrAssetContract.String()[2:]),
		Value:    evt.Rawdata,
		Height:   evt.Raw.BlockNumber,
		Fee:      fee,
	}
}

func convertLockProxyEvent(evt *nftlp.PolyNFTLockProxyLockEvent) *models.ProxyLockEvent {
	return &models.ProxyLockEvent{
		BlockNumber:   evt.Raw.BlockNumber,
		Method:        _eth_lock,
		TxHash:        evt.Raw.TxHash.String()[2:],
		FromAddress:   evt.FromAddress.String()[2:],
		FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
		ToChainId:     uint32(evt.ToChainId),
		ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
		ToAddress:     hex.EncodeToString(evt.ToAddress),
		Amount:        evt.TokenId,
	}
}

func verifyAndExecuteEvent2ProxyUnlockEvent(
	evt *eccm_abi.EthCrossChainManagerImplemetationVerifyHeaderAndExecuteTxEvent,
	fee uint64,
) *models.ECCMUnlockEvent {

	return &models.ECCMUnlockEvent{
		Method:   _eth_crosschainunlock,
		TxHash:   evt.Raw.TxHash.String()[2:],
		RTxHash:  basedef.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
		Contract: hex.EncodeToString(evt.ToContract),
		FChainId: uint32(evt.FromChainID),
		Height:   evt.Raw.BlockNumber,
		Fee:      fee,
	}
}

func convertUnlockProxyEvent(evt *nftlp.PolyNFTLockProxyUnlockEvent) *models.ProxyUnlockEvent {
	return &models.ProxyUnlockEvent{
		BlockNumber: evt.Raw.BlockNumber,
		Method:      _eth_unlock,
		TxHash:      evt.Raw.TxHash.String()[2:],
		ToAssetHash: strings.ToLower(evt.ToAssetHash.String()[2:]),
		ToAddress:   strings.ToLower(evt.ToAddress.String()[2:]),
		Amount:      evt.TokenId,
	}
}

func addERC721TokenStandard(
	wptxs []*models.WrapperTransaction,
	srcTxs []*models.SrcTransaction,
	dstTxs []*models.DstTransaction) {

	for _, v := range wptxs {
		v.Standard = models.TokenTypeErc721
	}
	for _, v := range srcTxs {
		v.Standard = models.TokenTypeErc721
		if v.SrcTransfer != nil {
			v.SrcTransfer.Standard = models.TokenTypeErc721
		}
	}
	for _, v := range dstTxs {
		v.Standard = models.TokenTypeErc721
		if v.DstTransfer != nil {
			v.DstTransfer.Standard = models.TokenTypeErc721
		}
	}
}
