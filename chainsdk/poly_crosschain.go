package chainsdk

import (
	"bytes"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/beego/beego/v2/core/logs"
	"github.com/btcsuite/btcd/btcec"
	ecm "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/sm2"
	polysdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/common"
	vconfig "github.com/polynetwork/poly/consensus/vbft/config"
	"github.com/polynetwork/poly/core/types"
)

//const (
//	block2wait uint64 = 1
//)

func NewPolySdkAndSetChainID(url string) (*PolySDK, error) {
	s := NewPolySDK(url)
	blk, err := s.GetBlockByHeight(0)
	if err != nil {
		return nil, err
	}
	hdr := blk.Header
	s.sdk.SetChainId(hdr.ChainID)
	return s, nil
}

// client的账户列表就是poly共识节点账户列表，可以通过注册和取消账户的方式实现bookKeeper的变更
func (s *PolySDK) RegNode(node *polysdk.Account, signer *polysdk.Account, validators []*polysdk.Account) error {
	peer := vconfig.PubkeyID(node.PublicKey)

	if err := s.RegisterCandidate(peer, signer); err != nil {
		return err
	} else {
		log.Info("register %s success!", peer)
	}
	if err := s.ApproveCandidate(peer, validators); err != nil {
		return err
	} else {
		log.Info("approve %s success", peer)
	}
	return s.CommitPolyDpos(validators)
}

func (s *PolySDK) SyncGenesisBlock(
	selfChainID uint64,
	validators []*polysdk.Account,
	genesisHeader []byte,
) error {

	if txhash, err := s.sdk.Native.Hs.SyncGenesisHeader(
		selfChainID,
		genesisHeader,
		validators,
	); err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("side chain already synced")
			return nil
		}
		return err
	} else {
		return s.waitPolyTx(txhash)
	}
}

func (s *PolySDK) RegisterSideChain(
	owner *polysdk.Account,
	chainID,
	blockToWait,
	router uint64,
	eccdAddr ecm.Address,
	sideChainName string,
) error {

	eccd, err := hex.DecodeString(strings.Replace(eccdAddr.Hex(), "0x", "", 1))
	if err != nil {
		return fmt.Errorf("failed to decode eccd address, err: %s", err)
	}

	if txhash, err := s.sdk.Native.Scm.RegisterSideChain(
		owner.Address,
		chainID,
		router,
		sideChainName,
		blockToWait,
		eccd,
		owner,
	); err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Info("chain %d already registered", chainID)
			return nil
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Info("chain %d already requested", chainID)
			return nil
		}
		return err
	} else {
		return s.waitPolyTx(txhash)
	}
}

func (s *PolySDK) RegisterSideChainExt(
	owner *polysdk.Account,
	chainID,
	blockToWait,
	router uint64,
	eccdAddr ecm.Address,
	sideChainName string,
	extra []byte,
) error {

	eccd, err := hex.DecodeString(strings.Replace(eccdAddr.Hex(), "0x", "", 1))
	if err != nil {
		return fmt.Errorf("failed to decode eccd address, err: %s", err)
	}

	if txhash, err := s.sdk.Native.Scm.RegisterSideChainExt(
		owner.Address,
		chainID,
		router,
		sideChainName,
		blockToWait,
		eccd,
		extra,
		owner,
	); err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Info("chain %d already registered", chainID)
			return nil
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Info("chain %d already requested", chainID)
			return nil
		}
		return err
	} else {
		return s.waitPolyTx(txhash)
	}
}

func (s *PolySDK) ApproveRegisterSideChain(chainID uint64, validators []*polysdk.Account) error {
	var (
		txhash common.Uint256
		err    error
	)
	for i, acc := range validators {
		txhash, err = s.sdk.Native.Scm.ApproveRegisterSideChain(chainID, acc)
		if err != nil {
			return fmt.Errorf("no%d - failed to approve %d: %v", i, chainID, err)
		}
		log.Info("No%d: successful to approve register side chain %d: ( acc: %s, txhash: %s )",
			i, chainID, acc.Address.ToHexString(), txhash.ToHexString())
	}
	return s.waitPolyTx(txhash)
}

