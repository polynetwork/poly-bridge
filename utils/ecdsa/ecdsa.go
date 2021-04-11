package ecdsa

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Key2address(key *ecdsa.PrivateKey) common.Address {
	publicKey := key.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
