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

package stakedao

import (
	"encoding/json"
	"fmt"
	"poly-bridge/basedef"
	"poly-bridge/models"
)

type StakeDao struct {
	chains map[uint64]*models.Chain
}

func NewStakeDao() *StakeDao {
	stakeDao := &StakeDao{}
	chains := make(map[uint64]*models.Chain)
	chains[2] = &models.Chain{
		ChainId: new(uint64),
		Height:  9691129,
	}
	chains[0] = &models.Chain{
		ChainId: new(uint64),
		Height:  1641496,
	}
	chains[6] = &models.Chain{
		ChainId: new(uint64),
		Height:  5698327,
	}
	chains[5] = &models.Chain{
		ChainId: new(uint64),
		Height:  5525704,
	}
	chains[4] = &models.Chain{
		ChainId: new(uint64),
		Height:  0,
	}
	chains[7] = &models.Chain{
		ChainId: new(uint64),
		Height:  2479753,
	}
	chains[3] = &models.Chain{
		ChainId: new(uint64),
		Height:  14017221,
	}

	for k, v := range chains {
		*v.ChainId = k
	}
	stakeDao.chains = chains
	return stakeDao
}

func (dao *StakeDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	{
		json, _ := json.Marshal(chain)
		fmt.Printf("chain: %s\n", json)
	}
	{
		json, _ := json.Marshal(wrapperTransactions)
		fmt.Printf("wrapperTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(srcTransactions)
		fmt.Printf("srcTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(polyTransactions)
		fmt.Printf("polyTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(dstTransactions)
		fmt.Printf("dstTransactions: %s\n", json)
	}
	return nil
}

func (dao *StakeDao) RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error {
	return nil
}

func (dao *StakeDao) GetChain(chainId uint64) (*models.Chain, error) {
	return dao.chains[chainId], nil
}

func (dao *StakeDao) UpdateChain(chain *models.Chain) error {
	dao.chains[*chain.ChainId] = chain
	return nil
}

func (dao *StakeDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap) error {
	return nil
}

func (dao *StakeDao) AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error {
	return nil
}

func (dao *StakeDao) RemoveTokenMaps(tokenMaps []*models.TokenMap) error {
	return nil
}

func (dao *StakeDao) RemoveTokens(tokens []string) error {
	return nil
}

func (dao *StakeDao) Name() string {
	return basedef.SERVER_STAKE
}
