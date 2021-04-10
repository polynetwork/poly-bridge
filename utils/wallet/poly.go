package wallet

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	polysdk "github.com/polynetwork/poly-go-sdk"
)

var (
	sdk   *polysdk.PolySdk
	ponce sync.Once
)

func initPolySdk() {
	ponce.Do(func() {
		sdk = polysdk.NewPolySdk()
	})
}

func LoadPolyAccountList(keystore string, pwd string) ([]*polysdk.Account, error) {
	fs, err := ioutil.ReadDir(keystore)
	if err != nil {
		return nil, err
	}

	list := make([]*polysdk.Account, 0)
	for _, f := range fs {
		fullPath := path.Join(keystore, f.Name())
		fmt.Println("full path is ", fullPath)
		acc, err := LoadPolyAccount(fullPath, pwd)
		if err != nil {
			panic(err)
		}
		list = append(list, acc)
	}
	return list, nil
}

func LoadPolyAccount(path string, pwd string) (*polysdk.Account, error) {
	acc, err := getPolyAccountByPassword(path, []byte(pwd))
	if err != nil {
		return nil, fmt.Errorf("failed to get poly account, err: %s", err)
	}
	return acc, nil
}

func getPolyAccountByPassword(path string, pwd []byte) (*polysdk.Account, error) {

	initPolySdk()

	wallet, err := sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("open wallet error: %v", err)
	}

	acc, err := wallet.GetDefaultAccount(pwd)
	if err == nil {
		return acc, nil
	}

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		curPwd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error, try it again......")
			continue
		}
		curPwd = strings.Trim(curPwd, " ")
		curPwd = strings.Trim(curPwd, "\r")
		curPwd = strings.Trim(curPwd, "\n")
		if acc, err := wallet.GetDefaultAccount([]byte(curPwd)); err == nil {
			return acc, nil
		} else {
			fmt.Printf("password invalid, err %s, try it again......\r\n", err.Error())
		}
	}

	return acc, nil
}
