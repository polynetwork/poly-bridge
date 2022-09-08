package toolsmethod

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func CreateAccount() {
	path := os.Getenv("path")
	if path == "" {
		panic(fmt.Sprintf("path is null "))
	}
	pass := os.Getenv("pass")
	if pass == "" {
		fmt.Println("pass is test")
		pass = "test"
	}
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(pass)
	if err != nil {
		panic(fmt.Sprint("NewAccount err", err))
	}
	fmt.Println("addr:", account.Address.Hex())
}

func getPrivateKey(path, pass, addr string) (string, error) {
	keys := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	for _, v := range keys.Accounts() {
		if v.Address.Hex() == common.HexToAddress(addr).Hex() {
			keyJson, err := os.ReadFile(v.URL.Path)
			if err != nil {
				return "", fmt.Errorf("Failed to read the keyJson err:%v, addr:%v", err, addr)
			}
			key, err := keystore.DecryptKey(keyJson, pass)
			if err != nil {
				return "", fmt.Errorf("DecryptKey the keystore err:%v, addr:%v", err, addr)
			}
			privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
			return "0x" + privateKey, nil
		}
	}
	return "", fmt.Errorf("not this addr")
}
