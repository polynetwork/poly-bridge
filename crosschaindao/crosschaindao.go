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

package crosschaindao

import (
	"github.com/polynetwork/poly-bridge/conf"
	"github.com/polynetwork/poly-bridge/crosschaindao/explorerdao"
	"github.com/polynetwork/poly-bridge/crosschaindao/stakedao"
	"github.com/polynetwork/poly-bridge/crosschaindao/swapdao"
	"github.com/polynetwork/poly-bridge/models"
)

type CrossChainDao interface {
	UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error
	GetChain(chainId uint64) (*models.Chain, error)
	UpdateChain(chain *models.Chain) error
	Name() string
}

func NewCrossChainDao(server string, dbCfg *conf.DBConfig) CrossChainDao {
	if server == conf.SERVER_POLY_SWAP {
		return swapdao.NewSwapDao(dbCfg)
	} else if server == conf.SERVER_EXPLORER {
		return explorerdao.NewExplorerDao(dbCfg)
	} else if server == conf.SERVER_STAKE {
		return stakedao.NewStakeDao()
	} else {
		return nil
	}
}
