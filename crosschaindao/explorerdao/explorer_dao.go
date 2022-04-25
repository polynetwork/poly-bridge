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

package explorerdao

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	serverconf "poly-bridge/conf"
	"poly-bridge/models"
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
	SrcTransfer *SrcTransfer `gorm:"foreignKey:TxHash;references:Hash"`
}

func (SrcTransaction) TableName() string {
	return "fchain_tx"
}

type SrcTransfer struct {
	TxHash     string         `gorm:"column:txhash"`
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
	DstTransfer *DstTransfer `gorm:"foreignKey:TxHash;references:Hash"`
}

func (DstTransaction) TableName() string {
	return "tchain_tx"
}

type DstTransfer struct {
	TxHash  string         `gorm:"column:txhash"`
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

type Token struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	Token     string `gorm:"column:xtoken;default:''"`
	Hash      string `gorm:"primaryKey;column:hash"`
	Name      string `gorm:"column:xname"`
	Type      string `gorm:"column:xtype"`
	Precision string `gorm:"column:xprecision"`
	Desc      string `gorm:"column:xdesc"`
}

func (Token) TableName() string {
	return "chain_token"
}

type TokenBind struct {
	SrcHash string `gorm:"column:hash_src"`
	DstHash string `gorm:"column:hash_dest"`
}

func (TokenBind) TableName() string {
	return "chain_token_bind"
}

type ExplorerDao struct {
	dbCfg  *conf.DBConfig
	db     *gorm.DB
	backup bool
}

