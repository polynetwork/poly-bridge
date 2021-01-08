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

package bridgesdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PolySwapResp struct {
	Version string
	URL     string
}

type CheckFeesReq struct {
	Hashs []string `json:"Hashs"`
}

type CheckFeeRsp struct {
	Hash string `json:"Hash"`
	HasPay bool `json:"HasPay"`
	Amount string `json:"Amount"`
	Error string `json:"Error"`
}

type CheckFeesRsp struct {
	CheckFees []*CheckFeeRsp `json:"CheckFees"`
}

type BridgeSdk struct {
	url       string
}

func NewBridgeSdk(url string) *BridgeSdk {
	return &BridgeSdk{
		url:       url,
	}
}

func (sdk *BridgeSdk) CheckFee(hashs []string) ([]*CheckFeeRsp, error) {
	checkFeesReq := &CheckFeesReq{Hashs:hashs}
	requestJson, err := json.Marshal(checkFeesReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", sdk.url + "checkfee", strings.NewReader(string(requestJson)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	checkFeesRsp := new(CheckFeesRsp)
	err = json.Unmarshal(respBody, checkFeesRsp)
	if err != nil {
		return nil, err
	}
	return checkFeesRsp.CheckFees, nil
}

func (sdk *BridgeSdk) Info() (bool, error) {
	req, err := http.NewRequest("GET", sdk.url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accepts", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	polySwapResp := new(PolySwapResp)
	err = json.Unmarshal(respBody, polySwapResp)
	if err != nil {
		return false, err
	}
	return true, nil
}