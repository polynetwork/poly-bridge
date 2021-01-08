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

package test

import (
	"encoding/json"
	"fmt"
	"poly-bridge/bridgesdk"
	"testing"
)

func TestBridageSdk(t *testing.T) {
	sdk := bridgesdk.NewBridgeSdkPro([]string{"http://40.115.153.174:30330/v1/"}, 1)
	rsp, err := sdk.CheckFee([]string{"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729"})
	if err != nil {
		panic(err)
	}
	rspJson, _ := json.Marshal(rsp)
	fmt.Printf("rsp: %s\n", string(rspJson))
}
