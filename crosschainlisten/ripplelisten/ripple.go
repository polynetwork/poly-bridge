package ripplelisten

import (
	"encoding/hex"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/rubblelabs/ripple/data"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
	"unicode/utf8"
)

type RippleChainListen struct {
	rippleCfg *conf.ChainListenConfig
	rippleSdk *chainsdk.RippleSdkPro
}

func NewRippleChainListen(cfg *conf.ChainListenConfig) *RippleChainListen {
	rippleListen := &RippleChainListen{}
	rippleListen.rippleCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewRippleSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	rippleListen.rippleSdk = sdk
	return rippleListen
}

func (this *RippleChainListen) GetLatestHeight() (uint64, error) {
	return this.rippleSdk.GetLatestHeight()
}

func (this *RippleChainListen) GetChainListenSlot() uint64 {
	return this.rippleCfg.ListenSlot
}

func (this *RippleChainListen) GetChainId() uint64 {
	return this.rippleCfg.ChainId
}

func (this *RippleChainListen) GetChainName() string {
	return this.rippleCfg.ChainName
}

func (this *RippleChainListen) GetDefer() uint64 {
	return this.rippleCfg.Defer
}

func (this *RippleChainListen) GetBatchSize() uint64 {
	return this.rippleCfg.BatchSize
}

func (this *RippleChainListen) GetXRP() string {
	return this.rippleSdk.GetXRP()
}

func (this *RippleChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.rippleCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	for _, v := range this.rippleCfg.ExtendNodes {
		height, err := this.getExtendLatestHeight(v.Url)
		if err == nil {
			return height, nil
		}
	}
	return this.GetLatestHeight()
}

func (this *RippleChainListen) getExtendLatestHeight(url string) (uint64, error) {
	info := chainsdk.NewRippleInfo(url)
	return info.GetLastHeight()
}

