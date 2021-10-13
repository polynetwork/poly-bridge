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

package ethereumlisten

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"poly-bridge/models"
)

var pltLockABIMap map[string]abi.Event

const (
	pltProxyAbiJsonStr = `[
	{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"fromAssetHash","type":"address"},{"indexed":false,"internalType":"address","name":"fromAddress","type":"address"},{"indexed":false,"internalType":"uint64","name":"toChainId","type":"uint64"},{"indexed":false,"internalType":"bytes","name":"toAssetHash","type":"bytes"},{"indexed":false,"internalType":"bytes","name":"toAddress","type":"bytes"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"lock","type":"event"},
	{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"toAssetHash","type":"address"},{"indexed":false,"internalType":"address","name":"toAddress","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"unlock","type":"event"}
]`
)

type LockEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
}

type UnlockEvent struct {
	ToAssetHash common.Address
	ToAddress   common.Address
	Amount      *big.Int
}

func init() {
	pltLockABIMap = make(map[string]abi.Event)
	ab, err := abi.JSON(strings.NewReader(pltProxyAbiJsonStr))
	if err != nil {
		panic(err)
	} else {
		e1 := ab.Events["lock"]
		pltLockABIMap["lock"] = e1
		e2 := ab.Events["unlock"]
		pltLockABIMap["unlock"] = e2
	}
}

func (this *EthereumChainListen) GetPaletteLockProxyLockEvent(hash common.Hash) (ev *models.ProxyLockEvent, err error) {
	var (
		receipt   *types.Receipt
		event     *types.Log
		lockEvent *LockEvent
	)

	proxyAddr := common.HexToAddress("0x0000000000000000000000000000000000000103")
	if receipt, err = this.ethSdk.GetTransactionReceipt(hash); err != nil {
		return
	}

	abEvent := pltLockABIMap["lock"]
	for _, e := range receipt.Logs {
		eid := common.BytesToHash(e.Topics[0][:])
		if eid != abEvent.ID {
			continue
		} else {
			event = e
		}
	}
	if event == nil {
		err = fmt.Errorf("can not find proxy lock event")
		return
	}
	if event.Address != proxyAddr {
		err = fmt.Errorf("expect proxy addr %s, got %s", proxyAddr.Hex(), event.Address.Hex())
		return
	}

	if lockEvent, err = unpackLockEvent(event.Data, abEvent); err != nil {
		return
	}

	ev = &models.ProxyLockEvent{
		Amount:        lockEvent.Amount,
		FromAddress:   lockEvent.FromAddress.String()[2:],
		FromAssetHash: strings.ToLower(lockEvent.FromAssetHash.String()[2:]),
		ToChainId:     uint32(lockEvent.ToChainId),
		ToAssetHash:   hex.EncodeToString(lockEvent.ToAssetHash),
		ToAddress:     hex.EncodeToString(lockEvent.ToAddress),
	}

	return
}

func (this *EthereumChainListen) GetPaletteLockProxyUnlockEvent(hash common.Hash) (toAddress, toAsset common.Address, amount *big.Int, err error) {
	var (
		receipt     *types.Receipt
		event       *types.Log
		unlockEvent *UnlockEvent
	)

	proxyAddr := common.HexToAddress("0x0000000000000000000000000000000000000103")

	if receipt, err = this.ethSdk.GetTransactionReceipt(hash); err != nil {
		return
	}
	if length := len(receipt.Logs); length < 3 {
		err = fmt.Errorf("invalid receipt %s, logs length expect 3, got %d", hash.Hex(), length)
	}

	abEvent := pltLockABIMap["unlock"]
	for _, e := range receipt.Logs {
		eid := common.BytesToHash(e.Topics[0][:])
		if eid != abEvent.ID {
			continue
		} else {
			event = e
		}
	}
	if event == nil {
		err = fmt.Errorf("can not find proxy unlock event")
		return
	}
	if event.Address != proxyAddr {
		err = fmt.Errorf("expect proxy addr %s, got %s", proxyAddr.Hex(), event.Address.Hex())
	}

	if unlockEvent, err = unpackUnlockEvent(event.Data, abEvent); err != nil {
		return
	}
	toAddress = unlockEvent.ToAddress
	toAsset = unlockEvent.ToAssetHash
	amount = unlockEvent.Amount
	return
}

func unpackLockEvent(enc []byte, ab abi.Event) (*LockEvent, error) {
	if unpacked, err := ab.Inputs.Unpack(enc); err == nil {
		event := new(LockEvent)
		err = ab.Inputs.Copy(event, unpacked)
		return event, err
	} else {
		return nil, err
	}
}

func unpackUnlockEvent(enc []byte, ab abi.Event) (*UnlockEvent, error) {
	if unpacked, err := ab.Inputs.Unpack(enc); err == nil {
		event := new(UnlockEvent)
		err = ab.Inputs.Copy(event, unpacked)
		return event, err
	} else {
		return nil, err
	}
}
