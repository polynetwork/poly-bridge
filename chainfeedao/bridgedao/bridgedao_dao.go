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

package bridgedao

import (
	"fmt"
	"github.com/polynetwork/bridge-common/metrics"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
)

type BridgeDao struct {
	dbCfg *conf.DBConfig
	db    *gorm.DB
}

func NewBridgeDao(dbCfg *conf.DBConfig) *BridgeDao {
	swapDao := &BridgeDao{
		dbCfg: dbCfg,
	}
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	swapDao.db = db
	return swapDao
}

func (dao *BridgeDao) GetFees() ([]*models.ChainFee, error) {
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
func (dao *BridgeDao) SaveFees(fees []*models.ChainFee) error {
	if fees != nil && len(fees) > 0 {
		res := dao.db.Save(fees)
		if res.Error != nil {
			return res.Error
		}
	}
	chainFees := make([]*models.ChainFee, 0)
	dao.db.Preload("TokenBasic").Find(&chainFees)
	for _, v := range chainFees {
		if v.ProxyFee.Cmp(big.NewInt(0)) > 0 {
			proxyFee := decimal.NewFromBigInt(&v.ProxyFee.Int, 0).Div(decimal.NewFromInt(basedef.FEE_PRECISION)).Div(decimal.New(1, int32(v.TokenBasic.Precision))).
				Mul(decimal.New(1, 4)).StringFixed(2)
			metrics.Record(proxyFee, "proxyFee_chain/10^4.%v", v.ChainId)
		}
	}
	return nil
}

func (dao *BridgeDao) Name() string {
	return basedef.SERVER_POLY_BRIDGE
}