func (this *RippleChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {

	ledger, err := this.rippleSdk.GetLedgerByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	txs := ledger.Ledger.Transactions
	time := uint64(ledger.Ledger.CloseTime.Time().Unix())
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	wrapperDetails := make([]*models.WrapperDetail, 0)

	for _, txData := range txs {
		if txData.MetaData.TransactionResult.Success() { // tx status is success
			if payment, ok := txData.Transaction.(*data.Payment); ok && // payment tx
				payment.Amount.Currency.Machine() == "XRP" { // payment xrp
				//srcTx
				if len(payment.Memos) > 1 {
					continue
				}
				hash := strings.ToLower(payment.Hash.String())
				sequence := uint64(payment.Sequence)
				fromAccount := payment.Account.String()
				toAccount := payment.Destination.String()
				fee := big.NewInt(0)
				paymentfee, err := payment.Fee.NonNative()
				if err == nil {
					if feeAmount, ok := new(big.Int).SetString(paymentfee.String(), 10); ok {
						fee = feeAmount
					}
				}
				nonNative, err := txData.MetaData.DeliveredAmount.NonNative()
				if err != nil {
					logs.Error("chian :%v, height: %v, txhasah: %v, txData.MetaData.DeliveredAmount.NonNative() err: %v", this.GetChainName(), height, payment.Hash.String(), err)
				}
				amount, ok := new(big.Int).SetString(nonNative.String(), 10)
				if !ok {
					logs.Error("chian :%v, height: %v, txhasah: %v, convert amount to big int failed", this.GetChainName(), height, payment.Hash.String())
				}

				if isContract(toAccount, this.rippleCfg.CCMContract) {
					type CrossChainInfo struct {
						DstChain   uint64
						DstAddress string
						DstAsset   string
					}
					crossChainInfo := new(CrossChainInfo)
					if len(payment.Memos) == 0 {
						continue
					}
					memoData, err := hex.DecodeString(payment.Memos[0].Memo.MemoData.String())
					if err != nil {
						logs.Error("HandleNewBlock: DecodeString MemoData error: %v, chain : %v, txHash is: %s", err, this.GetChainName(), hash)
						continue
					}
					if !utf8.ValidString(string(memoData)) {
						logs.Error("HandleNewBlock: memoData ValidString error: %v, chain : %v, txHash is: %s", err, this.GetChainName(), hash)
						continue
					}
					err = json.Unmarshal(memoData, crossChainInfo)
					if err != nil {
						logs.Error("HandleNewBlock: deserialize cross chain info error: %v, chain : %v, txHash is: %s", err, this.GetChainName(), hash)
						continue
					}

					param, _ := json.Marshal(payment.Memos)

					srcTransactions = append(srcTransactions, &models.SrcTransaction{
						Hash:       hash,
						ChainId:    this.GetChainId(),
						Standard:   models.TokenTypeErc20,
						State:      1,
						Time:       time,
						Fee:        models.NewBigInt(fee),
						Height:     height,
						User:       fromAccount,
						DstChainId: crossChainInfo.DstChain,
						Contract:   toAccount,
						Key:        hash,
						Param:      models.Format8190(string(param)),
						SrcTransfer: &models.SrcTransfer{
							TxHash:     hash,
							ChainId:    this.GetChainId(),
							Standard:   models.TokenTypeErc20,
							Time:       time,
							Asset:      this.GetXRP(),
							From:       fromAccount,
							To:         toAccount,
							Amount:     models.NewBigInt(amount),
							DstChainId: crossChainInfo.DstChain,
							DstAsset:   strings.ToLower(crossChainInfo.DstAsset),
							DstUser:    models.FormatString(crossChainInfo.DstAddress),
						},
					})
				} else if isContract(fromAccount, this.rippleCfg.CCMContract) {
					//dstTx
					dstTransactions = append(dstTransactions, &models.DstTransaction{
						Hash:       hash,
						ChainId:    this.GetChainId(),
						Standard:   models.TokenTypeErc20,
						State:      1,
						Time:       time,
						Fee:        models.NewBigInt(fee),
						Height:     height,
						SrcChainId: 0, //
						Contract:   fromAccount,
						PolyHash:   "", //
						Sequence:   sequence,
						DstTransfer: &models.DstTransfer{
							TxHash:   hash,
							ChainId:  this.GetChainId(),
							Standard: models.TokenTypeErc20,
							Time:     time,
							Asset:    this.GetXRP(),
							From:     fromAccount,
							To:       toAccount,
						},
					})
				} else if isContract(toAccount, this.rippleCfg.WrapperContract...) {
					type WrapperInfo struct {
						SrcAddress string
						DstChain   uint64
						DstAddress string
						Amount     string
						LockTxHash string
					}
					//wrapperTx
					if len(payment.Memos) > 0 {
						memoData, err := hex.DecodeString(payment.Memos[0].Memo.MemoData.String())
						if err == nil && len(memoData) > 0 && utf8.ValidString(string(memoData)) {
							wrapperInfo := new(WrapperInfo)
							err = json.Unmarshal(memoData, wrapperInfo)
							if err == nil {
								wrapperDetails = append(wrapperDetails, &models.WrapperDetail{
									WrapperHash:  strings.ToLower(models.FormatString(wrapperInfo.LockTxHash)),
									Hash:         hash,
									User:         models.FormatString(wrapperInfo.SrcAddress),
									SrcChainId:   this.GetChainId(),
									Standard:     models.TokenTypeErc20,
									BlockHeight:  height,
									Time:         time,
									DstChainId:   wrapperInfo.DstChain,
									DstUser:      models.FormatString(wrapperInfo.DstAddress),
									ServerId:     0,
									FeeTokenHash: this.GetXRP(),
									FeeAmount:    models.NewBigInt(amount),
									Status:       basedef.STATE_SOURCE_DONE,
								})
							}
						}
					}
				}
			}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, wrapperDetails, nil, len(srcTransactions), len(dstTransactions), nil
}

func isContract(srcContract string, contracts ...string) bool {
	if len(strings.TrimSpace(srcContract)) == 0 {
		return false
	}
	for _, v := range contracts {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		if srcContract == v {
			return true
		}
	}
	return false
}
