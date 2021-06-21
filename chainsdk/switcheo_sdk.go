/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package chainsdk

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/polynetwork/cosmos-poly-module/btcx"
	"github.com/polynetwork/cosmos-poly-module/ccm"
	"github.com/polynetwork/cosmos-poly-module/ft"
	"github.com/polynetwork/cosmos-poly-module/headersync"
	"github.com/polynetwork/cosmos-poly-module/lockproxy"
)

type SwitcheoSDK struct {
	client   *tmclient.HTTP
	cdc      *codec.Codec
	url string
}

func NewSwitcheoSDK(url string) *SwitcheoSDK {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("swth", "swthpub")
	config.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
	config.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")
	rawClient, err := tmclient.New(url, "/websocket")
	if err != nil {
		panic(err)
	}
	cdc := NewCDC()
	return &SwitcheoSDK{
		client:rawClient,
		url: url,
		cdc: cdc,
	}
}

func NewCDC() *codec.Codec {
	cdc := codec.New()
	bank.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	btcx.RegisterCodec(cdc)
	ccm.RegisterCodec(cdc)
	ft.RegisterCodec(cdc)
	headersync.RegisterCodec(cdc)
	lockproxy.RegisterCodec(cdc)
	return cdc
}

func (client *SwitcheoSDK) Status() (*ctypes.ResultStatus, error) {
	return client.client.Status()
}

func (client *SwitcheoSDK) Block(height *int64) (*ctypes.ResultBlock, error) {
	return client.client.Block(height)
}

func (client *SwitcheoSDK) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*ctypes.ResultTxSearch, error) {
	return client.client.TxSearch(query, prove, page, perPage, orderBy)
}

func (client *SwitcheoSDK) GetGas(tx []byte) uint64 {
	decoder := auth.DefaultTxDecoder(client.cdc)
	stdTx, err := decoder(tx)
	if err != nil {
		logs.Error("cosmos client GetGas err: %s", err.Error())
		return 0
	}
	aa, ok := stdTx.(authtypes.StdTx)
	if !ok {
		logs.Error("This is not cosmos std tx!")
		return 0
	}
	amount := aa.Fee.Amount.AmountOf("swth").BigInt()
	gas := big.NewInt(int64(aa.Fee.Gas))
	return amount.Div(amount, gas).Uint64()
}

func (sdk *SwitcheoSDK) GetCurrentBlockHeight() (uint64, error) {
	status, err := sdk.Status()
	if err != nil {
		logs.Error("GetLatestBlockHeight: get current block status %s", err.Error())
	}
	res := status.SyncInfo.LatestBlockHeight
	return uint64(res), nil
}