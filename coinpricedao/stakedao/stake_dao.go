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
	tokenBasics []*models.TokenBasic
}

func NewStakeDao() *StakeDao {
	stakeDao := &StakeDao{}
	tokenBasics := make([]*models.TokenBasic, 0)
	tokenBasicsJson := []byte(`[{"Name":"Ethereum","Precision":0,"AvgPrice":0,"AvgInd":0,"Time":0,"PriceMarkets":[{"TokenBasicName":"Ethereum","MarketName":"binance","Name":"ETHUSDT","Price":0,"Ind":0,"Time":0,"TokenBasic":null},{"TokenBasicName":"Ethereum","MarketName":"coinmarketcap","Name":"Ethereum","Price":0,"Ind":0,"Time":0,"TokenBasic":null}],"Tokens":[{"Hash":"0000000000000000000000000000000000000000","ChainId":2,"Name":"Ethereum","Precision":18,"TokenBasicName":"Ethereum","TokenBasic":null,"TokenMaps":null},{"Hash":"0000000000000000000000000000000000000005","ChainId":4,"Name":"Ethereum","Precision":18,"TokenBasicName":"Ethereum","TokenBasic":null,"TokenMaps":null}]},{"Name":"Neo","Precision":0,"AvgPrice":0,"AvgInd":0,"Time":0,"PriceMarkets":[{"TokenBasicName":"Neo","MarketName":"binance","Name":"NEOUSDT","Price":0,"Ind":0,"Time":0,"TokenBasic":null},{"TokenBasicName":"Neo","MarketName":"coinmarketcap","Name":"Neo","Price":0,"Ind":0,"Time":0,"TokenBasic":null}],"Tokens":[{"Hash":"0000000000000000000000000000000000000001","ChainId":2,"Name":"Neo","Precision":9,"TokenBasicName":"Neo","TokenBasic":null,"TokenMaps":null},{"Hash":"0000000000000000000000000000000000000006","ChainId":4,"Name":"Neo","Precision":9,"TokenBasicName":"Neo","TokenBasic":null,"TokenMaps":null}]}]`)
	err := json.Unmarshal(tokenBasicsJson, &tokenBasics)
	if err != nil {
		panic(err)
	}
	stakeDao.tokenBasics = tokenBasics
	return stakeDao
}

func (dao *StakeDao) SavePrices(tokens []*models.TokenBasic) error {
	{
		json, _ := json.Marshal(tokens)
		fmt.Printf("tokens: %s\n", json)
	}
	return nil
}

func (dao *StakeDao) GetTokens() ([]*models.TokenBasic, error) {
	return dao.tokenBasics, nil
}

func (dao *StakeDao) Name() string {
	return basedef.SERVER_STAKE
}