func NewExplorerDao(dbCfg *conf.DBConfig, backup bool) *ExplorerDao {
	explorerDao := &ExplorerDao{
		dbCfg:  dbCfg,
		backup: backup,
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
	explorerDao.db = db
	return explorerDao
}

func (dao *ExplorerDao) UpdateEvents(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
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
		for _, transaction := range newSrcTransactions {
			transaction.User = basedef.Hash2Address(transaction.ChainId, transaction.User)
			if transaction.SrcTransfer != nil {
				transaction.SrcTransfer.From = basedef.Hash2Address(transaction.SrcTransfer.ChainId, transaction.SrcTransfer.From)
				transaction.SrcTransfer.To = basedef.Hash2Address(transaction.SrcTransfer.ChainId, transaction.SrcTransfer.To)
				transaction.SrcTransfer.DstUser = basedef.Hash2Address(transaction.SrcTransfer.DstChainId, transaction.SrcTransfer.DstUser)
			}
			if transaction.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
				transaction.Hash, transaction.Key = transaction.Key, transaction.Hash
			}
			if transaction.SrcTransfer != nil {
				transaction.SrcTransfer.TxHash = transaction.Hash
			}
		}
		res := dao.db.Save(newSrcTransactions)
		if res.Error != nil {
			return res.Error
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
		for _, transaction := range newDstTransactions {
			if transaction.DstTransfer != nil {
				transaction.DstTransfer.From = basedef.Hash2Address(transaction.DstTransfer.ChainId, transaction.DstTransfer.From)
				transaction.DstTransfer.To = basedef.Hash2Address(transaction.DstTransfer.ChainId, transaction.DstTransfer.To)
			}
		}
		res := dao.db.Save(newDstTransactions)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (dao *ExplorerDao) RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error {
	dao.db.Where("`txhash` in ?", srcHashes).Delete(&SrcTransfer{})
	dao.db.Where("`txhash` in ?", srcHashes).Delete(&SrcTransaction{})

	dao.db.Where("`txhash` in ?", polyHashes).Delete(&PolyTransaction{})

	dao.db.Where("`txhash` in ?", dstHashes).Delete(&DstTransfer{})
	dao.db.Where("`txhash` in ?", dstHashes).Delete(&DstTransaction{})
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
	newChain.HeightSwap = newChain.Height
	return newChain, nil
}

func (dao *ExplorerDao) UpdateChain(chain *models.Chain) error {
	chainJson, err := json.Marshal(chain)
	if err != nil {
		return err
	}
	if dao.backup {
		return nil
	}
	newChain := new(Chain)
	err = json.Unmarshal(chainJson, newChain)
	if err != nil {
		return err
	}
	if chain.HeightSwap > chain.Height {
		newChain.Height = chain.HeightSwap
	}
	res := dao.db.Model(newChain).Updates(map[string]interface{}{
		"height": gorm.Expr("?", newChain.Height)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (dao *ExplorerDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap, servercfg *serverconf.Config) error {
	explorerTokens, explorerTokenMaps := dao.BuildTokens(tokens)
	if explorerTokens != nil && len(explorerTokens) > 0 {
		res := dao.db.Save(explorerTokens)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update explorer tokens failed!")
		}
	}
	for _, tokenMap := range tokenMaps {
		explorerTokenMaps = append(explorerTokenMaps, &TokenBind{
			SrcHash: tokenMap.SrcTokenHash,
			DstHash: tokenMap.DstTokenHash,
		})
	}
	if explorerTokenMaps != nil && len(explorerTokenMaps) > 0 {
		res := dao.db.Save(explorerTokenMaps)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update explorer tokens map failed!")
		}
	}
	return nil
}

func (dao *ExplorerDao) BuildTokens(tokens []*models.TokenBasic) ([]*Token, []*TokenBind) {
	explorerTokens := make([]*Token, 0)
	explorerTokenBinds := make([]*TokenBind, 0)
	for _, tokenBasic := range tokens {
		var srcToken *TokenBind
		for _, token := range tokenBasic.Tokens {
			explorerToken := &Token{
				Id:        token.ChainId,
				Token:     tokenBasic.PriceMarkets[0].Name,
				Hash:      token.Hash,
				Name:      token.Name,
				Type:      dao.tokenType(token.ChainId),
				Precision: fmt.Sprintf("%d", basedef.Int64FromFigure(int(token.Precision))),
				Desc:      token.TokenBasicName,
			}
			explorerTokens = append(explorerTokens, explorerToken)
			dstTokenHash := token.Hash
			if srcToken != nil {
				dstTokenHash = srcToken.SrcHash
			}
			explorerTokenBind := &TokenBind{
				SrcHash: token.Hash,
				DstHash: dstTokenHash,
			}
			if srcToken == nil {
				srcToken = explorerTokenBind
			}
			explorerTokenBinds = append(explorerTokenBinds, explorerTokenBind)
		}
	}
	return explorerTokens, explorerTokenBinds
}

func (dao *ExplorerDao) tokenType(chainId uint64) string {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		return "erc20"
	} else if chainId == basedef.NEO_CROSSCHAIN_ID {
		return "nep4"
	} else if chainId == basedef.HECO_CROSSCHAIN_ID {
		return "hrc20"
	} else if chainId == basedef.BSC_CROSSCHAIN_ID {
		return "bep20"
	} else if chainId == basedef.ONT_CROSSCHAIN_ID {
		return "oep4"
	} else if chainId == basedef.ONTEVM_CROSSCHAIN_ID {
		return "erc20"
	} else if chainId == basedef.OK_CROSSCHAIN_ID {
		return "kip20"
	} else {
		return ""
	}
}

func (dao *ExplorerDao) AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error {
	return nil
}

func (dao *ExplorerDao) RemoveTokenMaps(tokenMaps []*models.TokenMap) error {
	return nil
}

func (dao *ExplorerDao) RemoveTokens(tokens []string) error {
	return nil
}

func (dao *ExplorerDao) GetTokenBasicByHash(chainId uint64, hash string) (*models.Token, error) {
	return nil, nil
}

func (dao *ExplorerDao) GetDstTransactionByHash(hash string) (*models.DstTransaction, error) {
	return nil, nil
}

func (dao *ExplorerDao) Name() string {
	return basedef.SERVER_EXPLORER
}

type AssetStatistic struct {
	Xname          string
	Addressnum     uint32
	Amount         *models.BigInt `gorm:"type:varchar(64);not null"`
	AmountBtc      *models.BigInt `gorm:"type:varchar(64);not null"`
	AmountUsd      *models.BigInt `gorm:"type:varchar(64);not null"`
	Txnum          uint32
	Hash           string
	TokenBasicName string
}
type ChainInfo struct {
	Id    uint64 `gorm:"column:id"`
	Txin  int64  `gorm:"column:txin"`
	Txout int64  `gorm:"column:txout"`
}
