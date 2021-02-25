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
	fees []*models.ChainFee
}

func NewStakeDao() *StakeDao {
	stakeDao := &StakeDao{}
	fees := make([]*models.ChainFee, 0)
	feesJson := []byte(`[{"ChainId":2,"TokenBasicName":"Ethereum","TokenBasic":null,"MaxFee":0,"MinFee":0,"ProxyFee":0,"Ind":0},{"ChainId":4,"TokenBasicName":"Neo","TokenBasic":null,"MaxFee":0,"MinFee":0,"ProxyFee":0,"Ind":0},{"ChainId":8,"TokenBasicName":"Ethereum","TokenBasic":null,"MaxFee":0,"MinFee":0,"ProxyFee":0,"Ind":0}]`)
	err := json.Unmarshal(feesJson, &fees)
	if err != nil {
		panic(err)
	}
	stakeDao.fees = fees
	return stakeDao
}

func (dao *StakeDao) GetFees() ([]*models.ChainFee, error) {
	return dao.fees, nil
}
func (dao *StakeDao) SaveFees(fees []*models.ChainFee) error {
	{
		json, _ := json.Marshal(fees)
		fmt.Printf("fees: %s\n", json)
	}
	return nil
}

func (dao *StakeDao) Name() string {
	return basedef.SERVER_STAKE
}
