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

package polylisten

import (
	"context"
	"encoding/hex"
	"fmt"
	zcom "github.com/devfans/zion-sdk/contracts/native/cross_chain_manager/common"
	"github.com/devfans/zion-sdk/contracts/native/go_abi/cross_chain_manager_abi"
	"github.com/devfans/zion-sdk/contracts/native/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type PolyChainListen struct {
	polyCfg *conf.ChainListenConfig
	polySdk *chainsdk.EthereumSdkPro
}

func NewPolyChainListen(cfg *conf.ChainListenConfig) *PolyChainListen {
	polyListen := &PolyChainListen{}
	polyListen.polyCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	polyListen.polySdk = sdk
	return polyListen
}

func (this *PolyChainListen) GetLatestHeight() (uint64, error) {
	return this.polySdk.GetLatestHeight()
}

func (this *PolyChainListen) GetChainListenSlot() uint64 {
	return this.polyCfg.ListenSlot
}

func (this *PolyChainListen) GetChainId() uint64 {
	return this.polyCfg.ChainId
}

func (this *PolyChainListen) GetChainName() string {
	return this.polyCfg.ChainName
}

func (this *PolyChainListen) GetDefer() uint64 {
	return this.polyCfg.Defer
}

func (this *PolyChainListen) GetBatchSize() uint64 {
	return this.polyCfg.BatchSize
}

func (this *PolyChainListen) getECCMEventByBlockNumber(height uint64, tt uint64) ([]*models.PolyTransaction, error) {
	eccmContractAddress := utils.CrossChainManagerContractAddress
	client := this.polySdk.GetClient()
	if client == nil {
		return nil, fmt.Errorf("getECCMEventByBlockNumber GetClient error: nil")
	}
	eccmContract, err := cross_chain_manager_abi.NewCrossChainManagerFilterer(eccmContractAddress, client)

	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	crossChainEvents, err := eccmContract.FilterMakeProof(opt)
	if err != nil {
		return nil, fmt.Errorf("getECCMEventByBlockNumber, filter crossChainEvents :%s", err.Error())
	}
	for crossChainEvents.Next() {
		ev := crossChainEvents.Event
		param := new(zcom.ToMerkleValue)
		value, err := hex.DecodeString(ev.MerkleValueHex)
		if err != nil {
			fmt.Println("hex.DecodeString(ev.MerkleValueHex) err", err)
		}
		err = rlp.DecodeBytes(value, param)
		if err != nil {
			err = fmt.Errorf("rlp decode poly merkle value error %v", err)
			//return nil, err
			fmt.Println(err)
		}
		evt := crossChainEvents.Event
		fee := this.GetConsumeGas(evt.Raw.TxHash)
		polyTransactions = append(polyTransactions, &models.PolyTransaction{
			Hash:       evt.Raw.TxHash.String()[2:],
			ChainId:    this.GetChainId(),
			State:      1,
			Fee:        models.NewBigIntFromInt(int64(fee)),
			Height:     evt.Raw.BlockNumber,
			DstChainId: param.MakeTxParam.ToChainID,
			SrcChainId: param.FromChainID,
			SrcHash: func() string {
				switch param.FromChainID {
				case basedef.NEO_CROSSCHAIN_ID, basedef.NEO3_CROSSCHAIN_ID, basedef.ONT_CROSSCHAIN_ID:
					return basedef.HexStringReverse(hex.EncodeToString(param.MakeTxParam.CrossChainID))
				default:
					return hex.EncodeToString(param.MakeTxParam.CrossChainID)
				}
			}(),
			Time: tt,
		})
	}
	return polyTransactions, nil
}

func (this *PolyChainListen) GetConsumeGas(hash common.Hash) uint64 {
	tx, err := this.polySdk.GetTransactionByHash(hash)
	if err != nil {
		return 0
	}
	receipt, err := this.polySdk.GetTransactionReceipt(hash)
	if err != nil {
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}

func (this *PolyChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	block, err := this.polySdk.GetHeaderByNumber(height)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no poly block!")
	}
	tt := block.Time
	polyTransactions, err := this.getECCMEventByBlockNumber(height, tt)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	return nil, nil, polyTransactions, nil, 0, 0, nil
}

func (this *PolyChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.polyCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
