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
	"poly-swap/conf"
	"poly-swap/crosschaindao/explorer_dao"
	"poly-swap/crosschaindao/stake_dao"
	"poly-swap/crosschaindao/swap_dao"
	"poly-swap/models"
)

type CrossChainDao interface {
	UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error
	GetChain(chainId uint64) (*models.Chain, error)
	UpdateChain(chain *models.Chain) error
}

func NewCrossChainDao(server string, dbCfg *conf.DBConfig) CrossChainDao {
	if server == conf.SERVER_POLY_SWAP {
		return swap_dao.NewSwapDao(dbCfg)
	} else if server == conf.SERVER_EXPLORER {
		return explorer_dao.NewExplorerDao(dbCfg)
	} else if server == conf.SERVER_STAKE {
		return stake_dao.NewStakeDao()
	} else {
		return nil
	}
}
