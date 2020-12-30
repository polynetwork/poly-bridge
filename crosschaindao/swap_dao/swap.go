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
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
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

func (dao *SwapDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	if wrapperTransactions != nil && len(wrapperTransactions) > 0 {
		res := dao.db.Create(wrapperTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update wrapper Transactions failed!")
		}
	}
	if srcTransactions != nil && len(srcTransactions) > 0 {
		res := dao.db.Create(srcTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update src Transactions failed!")
		}
	}
	if polyTransactions != nil && len(polyTransactions) > 0 {
		res := dao.db.Create(polyTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update poly Transactions failed!")
		}
	}
	if dstTransactions != nil && len(dstTransactions) > 0 {
		res := dao.db.Create(dstTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update dst Transactions failed!")
		}
	}
	if chain != nil {
		res := dao.db.Save(chain)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update chain failed!")
		}
	}
	return nil
}

func (dao *SwapDao) GetChain(chainId uint64) (*models.Chain, error) {
	chain := new(models.Chain)
	res := dao.db.Where("chain_id = ?", chainId).First(chain)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	return chain, nil
}

func (dao *SwapDao) UpdateChain(chain *models.Chain) error {
	if chain == nil {
		return fmt.Errorf("no value!")
	}
	res := dao.db.Save(chain)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("no update!")
	}
	return nil
}

func (dao *SwapDao) Start() {
	go dao.Check()
}

func (dao *SwapDao) Check() {
	for {
		dao.check()
	}
}

func (dao *SwapDao) check() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("check events %s......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err := dao.checkHash()
			if err != nil {
				logs.Error("check - err: %s", err)
			}
		}
	}
}

func (dao *SwapDao) checkHash() error {
	polySrcRelations := make([]*models.PolySrcRelation, 0)
	dao.db.Debug().Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ?", "00000000").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("left join src_transactions on poly_transactions.src_hash = src_transactions.key").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	updatePolyTransactions := make([]*models.PolyTransaction, 0)
	for _, polySrcRelation := range polySrcRelations {
		if polySrcRelation.SrcTransaction != nil {
			polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
			updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
		}
	}
	if len(updatePolyTransactions) > 0 {
		dao.db.Save(updatePolyTransactions)
	}
	return nil
}
