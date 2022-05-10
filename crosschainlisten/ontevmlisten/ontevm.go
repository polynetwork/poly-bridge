package ontevmlisten

import (
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ontcommon "github.com/ontio/ontology-go-sdk/common"
	polycommon "github.com/polynetwork/poly/common"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"poly-bridge/go_abi/lock_proxy_abi"
	"poly-bridge/go_abi/wrapper_abi"
	"poly-bridge/models"
	"strings"
)

const (
	_PolyWrapperLock               = "PolyWrapperLock"
	_CrossChainEvent               = "CrossChainEvent"
	_VerifyHeaderAndExecuteTxEvent = "VerifyHeaderAndExecuteTxEvent"
	_LockEvent                     = "LockEvent"
	_UnlockEvent                   = "UnlockEvent"
)

type OntevmChainListen struct {
	ontevmCfg          *conf.ChainListenConfig
	ontSdk             *chainsdk.OntologySdkPro
	ccmAbiParsed       abi.ABI
	lockproxyAbiParsed abi.ABI
	wrapperAbiParsed   abi.ABI
}

func NewOntevmChainListen(cfg *conf.ChainListenConfig) *OntevmChainListen {
	ontevmListen := &OntevmChainListen{}
	ontevmListen.ontevmCfg = cfg
	//urls use ont url,listen ontevm must use ontsdk because filterlog
	//use ExtendNodes only here,others use GetNodesUrl
	urls := cfg.GetExtendNodesUrl()

	sdk := chainsdk.NewOntologySdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ontevmListen.ontSdk = sdk
	ccmAbiParsed, _ := abi.JSON(strings.NewReader(eccm_abi.EthCrossChainManagerABI))
	ontevmListen.ccmAbiParsed = ccmAbiParsed
	lockproxyAbiParsed, _ := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	ontevmListen.lockproxyAbiParsed = lockproxyAbiParsed
	wrapperAbiParsed, _ := abi.JSON(strings.NewReader(wrapper_abi.PolyWrapperABI))
	ontevmListen.wrapperAbiParsed = wrapperAbiParsed
	return ontevmListen
}

func (this *OntevmChainListen) GetLatestHeight() (uint64, error) {
	return this.ontSdk.GetCurrentBlockHeight()
}

func (this *OntevmChainListen) GetChainListenSlot() uint64 {
	return this.ontevmCfg.ListenSlot
}

func (this *OntevmChainListen) GetChainId() uint64 {
	return this.ontevmCfg.ChainId
}

func (this *OntevmChainListen) GetChainName() string {
	return this.ontevmCfg.ChainName
}

func (this *OntevmChainListen) parseOntolofyMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func (this *OntevmChainListen) GetDefer() uint64 {
	return this.ontevmCfg.Defer
}

func (this *OntevmChainListen) GetBatchSize() uint64 {
	return this.ontevmCfg.BatchSize
}

func (this *OntevmChainListen) isListeningContract(contract string, contracts ...string) bool {
	reverseContract := basedef.HexStringReverse(contract)
	for _, item := range contracts {
		if strings.EqualFold(reverseContract, item) {
			return true
		}
	}
	return false
}

