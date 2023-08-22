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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

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
	client *tmclient.HTTP
	cdc    *codec.Codec
	url    string
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
		client: rawClient,
		url:    url,
		cdc:    cdc,
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

func (client *SwitcheoSDK) Health() error {
	_, err := client.client.Health()
	return err
}

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Block struct {
			Header struct {
				Version struct {
					Block string `json:"block"`
				} `json:"version"`
				ChainID     string `json:"chain_id"`
				Height      string `json:"height"`
				Time        string `json:"time"`
				LastBlockID struct {
					Hash string `json:"hash"`
				} `json:"last_block_id"`
			} `json:"header"`
		} `json:"block"`
	} `json:"result"`
}

func (client *SwitcheoSDK) Block(height int64) (*BlockResponse, error) {
	method := "block"
	params := url.Values{}
	params.Add("height", fmt.Sprintf("\"%d\"", height))
	fullURL := func() string {
		if client.url[len(client.url)-1] == '/' {
			return client.url
		}
		return client.url + "/"
	}() + method + "?" + params.Encode()
	response, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}
	blockResponse := new(BlockResponse)
	err = json.Unmarshal(body, blockResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return blockResponse, nil
}

type TxSearchResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Txs []struct {
			Hash     string `json:"hash"`
			Height   string `json:"height"`
			Index    int    `json:"index"`
			TxResult struct {
				GasWanted string `json:"gas_wanted"`
				GasUsed   string `json:"gas_used"`
				Events    []struct {
					Type       string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
					Attributes []struct {
						Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
						Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
						Index bool   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
					} `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
				} `json:"events"`
				Codespace string `json:"codespace"`
			} `json:"tx_result"`
			Tx string `json:"tx"`
		} `json:"txs"`
		TotalCount string `json:"total_count"`
	} `json:"result"`
}

func (client *SwitcheoSDK) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*TxSearchResponse, error) {
	method := "tx_search"
	params := url.Values{}
	params.Add("query", fmt.Sprintf("\"%s\"", query))
	params.Add("prove", fmt.Sprintf("%t", prove))
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("per_page", fmt.Sprintf("%d", perPage))
	params.Add("order_by", fmt.Sprintf("\"%s\"", orderBy))
	fullURL := func() string {
		if client.url[len(client.url)-1] == '/' {
			return client.url
		}
		return client.url + "/"
	}() + method + "?" + params.Encode()
	response, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}
	txSearchResponse := new(TxSearchResponse)
	err = json.Unmarshal(body, txSearchResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return txSearchResponse, nil
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
		logs.Error("GetLatestBlockHeight switcheo: get current block status %s", err.Error())
		return 0, err
	}
	res := status.SyncInfo.LatestBlockHeight
	return uint64(res), nil
}
