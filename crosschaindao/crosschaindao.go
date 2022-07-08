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
	serverconf "poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"poly-bridge/crosschaindao/explorerdao"
	"poly-bridge/crosschaindao/stakedao"
	"poly-bridge/crosschaindao/swapdao"
	"poly-bridge/models"
)

type CrossChainDao interface {
	FillTxSpecialChain(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction, wrapperDetails []*models.WrapperDetail, polyDetails []*models.PolyDetail) ([]*models.WrapperTransaction, error)
	UpdateEvents(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction, wrapperDetails []*models.WrapperDetail, polyDetails []*models.PolyDetail) error
	RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error
	GetChain(chainId uint64) (*models.Chain, error)
	GetTokenBasicByHash(chainId uint64, hash string) (*models.Token, error)
	GetDstTransactionByHash(hash string) (*models.DstTransaction, error)
	UpdateChain(chain *models.Chain) error
	AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error
	AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap, servercfg *serverconf.Config) error
	RemoveTokens(tokens []string) error
	RemoveTokenMaps(tokenMaps []*models.TokenMap) error
	Name() string
}

func NewCrossChainDao(server string, backup bool, dbCfg *conf.DBConfig) CrossChainDao {
	if server == basedef.SERVER_POLY_SWAP {
		return swapdao.NewSwapDao(dbCfg, backup)
	} else if server == basedef.SERVER_POLY_BRIDGE {
		return bridgedao.NewBridgeDao(dbCfg, backup)
	} else if server == basedef.SERVER_EXPLORER {
		return explorerdao.NewExplorerDao(dbCfg, backup)
	} else if server == basedef.SERVER_STAKE {
		return stakedao.NewStakeDao()
	} else {
		return nil
	}
}
