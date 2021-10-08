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

package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/chainsdk"

	"github.com/ethereum/go-ethereum/common"
	polysdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/native/service/header_sync/bsc"
	"github.com/polynetwork/poly/native/service/header_sync/heco"
)

func SyncEthGenesisHeader2Poly(
	sideChainID uint64,
	sideChainSdk *chainsdk.EthereumSdk,
	polySdk *chainsdk.PolySDK,
	validators []*polysdk.Account,
) (err error) {

	curr, err := sideChainSdk.GetCurrentBlockHeight()
	if err != nil {
		return err
	}
	hdr, err := sideChainSdk.GetHeaderByNumber(curr)
	if err != nil {
		return err
	}

	headerEnc, err := hdr.MarshalJSON()
	if err != nil {
		return err
	}

	if err := polySdk.SyncGenesisBlock(sideChainID, validators, headerEnc); err != nil {
		return err
	}

	return nil
}

func SyncBscGenesisHeader2Poly(
	sideChainID uint64,
	sideChainSdk *chainsdk.EthereumSdk,
	polySdk *chainsdk.PolySDK,
	validators []*polysdk.Account,
) error {

	height, err := sideChainSdk.GetCurrentBlockHeight()
	if err != nil {
		return err
	}

	epochHeight := height - height%200
	pEpochHeight := epochHeight - 200

	hdr, err := sideChainSdk.GetHeaderByNumber(epochHeight)
	if err != nil {
		return err
	}
	phdr, err := sideChainSdk.GetHeaderByNumber(pEpochHeight)
	if err != nil {
		return err
	}
	fmt.Printf("epoch height %d, pEpoch height %d, phdr.extra length %d\r\n", epochHeight, pEpochHeight, len(phdr.Extra))
	pvalidators, err := bsc.ParseValidators(phdr.Extra[32 : len(phdr.Extra)-65])
	if err != nil {
		return err
	}

	if len(hdr.Extra) <= 65+32 {
		return fmt.Errorf("invalid epoch header at height:%d", epochHeight)
	}
	if len(phdr.Extra) <= 65+32 {
		return fmt.Errorf("invalid epoch header at height:%d", pEpochHeight)
	}

	genesisHeader := bsc.GenesisHeader{
		Header: *hdr,
		PrevValidators: []bsc.HeightAndValidators{
			{
				Height:     big.NewInt(int64(pEpochHeight)),
				Validators: pvalidators,
			},
		},
	}

	headerEnc, err := json.Marshal(genesisHeader)
	if err != nil {
		return err
	}

	if err := polySdk.SyncGenesisBlock(sideChainID, validators, headerEnc); err != nil {
		return err
	}

	return nil
}

func SyncHecoGenesisHeader2Poly(
	sideChainID uint64,
	sideChainSdk *chainsdk.EthereumSdk,
	polySdk *chainsdk.PolySDK,
	validators []*polysdk.Account,
) error {

	height, err := sideChainSdk.GetCurrentBlockHeight()
	if err != nil {
		return err
	}

	epochHeight := height - height%200
	pEpochHeight := epochHeight - 200

	hdr, err := sideChainSdk.GetHeaderByNumber(epochHeight)
	if err != nil {
		return err
	}
	phdr, err := sideChainSdk.GetHeaderByNumber(pEpochHeight)
	if err != nil {
		return err
	}
	pvalidators, err := heco.ParseValidators(phdr.Extra[32 : len(phdr.Extra)-65])
	if err != nil {
		return err
	}

	if len(hdr.Extra) <= 65+32 {
		return fmt.Errorf("invalid epoch header at height:%d", epochHeight)
	}
	if len(phdr.Extra) <= 65+32 {
		return fmt.Errorf("invalid epoch header at height:%d", pEpochHeight)
	}

	genesisHeader := bsc.GenesisHeader{
		Header: *hdr,
		PrevValidators: []bsc.HeightAndValidators{
			{
				Height:     big.NewInt(int64(pEpochHeight)),
				Validators: pvalidators,
			},
		},
	}
	headerEnc, err := json.Marshal(genesisHeader)
	if err != nil {
		return err
	}

	if err := polySdk.SyncGenesisBlock(sideChainID, validators, headerEnc); err != nil {
		return err
	}

	return nil
}

func SyncPolyGenesisHeader2Eth(
	polySDK *chainsdk.PolySDK,
	sideChainECCMOwnerKey *ecdsa.PrivateKey,
	sideChainSdk *chainsdk.EthereumSdk,
	sideChainECCM common.Address,
) error {

	// `epoch` related with the poly validators changing,
	// we can set it as 0 if poly validators never changed on develop environment.
	var RCEpoch uint32 = 0
	gB, err := polySDK.GetBlockByHeight(RCEpoch)
	if err != nil {
		return err
	}

	bookeepers, err := chainsdk.GetBookeeper(gB)
	if err != nil {
		return err
	}
	bookeepersEnc := chainsdk.AssembleNoCompressBookeeper(bookeepers)
	headerEnc := gB.Header.ToArray()

	if _, err := sideChainSdk.InitGenesisBlock(
		sideChainECCMOwnerKey,
		sideChainECCM,
		headerEnc,
		bookeepersEnc,
	); err != nil {
		return err
	}

	return nil
}
