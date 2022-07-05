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
	log "github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	bcommon "poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/go_abi/zk_abi"
	"poly-bridge/utils/decimal"
	"strconv"
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

type chainCache struct {
	ChainLogo        string
	ChainExplorerUrl string
	ChainFeeName     string
	ChainFeeLogo     string
}

var chainsNamesCache = map[uint64]string{}
var chainsCache = map[uint64]chainCache{}

func Init(chains []*Chain, chainFees []*ChainFee) {
	for _, chain := range chains {
		chainsNamesCache[chain.ChainId] = chain.Name
		for _, chainFee := range chainFees {
			if chain.ChainId == chainFee.ChainId {
				chainsCache[chain.ChainId] = chainCache{
					chain.ChainLogo,
					chain.ChainExplorerUrl,
					chainFee.TokenBasicName,
					chainFee.TokenBasic.Meta,
				}
			}
		}
		if _, ok := chainsCache[chain.ChainId]; !ok {
			chainsCache[chain.ChainId] = chainCache{
				chain.ChainLogo,
				chain.ChainExplorerUrl,
				"",
				"",
			}
		}
	}
}

func ChainId2Name(id uint64) string {
	name, ok := chainsNamesCache[id]
	if ok {
		return name
	}
	return fmt.Sprintf("%v", id)
}

func ChainId2ChainCache(id uint64) chainCache {
	cache, ok := chainsCache[id]
	if ok {
		return cache
	}
	return chainCache{}
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

	switch chain {
	case basedef.BTC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 8)
		return fee_new.Div(precision_new).String() + " BTC"
	case basedef.ONT_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 9)
		return fee_new.Div(precision_new).String() + " ONG"
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.NEO_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 8)
		return fee_new.Div(precision_new).String() + " GAS"
	case basedef.SWITCHEO_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 8)
		return fee_new.Div(precision_new).String() + " SWTH"
	case basedef.BSC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " BNB"
	case basedef.O3_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " O3"
	case basedef.HECO_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " HT"
	case basedef.OK_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " OKT"
	case basedef.MATIC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " MATIC"
	case basedef.NEO3_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 8)
		return fee_new.Div(precision_new).String() + " GAS"
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.FANTOM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " FTM"
	case basedef.XDAI_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		feeString := fee_new.Div(precision_new).String()
		if basedef.ENV == basedef.TESTNET {
			return feeString + " POA"
		}
		return feeString + " XDai"
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 12)
		return fee_new.Div(precision_new).String() + " ZIL"
	case basedef.AVAX_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " AVAX"
	case basedef.OPTIMISTIC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.METIS_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " METIS"
	case basedef.RINKEBY_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.BOBA_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.OASIS_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ROSE"
	case basedef.HARMONY_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ONE"
	case basedef.KCC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " KCS"
	case basedef.BYTOM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " BTM"
	case basedef.HSC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " HOO"
	case basedef.STARCOIN_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 9)
		return fee_new.Div(precision_new).String() + " STC"
	case basedef.KAVA_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " KAVA"
	case basedef.CUBE_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " CUBE"
	case basedef.ZKSYNC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"
	case basedef.CELO_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " CELO"
	case basedef.CLOVER_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " CLV"
	default:
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

func GetZkSyncL1Height(zkChain, l1Chain *Chain) (height uint64, err error) {
	var zkChainlistenCfg *conf.ChainListenConfig
	for _, cfg := range conf.GlobalConfig.ChainListenConfig {
		if cfg.ChainId == basedef.ZKSYNC_CROSSCHAIN_ID {
			zkChainlistenCfg = cfg
		}
	}

	if len(zkChainlistenCfg.L1Contract) == 0 {
		log.Info("GetZkSyncL1Height zkSync L1Contract not configured")
		return
	}

	var l1Latest uint64
	var l1Client *ethclient.Client

	if basedef.ENV == basedef.TESTNET {
		l1Latest, err = ethGetCurrentHeight(zkChainlistenCfg.L1Url)
		if err != nil {
			log.Info("GetZkSyncL1Height ethGetCurrentHeight failed", "l1Url", zkChainlistenCfg.L1Url, "error", err)
			return
		}
		l1Client, err = ethclient.Dial(zkChainlistenCfg.L1Url)
		if err != nil {
			log.Info("GetZkSyncL1Height create l1 client failed", "l1Url", zkChainlistenCfg.L1Url, "error", err)
			return
		}
	} else {
		if l1Chain == nil {
			log.Error("GetZkSyncL1Height l1Chain invalid")
			return
		}
		l1Latest = l1Chain.Height

		sdk := bcommon.GetSdk(basedef.ETHEREUM_CROSSCHAIN_ID)
		if pro, ok := sdk.(*chainsdk.EthereumSdkPro); ok {
			l1Client = pro.GetClient()
		} else {
			log.Error("GetZkSyncL1Height get l1Chain sdk failed")
			return
		}
	}
	h := l1Latest - zkChain.BackwardBlockNumber

	l1Getter, err := zk_abi.NewIGetters(common.HexToAddress(zkChainlistenCfg.L1Contract), l1Client)
	if err != nil {
		log.Error("GetZkSyncL1Height new zkSync l1 contract getter failed", "error", err)
		return
	}
	n, err := l1Getter.GetTotalBlocksExecuted(&bind.CallOpts{BlockNumber: big.NewInt(int64(h))})
	if err != nil {
		log.Error("GetZkSyncL1Height GetTotalBlocksExecuted failed", "error", err)
		return
	}
	height = uint64(n)
	log.Info("ZkSyncL1Height %d", height)
	return
}

