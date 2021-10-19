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
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	cosmos_types "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/crypto"
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
	if chainId == ETHEREUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == NEO_CROSSCHAIN_ID {
		addrHex, _ := hex.DecodeString(value)
		addr, _ := helper.UInt160FromBytes(addrHex)
		return helper.ScriptHashToAddress(addr)
	} else if chainId == BSC_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == HECO_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == ONT_CROSSCHAIN_ID {
		value = HexStringReverse(value)
		addr, _ := ontcommon.AddressFromHexString(value)
		return addr.ToBase58()
	} else if chainId == OK_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == SWITCHEO_CROSSCHAIN_ID {
		addr, _ := cosmos_types.AccAddressFromHex(value)
		return addr.String()
	} else if chainId == NEO3_CROSSCHAIN_ID {
		addrHex, _ := hex.DecodeString(value)
		addr := neo3_helper.UInt160FromBytes(addrHex)
		address := crypto.ScriptHashToAddress(addr, neo3_helper.DefaultAddressVersion)
		return address
	} else if chainId == BTC_CROSSCHAIN_ID {
		addrHex, _ := hex.DecodeString(value)
		return string(addrHex)
	} else if chainId == ARBITRUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	}
	return value
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

func Address2Hash(chainId uint64, value string) (string, error) {
	if chainId == ETHEREUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == NEO_CROSSCHAIN_ID {
		scripHash, err := helper.AddressToScriptHash(value)
		if err != nil {
			return value, err
		}
		addrBytes := scripHash.Bytes()
		addrHex := hex.EncodeToString(addrBytes)
		return addrHex, nil
	} else if chainId == BSC_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == HECO_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == ONT_CROSSCHAIN_ID {
		addr, err := ontcommon.AddressFromBase58(value)
		if err != nil {
			return value, err
		}
		addrHex := addr.ToHexString()
		return HexStringReverse(addrHex), nil
	} else if chainId == SWITCHEO_CROSSCHAIN_ID {
		//cosmos_types.
		addr, err := cosmos_types.AccAddressFromBech32(value)
		if err != nil {
			return value, err
		}
		hash := fmt.Sprint(addr)
		return hash, nil
	} else if chainId == ARBITRUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	}
	return value, nil
}

func ConfirmEnv(env string) {
	if ENV != env {
		logs.Error("Config env(%s) does not match build env(%s)", env, ENV)
		os.Exit(1)
	}
	logs.Info("Current env: %s", ENV)
}
