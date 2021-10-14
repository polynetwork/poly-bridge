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

package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/utils/decimal"
	"strings"
)

type Request struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint        `json:"id"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
}

// TxnReceipt is the receipt obtained over JSON/RPC from the ethereum client
type TxnReceipt struct {
	BlockHash         *common.Hash    `json:"blockHash"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed *hexutil.Big    `json:"cumulativeGasUsed"`
	GasUsed           *hexutil.Big    `json:"gasUsed"`
	TransactionHash   *common.Hash    `json:"transactionHash"`
	TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	From              *common.Address `json:"from"`
	To                *common.Address `json:"to"`
	Status            *hexutil.Big    `json:"status"`
	L1BlockNumber     *hexutil.Big    `json:"l1BlockNumber"`
}

var chainsNamesCache = map[uint64]string{}

func Init(chains []*Chain) {
	for _, chain := range chains {
		chainsNamesCache[chain.ChainId] = chain.Name
	}
}

func ChainId2Name(id uint64) string {
	name, ok := chainsNamesCache[id]
	if ok {
		return name
	}
	return fmt.Sprintf("%v", id)
}

type BigInt struct {
	big.Int
}

func NewBigIntFromInt(value int64) *BigInt {
	x := new(big.Int).SetInt64(value)
	return NewBigInt(x)
}

func NewBigInt(value *big.Int) *BigInt {
	return &BigInt{Int: *value}
}

func (bigInt *BigInt) Value() (driver.Value, error) {
	if bigInt == nil {
		return "null", nil
	}
	return bigInt.String(), nil
}

func (bigInt *BigInt) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("type error, %v", v)
	}
	str := string(value)
	if str == "null" || str == "nil" || str == "<nil>" {
		return nil
	}
	data, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", value)
	}
	bigInt.Int = *data
	return nil
}

func FormatAmount(precision uint64, amount *BigInt) string {
	precision_new := decimal.NewFromBigInt(big.NewInt(1), int32(precision))
	amount_new := decimal.NewFromBigInt(&amount.Int, 0)
	return amount_new.Div(precision_new).String()
}

func FormatFee(chain uint64, fee *BigInt) string {
	fee_new := decimal.NewFromBigInt(&fee.Int, 0)
	if chain == basedef.BTC_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(100000000), 0)
		return fee_new.Div(precision_new).String() + " BTC"
	} else if chain == basedef.ONT_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000), 0)
		return fee_new.Div(precision_new).String() + " ONG"
	} else if chain == basedef.ETHEREUM_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " ETH"
	} else if chain == basedef.NEO_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(100000000), 0)
		return fee_new.Div(precision_new).String() + " GAS"
	} else if chain == basedef.SWITCHEO_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(100000000), 0)
		return fee_new.Div(precision_new).String() + " SWTH"
	} else if chain == basedef.BSC_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " BNB"
	} else if chain == basedef.O3_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " O3"
	} else if chain == basedef.HECO_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " HT"
	} else if chain == basedef.OK_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " OKT"
	} else if chain == basedef.MATIC_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		return fee_new.Div(precision_new).String() + " MATIC"
	} else if chain == basedef.NEO3_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(100000000), 0)
		return fee_new.Div(precision_new).String() + " GAS"
	} else {
		precision_new := decimal.New(int64(1), 0)
		return fee_new.Div(precision_new).String()
	}
}

func TxType2Name(ty uint32) string {
	return "cross chain transfer"
}
func Precent(a uint64, b uint64) string {
	c := float64(a) / float64(b)
	return fmt.Sprintf("%.2f%%", c*100)
}

func NullToZero(a **BigInt) {
	if *a == nil {
		*a = NewBigInt(new(big.Int).SetInt64(0))
	}
}

func GetL1BlockNumberOfArbitrumTx(hash string) (uint64, error) {
	txHash := common.HexToHash(hash)
	paras := []interface{}{txHash}

	reqPara := &Request{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionReceipt",
		Params:  paras,
		Id:      1,
	}
	reqJson, err := json.Marshal(reqPara)
	arbitrumConfig := conf.GlobalConfig.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
	req, err := http.NewRequest("POST", arbitrumConfig.Nodes[0].Url, strings.NewReader(string(reqJson)))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("RPC response status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	rpcRes := new(Response)
	err = decoder.Decode(&rpcRes)
	if err != nil {
		return 0, fmt.Errorf("GetL1BlockNumberOfArbitrumTx, decode rpcRes failed. err: %s", err)
	}

	receipt := new(TxnReceipt)
	err = json.Unmarshal(rpcRes.Result, receipt)
	if err != nil {
		return 0, fmt.Errorf("GetL1BlockNumberOfArbitrumTx, unmarshal rpcRes.Result err: %s", err)
	}

	if receipt.L1BlockNumber == nil {
		return 0, fmt.Errorf("GetL1BlockNumberOfArbitrumTx failed, receipt.L1BlockNumber is nil")
	}
	l1BlockNumber := receipt.L1BlockNumber.ToInt().Uint64()
	return l1BlockNumber, nil
}
