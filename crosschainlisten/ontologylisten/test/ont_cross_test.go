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
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"testing"
)

func TestOntCross(t *testing.T) {
	config := conf.NewConfig("./../../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	ontChainListenConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
	urls := ontChainListenConfig.GetNodesUrl()
	ontsdk := chainsdk.NewOntologySdkPro(urls, ontChainListenConfig.ListenSlot, basedef.ONT_CROSSCHAIN_ID)

	// AScExXzLbkZV32tDFdV7Uoq7ZhCT1bRCGp
	privateKey, err := keypair.WIF2Key([]byte("KyxsqZ45MCx3t2UbuG9P8h96TzyrzTXGRQnfs9nZKFx6YkjTfHqb"))
	if err != nil {
		panic(err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	account := &ontology_go_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}

	contract, _ := common.AddressFromHexString("")
	sdk, _ := ontsdk.GetSdk()
	tx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(2500, 400000, contract, []interface{}{"deposit", []interface{}{
		account.Address}})
	if err != nil {
		panic(err)
	}
	sdk.SetPayer(tx, account.Address)
	err = sdk.SignToTransaction(tx, account)
	if err != nil {
		panic(err)
	}
	txHash, err := sdk.SendTransaction(tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cross successful, hash: %s\n", txHash.ToHexString())
}
