package wallet

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"poly-bridge/utils/leveldb"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/howeyc/gopass"
)

func LoadEthAccount(storage *leveldb.LevelDBImpl, keystore string, address string, pwd string) (*ecdsa.PrivateKey, error) {
	filepath := path.Join(keystore, address)
	enc, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(enc) <= 64 {
		bz, err := hex.DecodeString(string(enc))
		if err != nil {
			return nil, err
		}
		return crypto.ToECDSA(bz)
	}

	key, err := repeatDecrypt(storage, enc, address, pwd)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}

func repeatDecrypt(storage *leveldb.LevelDBImpl, enc []byte, address string, pwd string) (key *keystore.Key, err error) {
	if existPwd, err := getEthPwdSession(storage, address); err == nil {
		return keystore.DecryptKey(enc, existPwd)
	}

	if key, err = keystore.DecryptKey(enc, pwd); err == nil {
		_ = setEthPwdSession(storage, address, pwd)
		return
	}

	fmt.Printf("please input password for ethereum account %s \r\n", address)

	//reader := bufio.NewReader(os.Stdin)
	var (
		curPwd   string
		curPwdBz []byte
	)
	for i := 0; i < 10; i++ {
		//curPwd, err = reader.ReadString('\n')
		if curPwdBz, err = gopass.GetPasswd(); err != nil {
			fmt.Println("input error, try it again......")
			continue
		}
		curPwd = string(curPwdBz)
		curPwd = strings.Trim(curPwd, " ")
		curPwd = strings.Trim(curPwd, "\r")
		curPwd = strings.Trim(curPwd, "\n")
		if key, err = keystore.DecryptKey(enc, curPwd); err == nil {
			_ = setEthPwdSession(storage, address, curPwd)
			return
		} else {
			fmt.Printf("password invalid, err %s, try it again......\r\n", err.Error())
		}
	}
	return
}

func setEthPwdSession(storage *leveldb.LevelDBImpl, address string, pwd string) error {
	key := formatEthKey(address)
	return storage.Set(key, []byte(pwd))
}

func getEthPwdSession(storage *leveldb.LevelDBImpl, address string) (string, error) {
	key := formatEthKey(address)
	bz, err := storage.Get(key)
	if err != nil {
		return "", err
	}
	return string(bz), nil
}

const ethPersistPrefix = "ethereum:account:"

func formatEthKey(address string) []byte {
	return []byte(fmt.Sprintf("%s:%s", ethPersistPrefix, address))
}
