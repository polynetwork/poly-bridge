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
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BridgeDao struct {
	dbCfg  *conf.DBConfig
	db     *gorm.DB
	backup bool
}

func NewBridgeDao(dbCfg *conf.DBConfig, backup bool) *BridgeDao {
	swapDao := &BridgeDao{
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
	swapDao.db = db
	return swapDao
}

func (dao *BridgeDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	if !dao.backup {
		if wrapperTransactions != nil && len(wrapperTransactions) > 0 {
			res := dao.db.Save(wrapperTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if srcTransactions != nil && len(srcTransactions) > 0 {
			res := dao.db.Save(srcTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if polyTransactions != nil && len(polyTransactions) > 0 {
			res := dao.db.Save(polyTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if dstTransactions != nil && len(dstTransactions) > 0 {
			res := dao.db.Save(dstTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if chain != nil {
			res := dao.db.Updates(chain)
			if res.Error != nil {
				return res.Error
			}
		}
		return nil
	} else {
		if wrapperTransactions != nil && len(wrapperTransactions) > 0 {
			for _, wrapperTransaction := range wrapperTransactions {
				wrapperTransaction.Status = 0
				res := dao.db.Updates(wrapperTransaction)
				if res.Error != nil {
					return res.Error
				}
			}
		}
		if srcTransactions != nil && len(srcTransactions) > 0 {
			res := dao.db.Save(srcTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if polyTransactions != nil && len(polyTransactions) > 0 {
			res := dao.db.Save(polyTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		if dstTransactions != nil && len(dstTransactions) > 0 {
			res := dao.db.Save(dstTransactions)
			if res.Error != nil {
				return res.Error
			}
		}
		return nil
	}
}

func (dao *BridgeDao) RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error {
	dao.db.Where("`tx_hash` in ?", srcHashes).Delete(&models.SrcTransfer{})
	dao.db.Where("`hash` in ?", srcHashes).Delete(&models.SrcTransaction{})
	dao.db.Where("`hash` in ?", srcHashes).Delete(&models.WrapperTransaction{})

	dao.db.Where("`hash` in ?", polyHashes).Delete(&models.PolyTransaction{})

	dao.db.Where("`tx_hash` in ?", dstHashes).Delete(&models.DstTransfer{})
	dao.db.Where("`hash` in ?", dstHashes).Delete(&models.DstTransaction{})
	return nil
}

func (dao *BridgeDao) GetChain(chainId uint64) (*models.Chain, error) {
	chain := new(models.Chain)
	res := dao.db.Where("chain_id = ?", chainId).First(chain)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	chain.HeightSwap = 0
	return chain, nil
}

func (dao *BridgeDao) UpdateChain(chain *models.Chain) error {
	if chain == nil {
		return fmt.Errorf("no value!")
	}
	if dao.backup {
		return nil
	}
	res := dao.db.Updates(chain)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("no update!")
	}
	return nil
}

func (dao *BridgeDao) AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error {
	if chain == nil || len(chain) == 0 {
		return nil
	}
	res := dao.db.Create(chain)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("add chain failed!")
	}
	if chainFees == nil || len(chainFees) == 0 {
		return nil
	}
	res = dao.db.Create(chainFees)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("add chain fee failed!")
	}
	return nil
}

func (dao *BridgeDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap) error {
	if tokens != nil && len(tokens) > 0 {
		res := dao.db.Save(tokens)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("add tokens failed!")
		}
	}
	addTokenMaps := dao.getTokenMapsFromToken(tokens)
	addTokenMaps = append(addTokenMaps, tokenMaps...)
	if addTokenMaps != nil && len(addTokenMaps) > 0 {
		res := dao.db.Save(addTokenMaps)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("add tokens map failed!")
		}
	}
	return nil
}

func (dao *BridgeDao) getTokenMapsFromToken(tokenBasics []*models.TokenBasic) []*models.TokenMap {
	tokenMaps := make([]*models.TokenMap, 0)
	for _, tokenBasic := range tokenBasics {
		for _, tokenSrc := range tokenBasic.Tokens {
			for _, tokenDst := range tokenBasic.Tokens {
				if tokenDst.ChainId != tokenSrc.ChainId {
					tokenMaps = append(tokenMaps, &models.TokenMap{
						SrcChainId:   tokenSrc.ChainId,
						SrcTokenHash: tokenSrc.Hash,
						DstChainId:   tokenDst.ChainId,
						DstTokenHash: tokenDst.Hash,
						Property:     1,
					})
				}
			}
		}
	}
	return tokenMaps
}

func (dao *BridgeDao) RemoveTokenMaps(tokenMaps []*models.TokenMap) error {
	for _, tokenMap := range tokenMaps {
		dao.db.Model(&models.TokenMap{}).Where("src_chain_id = ? and src_token_hash = ? and dst_chain_id = ? and dst_token_hash = ?",
			tokenMap.SrcChainId, strings.ToLower(tokenMap.SrcTokenHash), tokenMap.DstChainId, strings.ToLower(tokenMap.DstTokenHash)).Update("property", 0)
		/*
			dao.db.Where("src_chain_id = ? and src_token_hash = ? and dst_chain_id = ? and dst_token_hash = ?",
				tokenMap.SrcChainId, strings.ToLower(tokenMap.SrcTokenHash), tokenMap.DstChainId, strings.ToLower(tokenMap.DstTokenHash)).Delete(&models.TokenMap{})
		*/
	}
	return nil
}

func (dao *BridgeDao) RemoveTokens(tokens []string) error {
	for _, token := range tokens {
		err := dao.RemoveToken(token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao *BridgeDao) RemoveToken(token string) error {
	tokenBasic := new(models.TokenBasic)
	res := dao.db.Model(&models.TokenBasic{}).Where("name = ?", token).Preload("Tokens").Preload("PriceMarkets").First(tokenBasic)
	if res.Error != nil {
		return res.Error
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	tokenBasics = append(tokenBasics, tokenBasic)
	tokenMaps := dao.getTokenMapsFromToken(tokenBasics)
	for _, tokenMap := range tokenMaps {
		dao.db.Where("src_chain_id = ? and src_token_hash = ? and dst_chain_id = ? and dst_token_hash = ?",
			tokenMap.SrcChainId, strings.ToLower(tokenMap.SrcTokenHash), tokenMap.DstChainId, strings.ToLower(tokenMap.DstTokenHash)).Delete(&models.TokenMap{})
	}
	for _, token := range tokenBasic.Tokens {
		dao.db.Where("hash = ? and chain_id = ?", token.Hash, token.ChainId).Delete(&models.Token{})
	}
	for _, priceMarket := range tokenBasic.PriceMarkets {
		dao.db.Where("token_basic_name = ? and market_name = ?", priceMarket.TokenBasicName, priceMarket.MarketName).Delete(&models.PriceMarket{})
	}
	dao.db.Where("name = ?", tokenBasic.Name).Delete(&models.TokenBasic{})
	return nil
}

func (dao *BridgeDao) Name() string {
	return basedef.SERVER_POLY_BRIDGE
}

func (dao *BridgeDao) GetTokenBasics() ([]*models.TokenBasic, error) {
	tokens := make([]*models.TokenBasic, 0)
	res := dao.db.Preload("Tokens").Find(&tokens)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return tokens, res.Error
}

func (dao *BridgeDao) GetTokens() ([]*models.Token, error) {
	tokens := make([]*models.Token, 0)
	res := dao.db.Find(&tokens)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return tokens, res.Error
}

func (dao *BridgeDao) GetLastSrcTransferForToken(assetHashes [][]interface{}) (*models.SrcTransfer, error) {
	transfer := new(models.SrcTransfer)
	res := dao.db.Where("(chain_id, asset) in ?", assetHashes).Order("id desc").Limit(1).First(transfer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return transfer, res.Error
}

func (dao *BridgeDao) AggregateTokenBasicSrcTransfers(assetHashes [][]interface{}, min, max int64) (totalAmount *big.Int, totalCount uint64, err error) {
	var v struct {
		Sum   string
		Count uint64
	}
	res := dao.db.Model(&models.SrcTransfer{}).Select("SUM(amount) as sum, COUNT(*) as count").Where("(chain_id, asset) in ? AND id > ? AND id <= ?", assetHashes, min, max).First(&v)
	err = res.Error
	if res.Error == nil {
		sum := new(big.Float)
		sum.SetString(v.Sum)
		totalAmount, _ = sum.Int(nil)
		totalCount = v.Count
	}
	return
}

func (dao *BridgeDao) UpdateTokenBasicStatsWithCheckPoint(tokenBasic *models.TokenBasic, checkPoint int64) error {
	res := dao.db.Table("token_basics").Where("name = ? AND stats_update_time=?", tokenBasic.Name, checkPoint).Updates(map[string]interface{}{"total_amount": tokenBasic.TotalAmount, "total_count": tokenBasic.TotalCount, "stats_update_time": tokenBasic.StatsUpdateTime})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		logs.Warn("Token basic stats was updated %s", tokenBasic.Name)
	} else {
		logs.Info("Token basic stats successfully updated %s", tokenBasic.Name)
	}
	return nil
}

func (dao *BridgeDao) UpdateTokenAvailableAmount(hash string, chainId uint64, amount *big.Int) error {
	var v interface{}
	if len(amount.String()) > 64 {
		v = strings.Repeat("9", 64)
	} else {
		v = &models.BigInt{*amount}
	}

	res := dao.db.Table("tokens").Where("hash=? AND chain_id=?", hash, chainId).Update("available_amount", v)
	return res.Error
}

func (dao *BridgeDao) CalculateInTokenStatistics(lastId, nowId int64, tokenStatistics interface{}) error {
	res := dao.db.Raw("select count(*) in_counter,  CONVERT(sum(amount), DECIMAL(37, 0)) as in_amount, asset as hash, chain_id as chain_id from dst_transfers where id > ? and id <= ? group by chain_id, asset", lastId, nowId).
		Preload("Token").Preload("Token.TokenBasic").
		Find(tokenStatistics)
	return res.Error
}

func (dao *BridgeDao) CalculateOutTokenStatistics(lastId, nowId int64, tokenStatistics interface{}) error {
	res := dao.db.Raw("select count(*) out_counter,  CONVERT(sum(amount), DECIMAL(37, 0)) as out_amount, asset as hash, chain_id as chain_id from src_transfers where id > ? and id<= ? group by chain_id, asset", lastId, nowId).
		Preload("Token").Preload("Token.TokenBasic").
		Find(tokenStatistics)
	return res.Error
}

func (dao *BridgeDao) GetTokenStatistics(tokenStatistics interface{}) error {
	res := dao.db.Find(tokenStatistics)
	return res.Error
}

func (dao *BridgeDao) SaveTokenStatistic(tokenStatistic *models.TokenStatistic) error {
	res := dao.db.Save(tokenStatistic)
	return res.Error
}

func (dao *BridgeDao) GetNewDstTransfer() (*models.DstTransfer, error) {
	transfer := &models.DstTransfer{}
	res := dao.db.Last(transfer)
	return transfer, res.Error
}
func (dao *BridgeDao) GetNewDstTransaction() (*models.DstTransaction, error) {
	transaction := &models.DstTransaction{}
	res := dao.db.Last(transaction)
	return transaction, res.Error
}
func (dao *BridgeDao) GetNewSrcTransfer() (*models.SrcTransfer, error) {
	srcTransfer := &models.SrcTransfer{}
	res := dao.db.Debug().Last(srcTransfer)
	fmt.Println("GetNewSrcTransfer:", *srcTransfer)
	return srcTransfer, res.Error
}
func (dao *BridgeDao) GetNewSrcTransaction() (*models.SrcTransaction, error) {
	transaction := &models.SrcTransaction{}
	res := dao.db.Debug().Last(transaction)
	fmt.Println("GetNewSrcTransaction:", *transaction)
	return transaction, res.Error
}
func (dao *BridgeDao) GetNewTokenSta() (*models.TokenStatistic, error) {
	tokenStatistic := &models.TokenStatistic{}
	res := dao.db.Last(tokenStatistic)
	return tokenStatistic, res.Error
}
func (dao *BridgeDao) GetNewChainSta() (*models.ChainStatistic, error) {
	chainStatistic := &models.ChainStatistic{}
	logs.Info("qwertyuiop12-----")
	res := dao.db.Debug().Last(chainStatistic)
	logs.Info("qwertyuiop13-----")
	return chainStatistic, res.Error
}
func (dao *BridgeDao) CalculateChainStatisticAssets(chainStatistics interface{}) error {
	res := dao.db.Raw("select count(distinct addresses) as addresses, chain_id from (select  `from` as addresses, chain_id from src_transfers union select `to` as addresses, chain_id from dst_transfers) u group by chain_id").
		Scan(chainStatistics)
	return res.Error
}
func (dao *BridgeDao) GetChainStatistic(chainStatistic interface{}) error {
	res := dao.db.Find(chainStatistic)
	return res.Error
}

func (dao *BridgeDao) GetPolyTransaction() (*models.PolyTransaction, error) {
	polyTransaction := new(models.PolyTransaction)
	res := dao.db.Last(&polyTransaction)
	return polyTransaction, res.Error
}

func (dao *BridgeDao) CalculateInChainStatistics(lastId, nowId int64, chainStatistics interface{}) error {
	res := dao.db.Raw("select count(*) as `in`, chain_id from dst_transactions where id > ? and id <= ? group by chain_id", lastId, nowId).
		Scan(chainStatistics)
	return res.Error
}
func (dao *BridgeDao) CalculateOutChainStatistics(lastId, nowId int64, chainStatistics interface{}) error {
	res := dao.db.Raw("select count(*) as `out`, chain_id from src_transactions where id > ? and id <= ? group by chain_id", lastId, nowId).
		Scan(chainStatistics)
	return res.Error
}
func (dao *BridgeDao) CalculatePolyChainStatistic(lastId, nowId int64) (int64, error) {
	counter := struct {
		Counter int64
	}{}
	res := dao.db.Raw("select count(*) from poly_transactions where id > ? and id <= ?", lastId, nowId).
		Scan(&counter)
	return counter.Counter, res.Error
}

func (dao *BridgeDao) SaveChainStatistics(chainStatistics []*models.ChainStatistic) error {
	fmt.Println("baocunsave----------")
	res := dao.db.Debug().Save(chainStatistics)
	fmt.Println("baocunjieshu")
	return res.Error
}
func (dao *BridgeDao) GetNewAssetSta() (*models.AssetStatistic, error) {
	assetStatistic := new(models.AssetStatistic)
	err := dao.db.Debug().First(assetStatistic).
		Error
	fmt.Println("GetNewAssetSta assetStatistic:", assetStatistic)
	return assetStatistic, err
}
func (dao *BridgeDao) CalculateAsset(lastId, nowId int64) ([]*models.AssetStatistic, error) {
	assetStatistics := make([]*models.AssetStatistic, 0)
	err := dao.db.Debug().Raw("select CONVERT(sum(amount), DECIMAL(37, 0)) as amount, count(*) as txnum, b.token_basic_name  from src_transfers a inner join tokens b on a.chain_id = b.chain_id and a.asset = b.hash where a.id > ? and a.id <= ? group by token_basic_name", lastId, nowId).
		Preload("TokenBasic").Find(&assetStatistics).Error
	jsonAssetStatistics, _ := json.Marshal(assetStatistics[0])
	fmt.Println("CalculateAsset assetStatistics:" + string(jsonAssetStatistics))
	return assetStatistics, err
}
func (dao *BridgeDao) GetAssetStatistic() ([]*models.AssetStatistic, error) {
	assetStatistics := make([]*models.AssetStatistic, 0)
	err := dao.db.Find(&assetStatistics).Error
	return assetStatistics, err
}

func (dao *BridgeDao) SaveAssetStatistic(assetStatistic *models.AssetStatistic) (err error) {
	err = dao.db.Save(assetStatistic).Error
	return err
}
func (dao *BridgeDao) CalculateAssetAdress() ([]*models.AssetStatistic, error) {
	assetStatistics := make([]*models.AssetStatistic, 0)
	err := dao.db.Raw("select count(distinct `from`) as addressnum,b.token_basic_name from src_transfers a inner join tokens b on a.chain_id = b.chain_id and a.asset = b.hash  group by token_basic_name").
		Preload("TokenBasic").Find(&assetStatistics).Error
	return assetStatistics, err
}
func (dao *BridgeDao) UpdateAssetStatisticAdress(assetStatistic *models.AssetStatistic) (err error) {
	err = dao.db.Model(&models.AssetStatistic{}).
		Where("token_basic_name = ? ", assetStatistic.TokenBasicName).
		Update("addressnum", assetStatistic.Addressnum).Error
	return
}
func (dao *BridgeDao) GetBTCPrice() (*models.TokenBasic, error) {
	tokenBasicBTC := new(models.TokenBasic)
	err := dao.db.Where("name='WBTC'").First(tokenBasicBTC).Error
	return tokenBasicBTC, err
}