func (this *OntevmChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	block, err := this.ontSdk.GetBlockByHeight(uint32(height))
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	tt := uint64(block.Header.Timestamp)
	events, err := this.ontSdk.GetSmartContractEventByBlock(uint32(height))
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransfers := make([]*models.SrcTransfer, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	dstTransfers := make([]*models.DstTransfer, 0)
	for _, event := range events {
		for _, notify := range event.Notify {
			if this.isListeningContract(notify.ContractAddress, this.ontevmCfg.WrapperContract...) {
				storageLog, err := deserializeStorageLog(notify)
				if err != nil {
					continue
				}
				for _, topic := range storageLog.Topics {
					switch topic {
					case this.wrapperAbiParsed.Events["PolyWrapperLock"].ID:
						logs.Info("(wrapper) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, basedef.HexStringReverse(event.TxHash))
						var evt wrapper_abi.PolyWrapperPolyWrapperLock
						err = this.wrapperAbiParsed.UnpackIntoInterface(&evt, "PolyWrapperLock", storageLog.Data)
						if err != nil {
							continue
						}
						wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
							Hash:         basedef.HexStringReverse(event.TxHash),
							User:         models.FormatString(strings.ToLower(evt.Sender.String()[2:])),
							DstChainId:   evt.ToChainId,
							DstUser:      models.FormatString(hex.EncodeToString(evt.ToAddress)),
							FeeTokenHash: models.FormatString(strings.ToLower(evt.FromAsset.String()[2:])),
							FeeAmount:    models.NewBigInt(evt.Fee),
							ServerId:     evt.Id.Uint64(),
							BlockHeight:  evt.Raw.BlockNumber,
							Status:       basedef.STATE_SOURCE_DONE,
							Time:         tt,
							SrcChainId:   this.GetChainId(),
						})
					}
				}
			} else if this.isListeningContract(notify.ContractAddress, this.ontevmCfg.CCMContract) {
				storageLog, err := deserializeStorageLog(notify)
				if err != nil {
					continue
				}
				for _, topic := range storageLog.Topics {
					switch topic {
					case this.ccmAbiParsed.Events["CrossChainEvent"].ID:
						logs.Info("(ccm lock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, basedef.HexStringReverse(event.TxHash))
						var evt eccm_abi.EthCrossChainManagerCrossChainEvent
						err = this.ccmAbiParsed.UnpackIntoInterface(&evt, "CrossChainEvent", storageLog.Data)
						if err != nil {
							continue
						}
						srcTransactions = append(srcTransactions, &models.SrcTransaction{
							Hash:       basedef.HexStringReverse(event.TxHash),
							ChainId:    this.GetChainId(),
							State:      1,
							Time:       tt,
							Fee:        models.NewBigIntFromInt(int64(event.GasConsumed)),
							Height:     height,
							DstChainId: evt.ToChainId,
							Contract:   models.FormatString(evt.ProxyOrAssetContract.String()[2:]),
							Key:        hex.EncodeToString(evt.TxId),
							Param:      hex.EncodeToString(evt.Rawdata),
						})
					case this.ccmAbiParsed.Events["VerifyHeaderAndExecuteTxEvent"].ID:
						logs.Info("(ccm unlock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, basedef.HexStringReverse(event.TxHash))
						var evt eccm_abi.EthCrossChainManagerVerifyHeaderAndExecuteTxEvent
						err = this.ccmAbiParsed.UnpackIntoInterface(&evt, "VerifyHeaderAndExecuteTxEvent", storageLog.Data)
						if err != nil {
							continue
						}
						dstTransactions = append(dstTransactions, &models.DstTransaction{
							Hash:       basedef.HexStringReverse(event.TxHash),
							ChainId:    this.GetChainId(),
							State:      1,
							Time:       tt,
							Fee:        models.NewBigIntFromInt(int64(event.GasConsumed)),
							Height:     height,
							SrcChainId: evt.FromChainID,
							Contract:   models.FormatString(basedef.HexStringReverse(hex.EncodeToString(evt.ToContract))),
							PolyHash:   basedef.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
						})
					}
				}
			} else if this.isListeningContract(notify.ContractAddress, this.ontevmCfg.ProxyContract...) {
				storageLog, err := deserializeStorageLog(notify)
				if err != nil {
					continue
				}
				for _, topic := range storageLog.Topics {
					switch topic {
					case this.lockproxyAbiParsed.Events["LockEvent"].ID:
						logs.Info("(lockproxy lock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, basedef.HexStringReverse(event.TxHash))
						var evt lock_proxy_abi.LockProxyLockEvent
						err = this.lockproxyAbiParsed.UnpackIntoInterface(&evt, "LockEvent", storageLog.Data)
						if err != nil {
							continue
						}
						srcTransfers = append(srcTransfers, &models.SrcTransfer{
							TxHash:     basedef.HexStringReverse(event.TxHash),
							ChainId:    this.GetChainId(),
							Standard:   models.TokenTypeErc20,
							Time:       tt,
							Asset:      models.FormatString(strings.ToLower(evt.FromAssetHash.String()[2:])),
							Amount:     models.NewBigInt(evt.Amount),
							DstChainId: evt.ToChainId,
							DstAsset:   models.FormatString(hex.EncodeToString(evt.ToAssetHash)),
							DstUser:    models.FormatString(hex.EncodeToString(evt.ToAddress)),
							From:       models.FormatString(evt.FromAddress.String()[2:]),
						})
					case this.lockproxyAbiParsed.Events["UnlockEvent"].ID:
						logs.Info("(lockproxy unlock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, basedef.HexStringReverse(event.TxHash))
						var evt lock_proxy_abi.LockProxyUnlockEvent
						err = this.lockproxyAbiParsed.UnpackIntoInterface(&evt, "UnlockEvent", storageLog.Data)
						if err != nil {
							continue
						}
						dstTransfers = append(dstTransfers, &models.DstTransfer{
							TxHash:   basedef.HexStringReverse(event.TxHash),
							ChainId:  this.GetChainId(),
							Standard: models.TokenTypeErc20,
							Time:     tt,
							Asset:    models.FormatString(strings.ToLower(evt.ToAssetHash.String()[2:])),
							To:       models.FormatString(strings.ToLower(evt.ToAddress.String()[2:])),
							Amount:   models.NewBigInt(evt.Amount),
						})
					}
				}
			}
		}
	}
	for _, srcTransaction := range srcTransactions {
		for _, srcTransfer := range srcTransfers {
			if srcTransaction.Hash == srcTransfer.TxHash {
				srcTransaction.User = srcTransfer.From
				srcTransfer.To = models.FormatString(srcTransaction.Contract)
				srcTransaction.Standard = srcTransfer.Standard
				srcTransaction.SrcTransfer = srcTransfer
			}
		}
	}
	for _, dstTransaction := range dstTransactions {
		for _, dstTransfer := range dstTransfers {
			if dstTransaction.Hash == dstTransfer.TxHash {
				dstTransfer.From = dstTransaction.Contract
				dstTransaction.Standard = dstTransfer.Standard
				dstTransaction.DstTransfer = dstTransfer
			}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, len(srcTransactions), len(dstTransactions), nil
}

func (this *OntevmChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.ontevmCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}

type StorageLog struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

func (self *StorageLog) Deserialization(source *polycommon.ZeroCopySource) error {
	address, eof := source.NextAddress()
	if eof {
		return fmt.Errorf("StorageLog.address eof")
	}
	self.Address = common.Address(address)
	l, eof := source.NextUint32()
	if eof {
		return fmt.Errorf("StorageLog.l eof")
	}
	self.Topics = make([]common.Hash, 0, l)
	for i := uint32(0); i < l; i++ {
		h, _ := source.NextHash()
		if eof {
			return fmt.Errorf("StorageLog.h eof")
		}
		self.Topics = append(self.Topics, common.Hash(h))
	}
	data, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("StorageLog.Data eof")
	}
	self.Data = data
	return nil
}

func deserializeStorageLog(notify *ontcommon.NotifyEventInfo) (storageLog StorageLog, err error) {
	states, ok := notify.States.(string)
	if !ok {
		err = fmt.Errorf("err States.(string)")
		return
	}
	var data []byte
	data, err = hexutil.Decode(states)
	if err != nil {
		return
	}
	source := polycommon.NewZeroCopySource(data)
	err = storageLog.Deserialization(source)
	if err != nil {
		return
	}
	if len(storageLog.Topics) == 0 {
		err = fmt.Errorf("err storageLog.Topics is 0")
		return
	}
	return
}
