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

package basedef

import (
	"encoding/hex"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	cosmos_types "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/helper"
	neo3_helper "github.com/joeqian10/neo3-gogogo/helper"
	ontcommon "github.com/ontio/ontology/common"
)

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			logs.Error("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func Hash2Address(chainId uint64, value string) string {
	switch chainId {
	case NEO_CROSSCHAIN_ID:
		addrHex, _ := hex.DecodeString(value)
		addr, _ := helper.UInt160FromBytes(addrHex)
		return helper.ScriptHashToAddress(addr)
	case ONT_CROSSCHAIN_ID:
		value = HexStringReverse(value)
		addr, _ := ontcommon.AddressFromHexString(value)
		return addr.ToBase58()
	case SWITCHEO_CROSSCHAIN_ID:
		addr, _ := cosmos_types.AccAddressFromHex(value)
		return addr.String()
	case NEO3_CROSSCHAIN_ID:
		addrHex, _ := hex.DecodeString(value)
		addr := neo3_helper.UInt160FromBytes(addrHex)
		address := crypto.ScriptHashToAddress(addr, neo3_helper.DefaultAddressVersion)
		return address
	case BTC_CROSSCHAIN_ID:
		addrHex, _ := hex.DecodeString(value)
		return string(addrHex)
	case ZILLIQA_CROSSCHAIN_ID:
		addr, err := bech32.ToBech32Address(value)
		if err == nil {
			return addr
		}
		return value
	default:
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	}
}

func HexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := HexReverse(aa)
	return hex.EncodeToString(bb)
}

func String2Float64(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return v
}

func Int64FromFigure(figure int) int64 {
	x := int64(1)
	for i := 0; i < figure; i++ {
		x *= 10
	}
	return x
}

func StringInSlice(s string, ss []string) bool {
	for _, v := range ss {
		if strings.EqualFold(s, v) {
			return true
		}
	}
	return false
}

func Address2Hash(chainId uint64, value string) (string, error) {
	switch chainId {
	case NEO_CROSSCHAIN_ID:
		scripHash, err := helper.AddressToScriptHash(value)
		if err != nil {
			return value, err
		}
		addrBytes := scripHash.Bytes()
		addrHex := hex.EncodeToString(addrBytes)
		return addrHex, nil
	case ONT_CROSSCHAIN_ID:
		addr, err := ontcommon.AddressFromBase58(value)
		if err != nil {
			return value, err
		}
		addrHex := addr.ToHexString()
		return HexStringReverse(addrHex), nil
	case SWITCHEO_CROSSCHAIN_ID:
		addr, err := cosmos_types.AccAddressFromBech32(value)
		if err != nil {
			return value, err
		}
		hash := fmt.Sprint(addr)
		return hash, nil
	case ZILLIQA_CROSSCHAIN_ID:
		addr, err := bech32.FromBech32Addr(value)
		return addr, err
	default:
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	}
}

//lock item_proxy use
func Proxy2Address(chainId uint64, proxy string) string {
	if chainId == NEO_CROSSCHAIN_ID || chainId == ONT_CROSSCHAIN_ID || chainId == NEO3_CROSSCHAIN_ID {
		proxy = HexStringReverse(proxy)
	}
	return Hash2Address(chainId, proxy)
}

func ConfirmEnv(env string) {
	if ENV != env {
		logs.Error("Config env(%s) does not match build env(%s)", env, ENV)
		os.Exit(1)
	}
	logs.Info("Current env: %s", ENV)
}

func GetChainName(id uint64) string {
	switch id {
	case ZION_CROSSCHAIN_ID:
		return "ZION"
	case ZIONMAIN_CROSSCHAIN_ID:
		return "ZIONMAIN"
	case SIDECHAIN_CROSSCHAIN_ID:
		return "SIDECHAIN"
	case ETHEREUM_CROSSCHAIN_ID:
		return "Ethereum"
	case ONT_CROSSCHAIN_ID:
		return "Ontology"
	case NEO_CROSSCHAIN_ID:
		return "Neo"
	case BSC_CROSSCHAIN_ID:
		return "Bsc"
	case HECO_CROSSCHAIN_ID:
		return "Heco"
	case O3_CROSSCHAIN_ID:
		return "O3"
	case OK_CROSSCHAIN_ID:
		return "OK"
	case MATIC_CROSSCHAIN_ID:
		return "Polygon"
	case ARBITRUM_CROSSCHAIN_ID:
		return "Arbitrum"
	case XDAI_CROSSCHAIN_ID:
		return "XDai"
	case BTC_CROSSCHAIN_ID:
		return "BTC"
	case NEO3_CROSSCHAIN_ID:
		return "Neo3"
	case PLT_CROSSCHAIN_ID:
		return "PLT"
	case KOVAN_CROSSCHAIN_ID:
		return "KOVAN"
	case RINKEBY_CROSSCHAIN_ID:
		return "RINKEBY"
	case GOERLI_CROSSCHAIN_ID:
		return "GOERLI"

	default:
		return fmt.Sprintf("Unknown(%d)", id)
	}
}