func FormatString(data string) string {
	if len(data) > 64 {
		return data[:64]
	}
	return data
}

func GetTokenType(chainId uint64, standard uint8) string {
	tokenType := ""
	switch standard {
	case TokenTypeErc20:
		tokenType = "20"
	case TokenTypeErc721:
		tokenType = "721"
	default:
		tokenType = "20"
	}
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID, basedef.SWITCHEO_CROSSCHAIN_ID, basedef.PLT_CROSSCHAIN_ID, basedef.ZILLIQA_CROSSCHAIN_ID,
		basedef.MATIC_CROSSCHAIN_ID, basedef.ARBITRUM_CROSSCHAIN_ID, basedef.XDAI_CROSSCHAIN_ID, basedef.AVAX_CROSSCHAIN_ID, basedef.FANTOM_CROSSCHAIN_ID,
		basedef.OPTIMISTIC_CROSSCHAIN_ID, basedef.METIS_CROSSCHAIN_ID, basedef.BOBA_CROSSCHAIN_ID, basedef.RINKEBY_CROSSCHAIN_ID, basedef.OASIS_CROSSCHAIN_ID:
		return "ERC" + "-" + tokenType
	case basedef.ONT_CROSSCHAIN_ID:
		if standard == TokenTypeErc721 {
			return "NFT"
		}
		return "OEP-4"
	case basedef.NEO_CROSSCHAIN_ID:
		if standard == TokenTypeErc721 {
			return "NFT"
		}
		return "NEP-4"
	case basedef.BSC_CROSSCHAIN_ID:
		return "BEP" + "-" + tokenType
	case basedef.HECO_CROSSCHAIN_ID:
		return "HRC" + "-" + tokenType
	case basedef.OK_CROSSCHAIN_ID:
		return "KIP" + "-" + tokenType
	case basedef.NEO3_CROSSCHAIN_ID:
		if standard == TokenTypeErc721 {
			return "NFT"
		}
		return "NEP-17"
	case basedef.HARMONY_CROSSCHAIN_ID:
		return "HRC" + "-" + tokenType
	case basedef.KCC_CROSSCHAIN_ID:
		return "KRC" + "-" + tokenType
	case basedef.BYTOM_CROSSCHAIN_ID:
		return "BAP" + "-" + tokenType
	case basedef.HSC_CROSSCHAIN_ID:
		return "ORC" + "-" + tokenType
	case basedef.KAVA_CROSSCHAIN_ID:
		return "ERC" + "-" + tokenType
	case basedef.CUBE_CROSSCHAIN_ID:
		return "CRC" + "-" + tokenType
	case basedef.STARCOIN_CROSSCHAIN_ID:
		if standard == TokenTypeErc721 {
			return "Starcoin NFT"
		}
		return "Starcoin Token"
	default:
		return "ERC" + "-" + tokenType
	}
}

type heightReq struct {
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      uint     `json:"id"`
}

type heightRep struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      uint   `json:"id"`
}

func ethGetCurrentHeight(url string) (height uint64, err error) {
	req := &heightReq{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		ID:      1,
	}
	data, _ := json.Marshal(req)

	body, err := jsonRequest(url, data)
	if err != nil {
		return
	}

	var resp heightRep
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	height, err = strconv.ParseUint(resp.Result, 0, 64)
	if err != nil {
		return
	}

	return
}

func jsonRequest(url string, data []byte) (result []byte, err error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
