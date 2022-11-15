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
	"github.com/beego/beego/v2/core/logs"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/nep17"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/vm"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"strconv"
	"strings"
)

type Neo3Sdk struct {
	client *rpc.RpcClient
	url    string
}
type Nep11Property struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	Series    string `json:"series"`
	Supply    string `json:"supply"`
	Thumbnail string `json:"thumbnail"`
	TokenURI  string `json:"tokenURI"`
}
type GetNep11OwnedByContractHashAddressReq struct {
	Address      string
	ContractHash string
	Limit        int
	Skip         int
}
type NeoRpcRequest struct {
	JsonRpc string      `json:"jsonRpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

type GetNep11OwnedByContractHashAddressRsp struct {
	Id     int `json:"id"`
	Result struct {
		Result []struct {
			Id          string  `json:"_id"`
			Blockhash   string  `json:"blockhash"`
			Contract    string  `json:"contract"`
			From        *string `json:"from"`
			Frombalance string  `json:"frombalance"`
			Timestamp   int64   `json:"timestamp"`
			To          string  `json:"to"`
			Tobalance   string  `json:"tobalance"`
			TokenId     string  `json:"tokenId"`
			Txid        string  `json:"txid"`
			Value       string  `json:"value"`
		} `json:"result"`
		TotalCount int `json:"totalCount"`
	} `json:"result"`
	Error interface{} `json:"error"`
}

func NewNeo3Sdk(url string) *Neo3Sdk {
	return &Neo3Sdk{
		client: rpc.NewClient(url),
		url:    url,
	}
}

func (sdk *Neo3Sdk) GetClient() *rpc.RpcClient {
	return sdk.client
}

func (sdk *Neo3Sdk) GetUrl() string {
	return sdk.url
}

func (sdk *Neo3Sdk) GetBlockCount() (uint64, error) {
	res := sdk.client.GetBlockCount()
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *Neo3Sdk) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
	res := sdk.client.GetBlock(strconv.Itoa(int(index)))
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *Neo3Sdk) GetApplicationLog(txId string) (*models.RpcApplicationLog, error) {
	res := sdk.client.GetApplicationLog(txId)
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *Neo3Sdk) GetTransactionHeight(hash string) (uint64, error) {
	res := sdk.client.GetTransactionHeight(hash)
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *Neo3Sdk) SendRawTransaction(txHex string) (bool, error) {
	res := sdk.client.SendRawTransaction(txHex)
	if res.HasError() {
		return false, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return true, nil
}

func (sdk *Neo3Sdk) Nep17Info(hash string) (string, string, int64, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return "", "", 0, err
	}
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	decimal, err := nep17.Decimals()
	if err != nil {
		return "", "", 0, err
	}
	symbol, err := nep17.Symbol()
	if err != nil {
		return "", "", 0, err
	}
	return hash, symbol, int64(decimal), nil
}

func (sdk *Neo3Sdk) Nep11OwnerOf(assetHash, tokenId string) (string, error) {
	method := "ownerOf"
	var params []models.RpcContractParameter
	tokenIdBase64 := helper.HexToBytes(ConvertTokenIdFromIntStr2HexStr(tokenId))
	params = append(params, models.RpcContractParameter{
		Type:  "ByteArray",
		Value: tokenIdBase64,
	})
	response := sdk.client.InvokeFunction(assetHash, method, params, nil, false)
	errResp := response.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return "", errResp.NetError
		}
		return "", fmt.Errorf("failed to get nep11 owner,%s", errResp.Error.Message)
	}
	stack, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return "", err
	}
	return EncodedHashToNeo3Addr(stack[0].Value.(string))
}

func (sdk *Neo3Sdk) Nep11BalanceOf(assetHash, owner string) (string, error) {
	method := "balanceOf"
	ownerHash160, err := Neo3AddrToHash160(owner)
	if err != nil {
		return "", err
	}
	var params []models.RpcContractParameter
	params = append(params, models.RpcContractParameter{
		Type:  "Hash160",
		Value: ownerHash160,
	})
	response := sdk.client.InvokeFunction(assetHash, method, params, nil, false)
	errResp := response.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return "", errResp.NetError
		}
		return "", fmt.Errorf("failed to get nep11 balance,%s", errResp.Error.Message)
	}
	stack, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return "", err
	}
	return stack[0].Value.(string), nil
}

func (sdk *Neo3Sdk) Nep11TokensOfWithBatchInvoke(assetHash, owner string) ([]string, error) {
	method := "tokensOf"
	res := make([]string, 0)
	ownerHash160, _ := Neo3AddrToHash160(owner)
	var params []models.RpcContractParameter
	params = append(params, models.RpcContractParameter{
		Type:  "Hash160",
		Value: ownerHash160,
	})
	countStr, err := sdk.Nep11BalanceOf(assetHash, owner)
	if err != nil {
		return res, err
	}
	count, err := strconv.ParseInt(countStr, 10, 32)
	if err != nil {
		return res, err
	}
	resp, err := sdk.client.InvokeFunctionAndIterate(assetHash, method, params, nil, false, int32(count))
	if err != nil {
		return res, err
	}
	for _, v := range resp[0] {
		tokenId, _ := crypto.Base64Decode(v.Value.(string))
		res = append(res, ConvertTokenIdFromHexStr2IntStr(helper.BytesToHex(tokenId)))
	}
	return res, nil
}

func (sdk *Neo3Sdk) Nep11TokensOf(assetHash, owner string, start, length int) ([]string, error) {
	skip := start * length
	res := make([]string, 0)
	rsp, err := GetNep11TokenInfoByRPC(owner, assetHash, length, skip)
	if err != nil {
		return nil, err
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("rpc resp error, %v", rsp.Error)
	}
	for _, v := range rsp.Result.Result {
		tokenId, _ := crypto.Base64Decode(v.TokenId)
		res = append(res, ConvertTokenIdFromHexStr2IntStr(helper.BytesToHex(tokenId)))
	}
	return res, nil
}

func (sdk *Neo3Sdk) Nep11PropertiesByBatchInvoke(assetHash string, tokenIds []string) ([]*Nep11Property, error) {
	method := "properties"
	res := make([]*Nep11Property, 0)
	sb := sc.NewScriptBuilder()
	assetHash160, err := helper.UInt160FromString(assetHash)
	if err != nil {
		return res, err
	}
	for _, tokenId := range tokenIds {
		cp := sc.ContractParameter{
			Type:  sc.ByteArray,
			Value: helper.HexToBytes(ConvertTokenIdFromIntStr2HexStr(tokenId)),
		}
		sb.EmitDynamicCall(assetHash160, method, []interface{}{cp})
	}
	bs, err := sb.ToArray()
	if err != nil {
		return res, err
	}
	script := crypto.Base64Encode(bs)
	resp := sdk.client.InvokeScript(script, nil, false)
	errResp := resp.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return res, errResp.NetError
		}
		return res, fmt.Errorf("failed to get nep11 properties,%s", errResp.Error.Message)
	}
	stacks, err := rpc.PopInvokeStacks(resp)
	if err != nil {
		return res, err
	}
	var m map[string]string
	for _, stack := range stacks {
		m = make(map[string]string)
		val1 := stack.Value.([]interface{})
		for _, v := range val1 {
			key, _ := crypto.Base64Decode(v.(map[string]interface{})["key"].(map[string]interface{})["value"].(string))
			value, _ := crypto.Base64Decode(v.(map[string]interface{})["value"].(map[string]interface{})["value"].(string))
			m[string(key)] = string(value)
		}
		property := &Nep11Property{}
		arr, _ := json.Marshal(m)
		_ = json.Unmarshal(arr, &property)
		res = append(res, property)
	}
	return res, nil
}

func (sdk *Neo3Sdk) Nep11UriByBatchInvoke(assetHash string, tokenIds []string) (map[string]string, error) {
	method := "properties"
	res := make(map[string]string)
	sb := sc.NewScriptBuilder()
	assetHash160, err := helper.UInt160FromString(assetHash)
	if err != nil {
		return res, err
	}
	for _, tokenId := range tokenIds {
		cp := sc.ContractParameter{
			Type:  sc.ByteArray,
			Value: helper.HexToBytes(ConvertTokenIdFromIntStr2HexStr(tokenId)),
		}
		sb.EmitDynamicCall(assetHash160, method, []interface{}{cp})
	}
	bs, err := sb.ToArray()
	if err != nil {
		return res, err
	}
	script := crypto.Base64Encode(bs)
	resp := sdk.client.InvokeScript(script, nil, false)
	errResp := resp.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return res, errResp.NetError
		}
		return res, fmt.Errorf("failed to get nep11 properties,%s", errResp.Error.Message)
	}
	stacks, err := rpc.PopInvokeStacks(resp)
	if err != nil {
		return res, err
	}
	for i, stack := range stacks {
		val1 := stack.Value.([]interface{})
		for _, v := range val1 {
			key, _ := crypto.Base64Decode(v.(map[string]interface{})["key"].(map[string]interface{})["value"].(string))
			if string(key) == "tokenURI" {
				value, _ := crypto.Base64Decode(v.(map[string]interface{})["value"].(map[string]interface{})["value"].(string))
				tokenUri := string(value)
				if strings.HasPrefix(tokenUri, "ipfs.io") {
					tokenUri = "https://" + tokenUri
				}
				res[tokenIds[i]] = tokenUri
			}
		}
	}
	return res, nil
}

func (sdk *Neo3Sdk) Nep11PropertiesByInvoke(assetHash, tokenId string) (*Nep11Property, error) {
	method := "properties"
	res := &Nep11Property{}
	sb := sc.NewScriptBuilder()
	assetHash160, err := helper.UInt160FromString(assetHash)
	if err != nil {
		return res, err
	}
	cp := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: helper.HexToBytes(ConvertTokenIdFromIntStr2HexStr(tokenId)),
	}
	sb.EmitDynamicCall(assetHash160, method, []interface{}{cp})
	bs, err := sb.ToArray()
	if err != nil {
		return res, err
	}
	script := crypto.Base64Encode(bs)
	resp := sdk.client.InvokeScript(script, nil, false)
	errResp := resp.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return res, errResp.NetError
		}
		return res, fmt.Errorf("failed to get nep11 properties,%s", errResp.Error.Message)
	}
	stacks, err := rpc.PopInvokeStacks(resp)
	if err != nil {
		return res, err
	}
	var m map[string]string
	stack := stacks[0]
	m = make(map[string]string)
	val1 := stack.Value.([]interface{})
	for _, v := range val1 {
		key, _ := crypto.Base64Decode(v.(map[string]interface{})["key"].(map[string]interface{})["value"].(string))
		value, _ := crypto.Base64Decode(v.(map[string]interface{})["value"].(map[string]interface{})["value"].(string))
		m[string(key)] = string(value)
	}
	arr, _ := json.Marshal(m)
	_ = json.Unmarshal(arr, &res)
	return res, nil
}

//Nep11PropertiesByRPC cannot use tokenId that start with character
func (sdk *Neo3Sdk) Nep11PropertiesByRPC(assetHash, tokenId string) (*Nep11Property, error) {
	response := sdk.client.GetNep11Properties(assetHash, ConvertTokenIdFromIntStr2HexStr(tokenId))
	errResp := response.ErrorResponse
	if errResp.HasError() {
		if errResp.NetError != nil {
			return nil, errResp.NetError
		}
		return nil, fmt.Errorf("failed to get nep11 properties,%s", errResp.Error.Message)
	}
	property := &Nep11Property{}
	arr, _ := json.Marshal(response.Result)
	_ = json.Unmarshal(arr, &property)
	if property == nil {
		return nil, fmt.Errorf("no properties found")
	}
	return property, nil
}

func (sdk *Neo3Sdk) Nep11TokenUri(assetHash, tokenId string) (string, error) {
	property, err := sdk.Nep11PropertiesByInvoke(assetHash, tokenId)
	if err != nil {
		return "", err
	}
	tokenUri := property.TokenURI
	if tokenUri == "" {
		return "", fmt.Errorf("no token uri")
	}
	if strings.HasPrefix(tokenUri, "ipfs.io") {
		tokenUri = "https://" + tokenUri
	}
	return tokenUri, nil
}

func (sdk *Neo3Sdk) Nep17Balance(hash string, addr string) (*big.Int, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	addrHash, err := helper.UInt160FromString(addr)
	if err != nil {
		logs.Info("Nep17Balance err: %s", err)
		return new(big.Int).SetUint64(0), err
	}
	logs.Info("Nep17Balance addrHash: %+v", addrHash)
	return nep17.BalanceOf(addrHash)
}

func (sdk *Neo3Sdk) Nep17TotalSupply(hash string) (*big.Int, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	logs.Info("hash: %s", hash)
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	return nep17.TotalSupply()
}

type InvokeStack struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Convert converts interface{} "Value" to string or []InvokeStack or map[InvokeStack]InvokeStack depending on the "Type"
func (s *InvokeStack) Convert() {
	switch s.Type {
	case vm.Array.String():
		vs := s.Value.([]interface{})
		result := make([]InvokeStack, len(vs))
		for i, v := range vs {
			m := v.(map[string]interface{})
			s2 := InvokeStack{
				Type:  m["type"].(string),
				Value: m["value"],
			}
			s2.Convert()
			result[i] = s2
		}
		s.Value = result
		break
	case vm.Boolean.String():
		if b, ok := s.Value.(bool); ok {
			s.Value = strconv.FormatBool(b)
		}
		break
	case vm.Buffer.String(), vm.ByteString.String():
		// nothing to handle
		break
	case vm.Integer.String():
		if num, ok := s.Value.(int); ok {
			s.Value = strconv.Itoa(num)
		}
		// else if number in string, nothing to handle
		break
	case vm.Map.String():
		vs := s.Value.([]interface{})
		result := make(map[InvokeStack]InvokeStack)
		for _, v := range vs {
			m := v.(map[string]interface{})
			key := m["key"].(map[string]interface{})
			value := m["value"].(map[string]interface{})
			s2 := InvokeStack{
				Type:  key["type"].(string),
				Value: key["value"],
			}
			s3 := InvokeStack{
				Type:  value["type"].(string),
				Value: value["value"],
			}
			s2.Convert()
			s3.Convert()
			result[s2] = s3
		}
		s.Value = result
		break
	case vm.Pointer.String():
		if num, ok := s.Value.(int); ok {
			s.Value = strconv.Itoa(num)
		}
		break
	}
}

func (s *InvokeStack) ToParameter() (*sc.ContractParameter, error) {
	var parameter *sc.ContractParameter = new(sc.ContractParameter)
	var err error
	s.Convert()
	switch s.Type {
	case vm.Array.String():
		parameter.Type = sc.Array
		a := s.Value.([]InvokeStack)
		r := make([]sc.ContractParameter, len(a))
		for i := range a {
			t, err1 := a[i].ToParameter()
			if err1 != nil {
				err = err1
				break
			}
			r[i] = *t
		}
		break
	case vm.Boolean.String():
		parameter.Type = sc.Boolean
		parameter.Value, err = strconv.ParseBool(s.Value.(string))
		break
	case vm.Buffer.String(), vm.ByteString.String():
		parameter.Type = sc.ByteArray
		parameter.Value, err = crypto.Base64Decode(s.Value.(string))
		break
	case vm.Integer.String():
		parameter.Type = sc.Integer
		var b bool
		parameter.Value, b = new(big.Int).SetString(s.Value.(string), 10)
		if !b {
			err = fmt.Errorf("converting vm.Integer to sc.Integer failed")
		}
		break
	case vm.Map.String():
		parameter.Type = sc.Map
		parameter.Value = s.Value // map[InvokeStack]InvokeStack
	case vm.Pointer.String():
		parameter.Type = sc.Integer
		var b bool
		parameter.Value, b = new(big.Int).SetString(s.Value.(string), 10)
		if !b {
			err = fmt.Errorf("converting vm.Pointer to sc.Integer failed")
		}
		break
	default:
		err = fmt.Errorf("not supported stack item type")
	}
	return parameter, err
}

func EncodedHashToNeo3Addr(encodedHash string) (string, error) {
	decodedByte, err := crypto.Base64Decode(encodedHash)
	if err != nil {
		return "", err
	}
	hash160 := helper.UInt160FromBytes(decodedByte)
	return crypto.ScriptHashToAddress(hash160, helper.DefaultAddressVersion), nil
}

func Neo3AddrToHash160(addr string) (*helper.UInt160, error) {
	scriptHash, err := crypto.AddressToScriptHash(addr, helper.DefaultAddressVersion)
	return scriptHash, err
}

func Neo3AddrToReverseHash160(addr string) (*helper.UInt160, error) {
	data, err := crypto.Base58CheckDecode(addr)
	if err != nil {
		return nil, err
	}
	scriptHash := helper.UInt160FromBytes(basedef.HexReverse(data[1:]))
	return scriptHash, nil
}

func Hash160StrToNeo3Addr(hash string) (string, error) {
	hash160, _ := helper.UInt160FromString(hash)
	return crypto.ScriptHashToAddress(hash160, helper.DefaultAddressVersion), nil
}

func ReversedHash160ToNeo3Addr(reversedHash string) (string, error) {
	hash := basedef.HexStringReverse(reversedHash)
	hash160, _ := helper.UInt160FromString(hash)
	return crypto.ScriptHashToAddress(hash160, helper.DefaultAddressVersion), nil
}

func GetNep11TokenInfoByRPC(ownerHash160, assetHash string, limit, skip int) (*GetNep11OwnedByContractHashAddressRsp, error) {
	paras := GetNep11OwnedByContractHashAddressReq{
		Address:      ownerHash160,
		ContractHash: "0x" + assetHash,
		Limit:        limit,
		Skip:         skip,
	}
	reqPara := NeoRpcRequest{
		JsonRpc: "2.0",
		Method:  "GetNep11OwnedByContractHashAddress",
		Params:  paras,
		Id:      1,
	}
	reqJson, err := json.Marshal(reqPara)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://testmagnet.ngd.network", strings.NewReader(string(reqJson)))
	if err != nil {
		return nil, fmt.Errorf("fail to call neo3fura rpc method, %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail to call neo3fura rpc method, %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read neo3fura rpc resp, %v", err)
	}
	rsp := new(GetNep11OwnedByContractHashAddressRsp)
	_ = json.Unmarshal(body, &rsp)
	return rsp, err
}