func (s *PolySDK) RegisterCandidate(peer string, validator *polysdk.Account) error {
	txHash, err := s.sdk.Native.Nm.RegisterCandidate(peer, validator)
	if err != nil {
		if strings.Contains(err.Error(), "already") {
			log.Warn("candidate %s already registered: %v", peer, err)
			return nil
		}
		return fmt.Errorf("sendTransaction error: %v", err)
	}
	return s.waitPolyTx(txHash)
}

func (s *PolySDK) ApproveCandidate(peer string, validators []*polysdk.Account) error {
	var (
		txhash common.Uint256
		err    error
	)

	for index, validator := range validators {
		txhash, err = s.sdk.Native.Nm.ApproveCandidate(peer, validator)
		if err != nil {
			return fmt.Errorf("node-%d sendTransaction error: %v", index, err)
		}
		log.Info("node-%d approve %s", index, peer)
	}

	return s.waitPolyTx(txhash)
}

func (s *PolySDK) CommitPolyDpos(accArr []*polysdk.Account) error {
	txhash, err := s.sdk.Native.Nm.CommitDpos(accArr)
	if err != nil {
		return err
	}
	return s.waitPolyTx(txhash)
}

func (s *PolySDK) waitPolyTx(hash common.Uint256) error {
	var (
		h         uint32
		tick      = time.NewTicker(1 * time.Second)
		startTime = time.Now()
	)

	for range tick.C {
		h, _ = s.sdk.GetBlockHeightByTxHash(hash.ToHexString())
		curr, _ := s.sdk.GetCurrentBlockHeight()
		if h > 0 && curr > h {
			break
		}

		if startTime.Add(500 * time.Millisecond); startTime.Second() > 300 {
			return fmt.Errorf("tx( %s ) is not confirm for a long time ( over %d sec )",
				hash.ToHexString(), 300)
		}
	}

	return nil
}

func GetBookeeper(block *types.Block) ([]keypair.PublicKey, error) {
	info := new(vconfig.VbftBlockInfo)
	info.NewChainConfig = new(vconfig.ChainConfig)
	if err := json.Unmarshal(block.Header.ConsensusPayload, info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal consensus payload, err: %s", err)
	}

	if info.NewChainConfig == nil {
		return nil, fmt.Errorf("new chain config is nil")
	}

	bookkeepers := make([]keypair.PublicKey, 0)
	for _, peer := range info.NewChainConfig.Peers {
		log.Info("poly peer index %d id %s", peer.Index, peer.ID)
		keystr, _ := hex.DecodeString(peer.ID)
		key, _ := keypair.DeserializePublicKey(keystr)
		bookkeepers = append(bookkeepers, key)
	}
	bookkeepers = keypair.SortPublicKeys(bookkeepers)

	return bookkeepers, nil
}

func AssembleNoCompressBookeeper(bookeepers []keypair.PublicKey) []byte {
	publickeys := make([]byte, 0)
	for _, key := range bookeepers {
		publickeys = append(publickeys, GetOntNoCompressKey(key)...)
	}
	return publickeys
}

func GetOntNoCompressKey(key keypair.PublicKey) []byte {
	var buf bytes.Buffer
	switch t := key.(type) {
	case *ec.PublicKey:
		switch t.Algorithm {
		case ec.ECDSA:
			// Take P-256 as a special case
			if t.Params().Name == elliptic.P256().Params().Name {
				return ec.EncodePublicKey(t.PublicKey, false)
			}
			buf.WriteByte(byte(0x12))
		case ec.SM2:
			buf.WriteByte(byte(0x13))
		}
		label, err := GetCurveLabel(t.Curve.Params().Name)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(label)
		buf.Write(ec.EncodePublicKey(t.PublicKey, false))
	default:
		panic("err")
	}
	return buf.Bytes()
}

func GetCurveLabel(name string) (byte, error) {
	switch strings.ToUpper(name) {
	case strings.ToUpper(elliptic.P224().Params().Name):
		return 1, nil
	case strings.ToUpper(elliptic.P256().Params().Name):
		return 2, nil
	case strings.ToUpper(elliptic.P384().Params().Name):
		return 3, nil
	case strings.ToUpper(elliptic.P521().Params().Name):
		return 4, nil
	case strings.ToUpper(sm2.SM2P256V1().Params().Name):
		return 20, nil
	case strings.ToUpper(btcec.S256().Name):
		return 5, nil
	default:
		panic("err")
	}
}
