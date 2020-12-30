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

package explorer_dao

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
)

type Chain struct {
	ChainId uint64 `gorm:"column:id"`
	Name    string `gorm:"column:xname"`
	Height  uint64 `gorm:"column:height"`
	In      uint64 `gorm:"column:txin"`
	Out     uint64 `gorm:"column:txout"`
}

func (Chain) TableName() string {
	return "chain_info"
}

type SrcTransaction struct {
	Hash        string       `gorm:"column:txhash"`
	ChainId     uint64       `gorm:"column:chain_id"`
	State       uint64       `gorm:"column:state"`
	Time        uint64       `gorm:"column:tt"`
	Fee         uint64       `gorm:"column:fee"`
	Height      uint64       `gorm:"column:height"`
	User        string       `gorm:"column:xuser"`
	DstChainId  uint64       `gorm:"column:tchain"`
	Contract    string       `gorm:"column:contract"`
	Key         string       `gorm:"column:xkey"`
	Param       string       `gorm:"column:xparam"`
	SrcTransfer *SrcTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

func (SrcTransaction) TableName() string {
	return "fchain_tx"
}

type SrcTransfer struct {
	Hash       string         `gorm:"column:txhash"`
	ChainId    uint64         `gorm:"column:chain_id"`
	Time       uint64         `gorm:"column:tt"`
	Asset      string         `gorm:"column:asset"`
	From       string         `gorm:"column:xfrom"`
	To         string         `gorm:"column:xto"`
	Amount     *models.BigInt `gorm:"column:amount"`
	DstChainId uint64         `gorm:"column:tochainid"`
	DstAsset   string         `gorm:"column:toasset"`
	DstUser    string         `gorm:"column:touser"`
}

func (SrcTransfer) TableName() string {
	return "fchain_transfer"
}

type PolyTransaction struct {
	Hash       string `gorm:"column:txhash"`
	ChainId    uint64 `gorm:"column:chain_id"`
	State      uint64 `gorm:"column:state"`
	Time       uint64 `gorm:"column:tt"`
	Fee        uint64 `gorm:"column:fee"`
	Height     uint64 `gorm:"column:height"`
	SrcChainId uint64 `gorm:"column:fchain"`
	SrcHash    string `gorm:"column:ftxhash"`
	DstChainId uint64 `gorm:"column:tchain"`
	Key        string `gorm:"column:xkey"`
}

func (PolyTransaction) TableName() string {
	return "mchain_tx"
}

type PolySrcRelation struct {
	SrcHash         string
	SrcTransaction  *SrcTransaction `gorm:"foreignKey:Hash;references:SrcHash"`
	PolyHash        string
	PolyTransaction *PolyTransaction `gorm:"foreignKey:Hash;references:PolyHash"`
}

type DstTransaction struct {
	Hash        string       `gorm:"column:txhash"`
	ChainId     uint64       `gorm:"column:chain_id"`
	State       uint64       `gorm:"column:state"`
	Time        uint64       `gorm:"column:tt"`
	Fee         uint64       `gorm:"column:fee"`
	Height      uint64       `gorm:"column:height"`
	SrcChainId  uint64       `gorm:"column:fchain"`
	Contract    string       `gorm:"column:contract"`
	PolyHash    string       `gorm:"column:rtxhash"`
	DstTransfer *DstTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

func (DstTransaction) TableName() string {
	return "tchain_tx"
}

type DstTransfer struct {
	Hash    string         `gorm:"column:txhash"`
	ChainId uint64         `gorm:"column:chain_id"`
	Time    uint64         `gorm:"column:tt"`
	Asset   string         `gorm:"column:asset"`
	From    string         `gorm:"column:xfrom"`
	To      string         `gorm:"column:xto"`
	Amount  *models.BigInt `gorm:"column:amount"`
}

func (DstTransfer) TableName() string {
	return "tchain_transfer"
}

type ExplorerDao struct {
	dbCfg *conf.DBConfig
	db    *gorm.DB
}

func NewExplorerDao(dbCfg *conf.DBConfig) *ExplorerDao {
	explorerDao := &ExplorerDao{
		dbCfg: dbCfg,
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	explorerDao.db = db
	return explorerDao
}

func (dao *ExplorerDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	if srcTransactions != nil && len(srcTransactions) > 0 {
		srcTransactionsJson, err := json.Marshal(srcTransactions)
		if err != nil {
			return err
		}
		newSrcTransactions := make([]*SrcTransaction, 0)
		err = json.Unmarshal(srcTransactionsJson, &newSrcTransactions)
		if err != nil {
			return err
		}
		res := dao.db.Save(newSrcTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update src Transactions failed!")
		}
	}
	if polyTransactions != nil && len(polyTransactions) > 0 {
		polyTransactionsJson, err := json.Marshal(polyTransactions)
		if err != nil {
			return err
		}
		newPolyTransactions := make([]*PolyTransaction, 0)
		err = json.Unmarshal(polyTransactionsJson, &newPolyTransactions)
		if err != nil {
			return err
		}
		res := dao.db.Save(newPolyTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update poly Transactions failed!")
		}
	}
	if dstTransactions != nil && len(dstTransactions) > 0 {
		dstTransactionsJson, err := json.Marshal(dstTransactions)
		if err != nil {
			return err
		}
		newDstTransactions := make([]*DstTransaction, 0)
		err = json.Unmarshal(dstTransactionsJson, &newDstTransactions)
		if err != nil {
			return err
		}
		res := dao.db.Save(newDstTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update dst Transactions failed!")
		}
	}
	if chain != nil {
		chainJson, err := json.Marshal(chain)
		if err != nil {
			return err
		}
		newChain := new(Chain)
		err = json.Unmarshal(chainJson, newChain)
		if err != nil {
			return err
		}
		newChain.In = uint64(len(srcTransactions))
		newChain.Out = uint64(len(dstTransactions))
		res := dao.db.Debug().Model(newChain).Updates(map[string]interface{}{
			"txin":   gorm.Expr("txin + ?", newChain.In),
			"txout":  gorm.Expr("txout + ?", newChain.Out),
			"height": gorm.Expr("?", newChain.Height)})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update chain failed!")
		}
	}
	return nil
}

func (dao *ExplorerDao) GetChain(chainId uint64) (*models.Chain, error) {
	chain := new(Chain)
	res := dao.db.Where("id = ?", chainId).First(chain)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	chainJson, err := json.Marshal(chain)
	if err != nil {
		return nil, err
	}
	newChain := new(models.Chain)
	err = json.Unmarshal(chainJson, newChain)
	if err != nil {
		return nil, err
	}
	return newChain, nil
}

func (dao *ExplorerDao) UpdateChain(chain *models.Chain) error {
	chainJson, err := json.Marshal(chain)
	if err != nil {
		return err
	}
	newChain := new(Chain)
	err = json.Unmarshal(chainJson, newChain)
	if err != nil {
		return err
	}
	res := dao.db.Model(newChain).Updates(map[string]interface{}{
		"height": gorm.Expr("?", newChain.Height)})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("update chain failed!")
	}
	return nil
}