package wallet

import (
	"encoding/hex"
	"poly-bridge/utils/leveldb"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/assert"
)

func TestLoadEthAccount(t *testing.T) {
	keystoreDir := "/Users/dylen/software/nft-bridge/poly-nft-bridge/build/devnet/deploy_tool/keystore/eth/"
	storeDir := "/Users/dylen/software/nft-bridge/poly-nft-bridge/build/devnet/deploy_tool/leveldb/"
	account := "0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C"
	passphrase := "111111"
	storage := leveldb.NewLevelDBInstance(storeDir)
	key, err := LoadEthAccount(storage, keystoreDir, account, passphrase)
	assert.NoError(t, err)

	t.Log(key.PublicKey.X.String(), key.PublicKey.Y.String())
}

func TestDecryptKey(t *testing.T) {
	data := "12f66113274159e261c2361f076dce335f917cd9244437129dfdbb640a80a171"
	bz, err := hex.DecodeString(data)
	assert.NoError(t, err)
	ec, err := crypto.ToECDSA(bz)
	assert.NoError(t, err)

	addr := crypto.PubkeyToAddress(ec.PublicKey)
	t.Log(addr.Hex())
}
