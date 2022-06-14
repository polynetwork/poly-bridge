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
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	serverconf "poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
	"poly-bridge/utils/fee"
	"strings"
	"time"
)

type BridgeDao struct {
	dbCfg             *conf.DBConfig
	chainListenConfig []*conf.ChainListenConfig
	feeListenConfig   []*conf.FeeListenConfig
	db                *gorm.DB
	EstimateProxy     map[string]bool
	EstimateFeeMin    map[uint64]int64
	backup            bool
}

func (dao *BridgeDao) initEstimateProxy() {
	dao.EstimateProxy = make(map[string]bool, 0)
	proxyConfigs := dao.chainListenConfig
	for _, v := range proxyConfigs {
		for _, proxy := range v.OtherProxyContract {
			if proxy.ItemName == "O3V2" {
				dao.EstimateProxy[strings.ToUpper(proxy.ItemProxy)] = true
				dao.EstimateProxy[strings.ToUpper(basedef.HexStringReverse(proxy.ItemProxy))] = true
			}
		}
	}
	logs.Info("init EstimateProxy:", dao.EstimateProxy)
}

func (dao *BridgeDao) initEstimateFeeMin() {
	dao.EstimateFeeMin = make(map[uint64]int64, 0)
	feeListenConfig := dao.feeListenConfig
	for _, v := range feeListenConfig {
		dao.EstimateFeeMin[v.ChainId] = v.MinFee
	}
	logs.Info("init EstimateFeeMin:", dao.EstimateFeeMin)
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
func NewBridgeDaoCheckFee(dbCfg *conf.DBConfig, chainListenConfig []*conf.ChainListenConfig, feeListenConfig []*conf.FeeListenConfig, backup bool) *BridgeDao {
	swapDao := &BridgeDao{
		dbCfg:             dbCfg,
		chainListenConfig: chainListenConfig,
		feeListenConfig:   feeListenConfig,
		backup:            backup,
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
	db.AutoMigrate(
		&models.WrapperTransaction{},
	)
	swapDao.db = db
	swapDao.initEstimateProxy()
	swapDao.initEstimateFeeMin()
	return swapDao
}

func (dao *BridgeDao) UpdateEvents(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
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
			for _, v := range srcTransactions {
				if v.SrcTransfer != nil && v.SrcTransfer.TxHash != "" {
					res := dao.db.
						Table("src_transfers").
						Where("tx_hash = ?", v.SrcTransfer.TxHash).
						Updates(v.SrcTransfer)
					if res.Error != nil {
						return res.Error
					}
				}
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
			for _, v := range dstTransactions {
				if v.DstTransfer != nil && v.DstTransfer.TxHash != "" {
					res := dao.db.Table("dst_transfers").
						Where("tx_hash = ?", v.DstTransfer.TxHash).
						Updates(v.DstTransfer)
					if res.Error != nil {
						return res.Error
					}
				}
			}
		}
		return nil
	} else {
		if wrapperTransactions != nil && len(wrapperTransactions) > 0 {
			for _, wrapperTransaction := range wrapperTransactions {
				res := dao.db.Save(wrapperTransaction)
				if res.RowsAffected > 0 {
					logs.Info("backup wrapperTransaction hash:%v", wrapperTransaction.Hash)
				}
			}
		}
		if srcTransactions != nil && len(srcTransactions) > 0 {
			for _, srcTransaction := range srcTransactions {
				res := dao.db.Debug().Save(srcTransaction)
				if res.RowsAffected == 0 {
					res = dao.db.Debug().Model(&models.SrcTransaction{}).Where("hash = ?", srcTransaction.Hash).Update("key", srcTransaction.Key)
					if res.RowsAffected == 0 {
						continue
					}
				}
				logs.Info("backup srcTransaction hash:%v", srcTransaction.Hash)
				err := dao.db.Table("poly_transactions").Where("(poly_transactions.src_hash = ? or poly_transactions.key = ?) and poly_transactions.time > ? and poly_transactions.src_chain_id = ?", srcTransaction.Key, srcTransaction.Key, 1622476800, srcTransaction.ChainId).
					Update("src_hash", srcTransaction.Hash).Error
				if err != nil {
					return err
				}
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

func (dao *BridgeDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap, servercfg *serverconf.Config) error {
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

func (dao *BridgeDao) CalculateInTokenStatistics(chainId uint64, hash string, lastId, nowId int64) (*models.TokenStatistic, error) {
	tokenStatistic := new(models.TokenStatistic)
	res := dao.db.Raw("select count(*) in_counter,  CONVERT(sum(amount), DECIMAL(37, 0)) as in_amount, chain_id as chain_id, asset as hash  from dst_transfers where  chain_id = ? and asset = ? and id > ? and id <= ? group by chain_id,hash", chainId, hash, lastId, nowId).
		First(tokenStatistic)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return tokenStatistic, res.Error
}

func (dao *BridgeDao) CalculateOutTokenStatistics(chainId uint64, hash string, lastId, nowId int64) (*models.TokenStatistic, error) {
	tokenStatistic := new(models.TokenStatistic)
	res := dao.db.Raw("select count(*) out_counter,  CONVERT(sum(amount), DECIMAL(37, 0)) as out_amount, chain_id as chain_id, asset as hash from src_transfers where chain_id = ? and asset = ? and id > ? and id<= ? group by chain_id,hash", chainId, hash, lastId, nowId).
		First(tokenStatistic)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return tokenStatistic, res.Error
}

func (dao *BridgeDao) GetTokenStatistics() ([]*models.TokenStatistic, error) {
	tokenStatistics := make([]*models.TokenStatistic, 0)
	res := dao.db.Find(&tokenStatistics)
	return tokenStatistics, res.Error
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
	logs.Info("qChainStatistic,sqlGetNewDstTransaction-----")
	transaction := &models.DstTransaction{}
	res := dao.db.Last(transaction)
	logs.Info("qChainStatistic,sqlGetNewDstTransaction,res.Error-----", res.Error)
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
func (dao *BridgeDao) GetNewChainSta() (*models.ChainStatistic, error) {
	chainStatistic := &models.ChainStatistic{}
	res := dao.db.Debug().Last(chainStatistic)
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
func (dao *BridgeDao) GetChains() ([]*models.Chain, error) {
	chains := make([]*models.Chain, 0)
	res := dao.db.Find(&chains)
	return chains, res.Error
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
	var counter int64
	res := dao.db.Debug().Raw("select count(*) as counter from poly_transactions where id > ? and id <= ?", lastId, nowId).
		Scan(&counter)
	return counter, res.Error
}

func (dao *BridgeDao) SaveChainStatistics(chainStatistics []*models.ChainStatistic) error {
	res := dao.db.Debug().Save(chainStatistics)
	return res.Error
}
func (dao *BridgeDao) GetNewAssetSta() (*models.AssetStatistic, error) {
	assetStatistic := new(models.AssetStatistic)
	err := dao.db.Debug().First(assetStatistic).
		Error
	return assetStatistic, err
}
func (dao *BridgeDao) CalculateAssets(tokenBasicName string, lastId, nowId int64) ([]*models.AssetInfo, error) {
	assetInfos := make([]*models.AssetInfo, 0)
	err := dao.db.Debug().Raw("select CONVERT(sum(amount), DECIMAL(37, 0)) as amount, count(*) as txnum, b.token_basic_name, b.precision, c.price  from src_transfers a inner join tokens b on a.chain_id = b.chain_id and a.asset = b.hash left join token_basics c on c.name = b.token_basic_name where b.token_basic_name = ? and a.id > ? and a.id <= ? group by b.chain_id,b.`hash`", tokenBasicName, lastId, nowId).
		Find(&assetInfos).Error
	return assetInfos, err
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
		Find(&assetStatistics).Error
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
func (dao *BridgeDao) GetPropertytokenBasic() ([]*models.TokenBasic, error) {
	tokenBasics := make([]*models.TokenBasic, 0)
	err := dao.db.Where("property = ?", 1).
		Preload("Tokens").
		Find(&tokenBasics).Error
	return tokenBasics, err
}
func (dao *BridgeDao) GetTokenBasicByHash(chainId uint64, hash string) (*models.Token, error) {
	token := new(models.Token)
	err := dao.db.Where("chain_id = ? and hash = ?", chainId, hash).
		Preload("TokenBasic").
		First(token).Error
	return token, err
}

func (dao *BridgeDao) GetDstTransactionByHash(hash string) (*models.DstTransaction, error) {
	dstTransaction := new(models.DstTransaction)
	res := dao.db.Where("hash = ?", hash).First(dstTransaction)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	return dstTransaction, nil
}

type ChainAvgTime struct {
	ChainId uint64
	AvgTime int64
}

func (dao *BridgeDao) GetAvgTimeSrc2Poly(timeLast, timeNow int64) ([]*ChainAvgTime, error) {
	chainAvgTimes := make([]*ChainAvgTime, 0)
	err := dao.db.Raw("SELECT s.chain_id as chain_id,floor(AVG(p.time-s.time)) as avg_time from poly_transactions p left join src_transactions s on s.hash=p.src_hash where s.hash is not null and p.time >= ? and p.time < ?  group by s.chain_id", timeLast, timeNow).
		Find(&chainAvgTimes).Error
	return chainAvgTimes, err
}
func (dao *BridgeDao) GetAvgTimePoly2Dst(timeLast, timeNow int64) ([]*ChainAvgTime, error) {
	chainAvgTimes := make([]*ChainAvgTime, 0)
	err := dao.db.Raw("SELECT d.chain_id as chain_id,floor(AVG(d.time-p.time)) as avg_time from dst_transactions d left join poly_transactions p on p.hash=d.poly_hash where p.hash is not null  and d.time >= ? and d.time < ?  group by d.chain_id", timeLast, timeNow).
		Find(&chainAvgTimes).Error
	return chainAvgTimes, err
}

type TokenStatisticWithName struct {
	TokenBasicName string
	ChainId        uint64
	InAmount       *models.BigInt
	InAmountUsd    *models.BigInt
}

func (dao *BridgeDao) GetSourceTokenStatistics() ([]*TokenStatisticWithName, error) {
	sourceTokenStatistics := make([]*TokenStatisticWithName, 0)
	err := dao.db.Raw("SELECT b.token_basic_name,a.chain_id,a.in_amount,a.in_amount_usd from token_statistics a left join tokens b on a.chain_id=b.chain_id and a.`hash`=b.`hash` left join token_basics c on b.token_basic_name=c.`name` where c.chain_id=a.chain_id").
		Find(&sourceTokenStatistics).Error
	return sourceTokenStatistics, err
}

func (dao *BridgeDao) GetLockTokenStatistics() ([]*models.LockTokenStatistic, error) {
	lockTokenStatistics := make([]*models.LockTokenStatistic, 0)
	err := dao.db.Find(&lockTokenStatistics).
		Error
	return lockTokenStatistics, err
}

func (dao *BridgeDao) SaveLockTokenStatistics(lockTokenStatistics []*models.LockTokenStatistic) error {
	err := dao.db.Save(lockTokenStatistics).Error
	return err
}

func (dao *BridgeDao) FilterMissingWrapperTransactions() ([]*models.SrcTransaction, error) {
	srcTransactions := make([]*models.SrcTransaction, 0)
	startTime := time.Now().Add(-time.Hour * 24).Unix()
	endTime := time.Now().Add(-time.Hour).Unix()
	ignoreSrcChainIds := []uint64{basedef.O3_CROSSCHAIN_ID, basedef.SWITCHEO_CROSSCHAIN_ID}
	ignoreDstChainIds := []uint64{basedef.SWITCHEO_CROSSCHAIN_ID}

	var polyProxies []string
	for k, _ := range conf.PolyProxy {
		polyProxies = append(polyProxies, k)
	}

	res := dao.db.Debug().Where("time > ? and time < ?", startTime, endTime).
		Where("chain_id not in ? and dst_chain_id not in ?", ignoreSrcChainIds, ignoreDstChainIds).
		Where("(select count(1) from wrapper_transactions where src_transactions.hash=wrapper_transactions.hash) = 0").
		Where("UPPER(contract) in ?", polyProxies).
		Find(&srcTransactions)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return srcTransactions, res.Error
}

func (dao *BridgeDao) WrapperTransactionCheckFee(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction) error {
	logs.Info("check fee for poly wrapper Transaction when listening new block ")
	//get chain fee
	chainFees := make([]*models.ChainFee, 0)
	dao.db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	var curSrcTransaction *models.SrcTransaction
	var feePayFloat64, feeMinFloat64, PaidGasFloat64 float64
	for k, v := range wrapperTransactions {
		for _, srcTransaction := range srcTransactions {
			if srcTransaction.Hash == v.Hash {
				curSrcTransaction = srcTransaction
				break
			}
		}
		token := new(models.Token)
		dao.db.Where("hash = ? and chain_id = ?", v.FeeTokenHash, v.SrcChainId).Preload("TokenBasic").Find(token)

		chainFee, ok := chain2Fees[v.DstChainId]
		if !ok {
			v.IsPaid = false
			logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't DstChainId's fee", k)
			continue
		}
		//money paid in wrapper
		feePay, feeMin, gasPay := fee.CheckFeeCal(chainFee, token, v.FeeAmount)
		fmt.Println("cd ..", feePay, feeMin)
		// get optimistic L1 fee on ethereum
		if chainFee.ChainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
			ethChainFee, ok := chain2Fees[basedef.ETHEREUM_CROSSCHAIN_ID]
			if !ok {
				v.IsPaid = false
				logs.Info("check fee poly_hash %s NOT_PAID,chainFee hasn't ethereum fee", k)
				continue
			}

			L1MinFee, _, _, err := fee.GetL1Fee(ethChainFee, chainFee.ChainId)
			if err != nil {
				v.IsPaid = false
				logs.Info("check fee poly_hash %s NOT_PAID, get L1 fee failed. err=%v", k, err)
				continue
			}
			feeMin = new(big.Float).Add(feeMin, L1MinFee)
		}

		if _, in := dao.EstimateProxy[strings.ToUpper(curSrcTransaction.Contract)]; in {
			//is estimateGas proxy
			if gasPay.Cmp(new(big.Float).SetInt64(0)) <= 0 {
				v.IsPaid = false
				continue
			}
			if minFee, in := dao.EstimateFeeMin[v.DstChainId]; in {
				if minFee > 0 && minFee < 100 {
					gasPay = new(big.Float).Mul(gasPay, new(big.Float).SetInt64(100))
					gasPay = new(big.Float).Quo(gasPay, new(big.Float).SetInt64(minFee))
				}
			}
			PaidGasFloat64, _ = gasPay.Float64()
			PaidGas := decimal.NewFromFloat(PaidGasFloat64).Mul(decimal.NewFromInt(100))
			v.PaidGas = models.NewBigInt(PaidGas.BigInt())
			logs.Info("check fee poly_hash %s is EstimateProxy,PaidGas %v", k, v.PaidGas)
			continue
		}
		feeMinFloat64, _ = feeMin.Float64()
		feePayFloat64, _ = feePay.Float64()
		if feePay.Cmp(feeMin) >= 0 {
			v.IsPaid = true
			logs.Info("check fee poly_hash %s PAID,feePay %v >= feeMin %v", v.Hash, feePayFloat64, feeMinFloat64)
		} else {
			v.IsPaid = false
			logs.Info("check fee poly_hash %s NOT_PAID,feePay %v < FluctuatingFeeMin %v", v.Hash, feePayFloat64, feeMinFloat64)
		}
	}
	return nil
}
