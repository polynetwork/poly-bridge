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

package stake_dao

import (
	"encoding/json"
	"fmt"
	"poly-swap/models"
)

type StakeDao struct {
	chains map[uint64]*models.Chain
}

func NewStakeDao() *StakeDao {
	stakeDao := &StakeDao{}
	chains := make(map[uint64]*models.Chain)
	chains[2] = &models.Chain{
		ChainId: new(uint64),
		Name:    "Ethereum",
		Height:  9329384,
	}
	chains[0] = &models.Chain{
		ChainId: new(uint64),
		Name:    "Poly",
		Height:  1641496,
	}
	chains[8] = &models.Chain{
		ChainId: new(uint64),
		Name:    "BSC",
		Height:  4810050,
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

func (dao *StakeDao) GetChain(chainId uint64) (*models.Chain, error) {
	return dao.chains[chainId], nil
}

func (dao *StakeDao) UpdateChain(chain *models.Chain) error {
	dao.chains[*chain.ChainId] = chain
	return nil
}
