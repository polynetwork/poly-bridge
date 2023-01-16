package chainsdk

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	ripple_sdk "github.com/polynetwork/ripple-sdk"
	ripple_client "github.com/polynetwork/ripple-sdk/client"
	"github.com/rubblelabs/ripple/websockets"
	"math/big"
)

type RippleSdk struct {
	client *ripple_client.RpcClient
	url    string
}

func NewRippleSdk(url string) *RippleSdk {
	client := ripple_sdk.NewRippleSdk().NewRpcClient().SetAddress(url)
	return &RippleSdk{
		client: client,
		url:    url,
	}
}

func (rs *RippleSdk) GetUrl() string {
	return rs.url
}

func (rs *RippleSdk) GetCurrentBlockHeight() (uint64, error) {
	txBlock, err := rs.client.GetCurrentHeight()
	if err != nil {
		logs.Error("RippleSdk GetCurrentBlockHeight - cannot getLatestTxBlock, err: %s\n", err.Error())
		return 0, err
	}
	return uint64(txBlock), err
}

func (rs *RippleSdk) GetLedger(height uint64) (*websockets.LedgerResult, error) {
	return rs.client.GetLedger(uint32(height))
}

func (rs *RippleSdk) GetXRPBalance(addrhash string) (*big.Int, error) {
	acc, err := rs.client.GetAccountInfo(addrhash)
	if err != nil {
		logs.Error("RippleSdk GetTokenBalance err: %s\n", err.Error())
		return big.NewInt(0), err
	}
	amount, err := acc.AccountData.Balance.NonNative()
	if err != nil {
		logs.Error("RippleSdk GetTokenBalance err: %s\n", err.Error())
		return big.NewInt(0), err
	}
	balance, ok := new(big.Int).SetString(amount.String(), 10)
	if !ok {
		return big.NewInt(0), fmt.Errorf("RippleSdk GetXRPBalance SetString err")
	}
	return balance, nil
}

func (rs *RippleSdk) GetFee() (*big.Int, error) {
	fee, err := rs.client.GetFee()
	if err != nil {
		logs.Error("RippleSdk GetFee err: %s\n", err.Error())
		return big.NewInt(0), err
	}
	rippleFee, err := fee.Drops.OpenLedgerFee.NonNative()
	if err != nil {
		logs.Error("RippleSdk fee NonNative err: %s\n", err.Error())
		return big.NewInt(0), err
	}
	feeAmount, ok := new(big.Int).SetString(rippleFee.String(), 10)
	if !ok {
		return big.NewInt(0), fmt.Errorf("RippleSdk GetFee SetString err")
	}
	return feeAmount, nil
}
