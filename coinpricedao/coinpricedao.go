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

package coinpricedao

import (
	"poly-bridge/basedef"
	"poly-bridge/coinpricedao/stakedao"
	"poly-bridge/coinpricedao/swapdao"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type CoinPriceDao interface {
	GetTokens() ([]*models.TokenBasic, error)
	SavePrices(tokens []*models.TokenBasic) error
	Name() string
}

func NewCoinPriceDao(server string, dbCfg *conf.DBConfig) CoinPriceDao {
	if server == basedef.SERVER_STAKE {
		return stakedao.NewStakeDao()
	} else if server == basedef.SERVER_POLY_SWAP {
		return swapdao.NewSwapDao(dbCfg)
	} else {
		return nil
	}
}
