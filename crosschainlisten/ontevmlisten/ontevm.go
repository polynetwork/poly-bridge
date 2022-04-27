package ontevmlisten

import (
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	polycommon "github.com/polynetwork/poly/common"
	"math/big"
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
	ontevmSdk          *chainsdk.OntologySdkPro
	ccmAbiParsed       abi.ABI
	lockproxyAbiParsed abi.ABI
	wrapperAbiParsed   abi.ABI
}

func NewOntevmyChainListen(cfg *conf.ChainListenConfig) *OntevmChainListen {
	ontevmListen := &OntevmChainListen{}
	ontevmListen.ontevmCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewOntologySdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ontevmListen.ontevmSdk = sdk
	ccmAbiParsed, _ := abi.JSON(strings.NewReader(eccm_abi.EthCrossChainManagerABI))
	ontevmListen.ccmAbiParsed = ccmAbiParsed
	lockproxyAbiParsed, _ := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	ontevmListen.lockproxyAbiParsed = lockproxyAbiParsed
	wrapperAbiParsed, _ := abi.JSON(strings.NewReader(wrapper_abi.PolyWrapperABI))
	ontevmListen.wrapperAbiParsed = wrapperAbiParsed
	return ontevmListen
}

func (this *OntevmChainListen) GetLatestHeight() (uint64, error) {
	return this.ontevmSdk.GetCurrentBlockHeight()
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

func (this *OntevmChainListen) isListeningContract(contract string, contracts []string) bool {
	reverseContract := basedef.HexStringReverse(contract)
	for _, item := range contracts {
		if strings.EqualFold(reverseContract, item) {
			return true
		}
	}
	return false
}
func (this *OntevmChainListen) isListeningContract1(contract0 string, contract1 string) bool {
	reverseContract := basedef.HexStringReverse(contract0)
	if strings.EqualFold(reverseContract, contract1) {
		return true
	}
	return false
}

func (this *OntevmChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	block, err := this.ontevmSdk.GetBlockByHeight(uint32(height))
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	tt := uint64(block.Header.Timestamp)
	events, err := this.ontevmSdk.GetSmartContractEventByBlock(uint32(height))
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, event := range events {
		for _, notify := range event.Notify {
			if this.isListeningContract(notify.ContractAddress, this.ontevmCfg.WrapperContract) {
				states, ok := notify.States.(string)
				if !ok {
					continue
				}

				if len(storageLog.Topics) == 0 {
					continue
				}
				for _, topic := range storageLog.Topics {
					switch topic {
					case this.wrapperAbiParsed.Events["PolyWrapperLock"].ID:
						logs.Info("(wrapper) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
						var evt wrapper_abi.PolyWrapperPolyWrapperLock
						err = this.wrapperAbiParsed.UnpackIntoInterface(&event, "PolyWrapperLock", storageLog.Data)
						if err != nil {
							continue
						}
						wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
							Hash:         evt.Raw.TxHash.String()[2:],
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
			} else if this.isListeningContract1(notify.ContractAddress, this.ontevmCfg.CCMContract) {
				states, ok := notify.States.(string)
				if !ok {
					continue
				}
				var data []byte
				data, err = hexutil.Decode(states)
				if err != nil {
					continue
				}
				source := polycommon.NewZeroCopySource(data)
				var storageLog StorageLog
				err = storageLog.Deserialization(source)
				if err != nil {
					continue
				}
				if len(storageLog.Topics) == 0 {
					continue
				}
				for _, topic := range storageLog.Topics {
					switch topic {
					case this.wrapperAbiParsed.Events["CrossChainEvent"].ID:
						logs.Info("(lock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
						srcTransfer := &models.SrcTransfer{}
						for _, notifyNew := range event.Notify {
							statesNew := notifyNew.States.([]interface{})
							method, ok := statesNew[0].(string)
							if !ok {
								continue
							}
							method = this.parseOntolofyMethod(method)
							if method == _ont_lock {
								if len(statesNew) < 7 {
									continue
								}
								srcTransfer.ChainId = this.GetChainId()
								srcTransfer.TxHash = event.TxHash
								srcTransfer.Time = tt
								srcTransfer.From = statesNew[2].(string)
								srcTransfer.To = states[5].(string)
								srcTransfer.Asset = basedef.HexStringReverse(statesNew[1].(string))
								if len(srcTransfer.Asset) < 20 {
									continue
								}
								amount, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[6].(string)), 16)
								srcTransfer.Amount = models.NewBigInt(amount)
								toChain, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[3].(string)), 16)
								srcTransfer.DstChainId = toChain.Uint64()
								srcTransfer.DstAsset = statesNew[4].(string)
								srcTransfer.DstUser = statesNew[5].(string)
								if len(srcTransfer.From) > basedef.ADDRESS_LENGTH {
									srcTransfer.From = ""
								}
								if len(srcTransfer.To) > basedef.ADDRESS_LENGTH {
									srcTransfer.To = ""
								}
								if len(srcTransfer.DstUser) > basedef.ADDRESS_LENGTH {
									srcTransfer.DstUser = ""
								}
								break
							}
						}
						srcTransaction := &models.SrcTransaction{}
						srcTransaction.ChainId = this.GetChainId()
						srcTransaction.Hash = event.TxHash
						srcTransaction.State = uint64(event.State)
						srcTransaction.Fee = models.NewBigIntFromInt(int64(event.GasConsumed))
						srcTransaction.Time = tt
						srcTransaction.Height = height
						srcTransaction.User = srcTransfer.From
						srcTransaction.DstChainId = uint64(states[2].(float64))
						srcTransaction.Contract = basedef.HexStringReverse(states[5].(string))
						srcTransaction.Key = states[4].(string)
						srcTransaction.Param = states[6].(string)
						srcTransaction.SrcTransfer = srcTransfer
						srcTransactions = append(srcTransactions, srcTransaction)
					case _ont_crosschainunlock:
						logs.Info("(unlock) to chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
						if len(states) < 6 {
							continue
						}
						dstTransfer := &models.DstTransfer{}
						for _, notifyNew := range event.Notify {
							statesNew := notifyNew.States.([]interface{})
							method, ok := statesNew[0].(string)
							if !ok {
								continue
							}
							method = this.parseOntolofyMethod(method)
							if method == _ont_unlock {
								if len(statesNew) < 4 {
									continue
								}
								dstTransfer.ChainId = this.GetChainId()
								dstTransfer.TxHash = event.TxHash
								dstTransfer.Time = tt
								dstTransfer.From = states[5].(string)
								dstTransfer.To = statesNew[2].(string)
								dstTransfer.Asset = basedef.HexStringReverse(statesNew[1].(string))
								if len(dstTransfer.Asset) < 20 {
									continue
								}
								amount, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[3].(string)), 16)
								dstTransfer.Amount = models.NewBigInt(amount)
								break
							}
						}
						dstTransaction := &models.DstTransaction{}
						dstTransaction.ChainId = this.GetChainId()
						dstTransaction.Hash = event.TxHash
						dstTransaction.State = uint64(event.State)
						dstTransaction.Fee = models.NewBigIntFromInt(int64(event.GasConsumed))
						dstTransaction.Time = tt
						dstTransaction.Height = height
						dstTransaction.SrcChainId = uint64(states[3].(float64))
						dstTransaction.Contract = basedef.HexStringReverse(states[5].(string))
						dstTransaction.PolyHash = basedef.HexStringReverse(states[1].(string))
						dstTransaction.DstTransfer = dstTransfer
						dstTransactions = append(dstTransactions, dstTransaction)
					default:
						logs.Warn("ignore method: %s", contractMethod)
					}
				}
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

func deserializeStorageLog(states string) {

	var data []byte
	data, err = hexutil.Decode(states)
	if err != nil {
		continue
	}
	source := polycommon.NewZeroCopySource(data)
	var storageLog StorageLog
	err = storageLog.Deserialization(source)
	if err != nil {
		continue
	}
}
