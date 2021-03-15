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
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/explorerdao"
	"poly-bridge/crosschaindao/stakedao"
	"poly-bridge/crosschaindao/swapdao"
	"poly-bridge/models"
)

type CrossChainDao interface {
	UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error
	RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error
	GetChain(chainId uint64) (*models.Chain, error)
	UpdateChain(chain *models.Chain) error
	AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error
	AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap) error
	RemoveTokens(tokens []string) error
	RemoveTokenMaps(tokenMaps []*models.TokenMap) error
	Name() string
}

func NewCrossChainDao(server string, backup bool, dbCfg *conf.DBConfig) CrossChainDao {
	if server == basedef.SERVER_POLY_SWAP {
		return swapdao.NewSwapDao(dbCfg, backup)
	} else if server == basedef.SERVER_EXPLORER {
		return explorerdao.NewExplorerDao(dbCfg)
	} else if server == basedef.SERVER_STAKE {
		return stakedao.NewStakeDao()
	} else {
		return nil
	}
}
