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
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"testing"
)

func TestNeoCross(t *testing.T) {
	config := conf.NewConfig("./../../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	neoChainListenConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
	urls := neoChainListenConfig.GetNodesUrl()
	neoSdk := chainsdk.NewNeoSdkPro(urls, neoChainListenConfig.ListenSlot, basedef.NEO_CROSSCHAIN_ID)

	w, err := wallet.NewWalletFromFile("")
	if err != nil {
		panic(err)
	}
	err = w.DecryptAll("1")
	if err != nil {
		panic(err)
	}
	neoAccount := w.Accounts[0]

	sb := sc.NewScriptBuilder()
	scriptHash := helper.HexToBytes("104057f879009326250ee1f5d60e2efd925024e6")
	sb.MakeInvocationScript(scriptHash, "lock", []sc.ContractParameter{})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(neoAccount.Address)
	if err != nil {
		panic(err)
	}
	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0.02)

	tb := tx.NewTransactionBuilder("http://seed1.ngd.network:20332")
	itx, err := tb.MakeInvocationTransaction(script, from, nil, from, sysFee, netFee)
	if err != nil {
		panic(err)
	}

	// sign transaction
	err = tx.AddSignature(itx, neoAccount.KeyPair)
	if err != nil {
		panic(err)
	}
	rawTxString := itx.RawTransactionString()
	result, err := neoSdk.SendRawTransaction(rawTxString)
	if err != nil {
		panic(err)
	}
	if result {
		fmt.Printf("send transaction successful")
	}
	neoSdk.WaitTransactionConfirm(itx.HashString())
}
