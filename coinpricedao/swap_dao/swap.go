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

package swap_dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
)

type SwapDao struct {
	dbCfg *conf.DBConfig
	db    *gorm.DB
}

func NewSwapDao(dbCfg *conf.DBConfig) *SwapDao {
	swapDao := &SwapDao{
		dbCfg: dbCfg,
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	swapDao.db = db
	return swapDao
}

func (dao *SwapDao) SavePrices(tokens []*models.TokenBasic) error {
	if tokens != nil && len(tokens) > 0 {
		res := dao.db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(tokens)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (dao *SwapDao) GetTokens() ([]*models.TokenBasic, error) {
	tokens := make([]*models.TokenBasic, 0)
	res := dao.db.Preload("TokenPrices").Find(&tokens)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	return tokens, nil
}

func (dao *SwapDao) GetFees() ([]*models.ChainFee, error) {
	fees := make([]*models.ChainFee, 0)
	res := dao.db.Find(&fees)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	return fees, nil
}
func (dao *SwapDao) SaveFees(fees []*models.ChainFee) error {
	if fees != nil && len(fees) > 0 {
		res := dao.db.Debug().Save(fees)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
