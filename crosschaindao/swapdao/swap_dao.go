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

package swapdao

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/coinpricelisten/coinmarketcap"
	"poly-bridge/conf"
	serverconf "poly-bridge/conf"
	"poly-bridge/models"
	"strings"
	"time"
)

type SwapDao struct {
	dbCfg  *conf.DBConfig
	db     *gorm.DB
	backup bool
}

func NewSwapDao(dbCfg *conf.DBConfig, backup bool) *SwapDao {
	swapDao := &SwapDao{
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

func (dao *SwapDao) UpdateEvents(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction, wrapperDetails []*models.WrapperDetail, polySignDetails []*models.PolyDetail) error {
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

func (dao *SwapDao) RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error {
	dao.db.Where("`tx_hash` in ?", srcHashes).Delete(&models.SrcTransfer{})
	dao.db.Where("`hash` in ?", srcHashes).Delete(&models.SrcTransaction{})
	dao.db.Where("`hash` in ?", srcHashes).Delete(&models.WrapperTransaction{})

	dao.db.Where("`hash` in ?", polyHashes).Delete(&models.PolyTransaction{})

	dao.db.Where("`tx_hash` in ?", dstHashes).Delete(&models.DstTransfer{})
	dao.db.Where("`hash` in ?", dstHashes).Delete(&models.DstTransaction{})
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
	chain.Height = 0
	return chain, nil
}

func (dao *SwapDao) UpdateChain(chain *models.Chain) error {
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

func (dao *SwapDao) AddChains(chain []*models.Chain, chainFees []*models.ChainFee) error {
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

//now mainnet use polyswap
func (dao *SwapDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap, servercfg *serverconf.Config) error {
	if tokens != nil && len(tokens) > 0 {
		for _, basic := range tokens {
			for _, token := range basic.Tokens {
				token.Standard = basic.Standard
				token.Property = basic.Property
				if basic.Standard == models.TokenTypeErc721 {
					token.Name = basic.Name
				}
			}
		}
		if servercfg != nil {
			var coinmarketsdk *coinmarketcap.CoinMarketCapSdk
			for _, coinconfig := range servercfg.CoinPriceListenConfig {
				if coinconfig.MarketName == basedef.MARKET_COINMARKETCAP {
					coinmarketsdk = coinmarketcap.NewCoinMarketCapSdk(coinconfig)
					break
				}
			}

			coinIds := make([]string, 0)
			for _, tokenBasic := range tokens {
				if tokenBasic != nil && tokenBasic.PriceMarkets != nil && len(tokenBasic.PriceMarkets) > 0 && tokenBasic.Standard == models.TokenTypeErc20 {
					for _, priceMarket := range tokenBasic.PriceMarkets {
						if priceMarket.MarketName == basedef.MARKET_COINMARKETCAP && priceMarket.CoinMarketId > 0 {
							fmt.Printf("start update token:%v CoinMarketId:%v coinmarketcap price\n", tokenBasic.Name, priceMarket.CoinMarketId)
							coinIds = append(coinIds, fmt.Sprintf("%d", priceMarket.CoinMarketId))
						}
					}
				}
			}
			requestCoinIds := strings.Join(coinIds, ",")
			quotes, err := coinmarketsdk.QuotesLatest(requestCoinIds)
			if err != nil {
				time.Sleep(time.Second * 3)
				quotes, err = coinmarketsdk.QuotesLatest(requestCoinIds)
			}
			if err != nil {
				logs.Error("coinmarketsdk.QuotesLatest err:", err)
			}
			coinId2Price := make(map[int]*coinmarketcap.Ticker)
			if err == nil {
				for _, v := range quotes {
					coinId2Price[v.ID] = v
				}
				jsonQuotes, _ := json.MarshalIndent(quotes, "", "	")
				fmt.Println(string(jsonQuotes))
			}
			for _, tokenBasic := range tokens {
				if tokenBasic != nil && tokenBasic.Standard == models.TokenTypeErc20 && tokenBasic.PriceMarkets != nil {
					for _, priceMarket := range tokenBasic.PriceMarkets {
						if priceMarket.MarketName == basedef.MARKET_COINMARKETCAP {
							if tokenBasic.Price > 0 {
								priceMarket.Price = tokenBasic.Price
								priceMarket.Time = time.Now().Unix()
								priceMarket.Ind = 1
								fmt.Printf("end update token:%v CoinMarketId:%v coinmarketcap price%v \n", tokenBasic.Name, priceMarket.CoinMarketId, tokenBasic.Price)
							} else if priceMarket.Price > 0 {
								priceMarket.Time = time.Now().Unix()
								priceMarket.Ind = 1
								tokenBasic.Price = priceMarket.Price
								fmt.Printf("end update token:%v CoinMarketId:%v coinmarketcap price%v \n", tokenBasic.Name, priceMarket.CoinMarketId, tokenBasic.Price)
							} else {
								if priceMarket.CoinMarketId > 0 {
									priceTicker, ok := coinId2Price[priceMarket.CoinMarketId]
									if ok {
										priceMarket.Name = priceTicker.Name
										if priceTicker.Quote == nil || priceTicker.Quote["USD"] == nil {
											fmt.Printf(" There is no price for coin %s in CoinMarketCap!\n", tokenBasic.Name)
											continue
										}
										price, _ := new(big.Float).Mul(big.NewFloat(priceTicker.Quote["USD"].Price), big.NewFloat(float64(basedef.PRICE_PRECISION))).Int64()

										priceMarket.Price = price
										priceMarket.Time = time.Now().Unix()
										priceMarket.Ind = 1
										tokenBasic.Price = price
										fmt.Printf("end update token: %v CoinMarketId: %v coinmarketcap price: %v \n", tokenBasic.Name, priceMarket.CoinMarketId, tokenBasic.Price)
										break
									}
								}
							}
						}
					}
				}
			}
		}

		res := dao.db.Save(tokens)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			logs.Error("add tokenbascis failed. please check if it already exist!")
			//return fmt.Errorf("add tokens failed!")
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

func (dao *SwapDao) getTokenMapsFromToken(tokenBasics []*models.TokenBasic) []*models.TokenMap {
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
						Property:     tokenBasic.Property,
						Standard:     tokenBasic.Standard,
					})
				}
			}
		}
	}
	return tokenMaps
}

func (dao *SwapDao) RemoveTokenMaps(tokenMaps []*models.TokenMap) error {
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

func (dao *SwapDao) RemoveTokens(tokens []string) error {
	for _, token := range tokens {
		err := dao.RemoveToken(token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao *SwapDao) RemoveToken(token string) error {
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
		if priceMarket.MarketName == basedef.MARKET_COINMARKETCAP {
			dao.db.Where("token_basic_name = ? and market_name = ?", priceMarket.TokenBasicName, priceMarket.MarketName).Delete(&models.PriceMarket{})
		}
	}
	dao.db.Where("name = ?", tokenBasic.Name).Delete(&models.TokenBasic{})
	return nil
}

func (dao *SwapDao) Name() string {
	return basedef.SERVER_POLY_SWAP
}

func (dao *SwapDao) UpdateNFTProfileTokenName(oldName, newName string) {
	dao.db.Table("nft_profiles").
		Where("token_basic_name = ?", oldName).
		Update("token_basic_name", newName)
}

func (dao *SwapDao) GetTokenBasicByHash(chainId uint64, hash string) (*models.Token, error) {
	return nil, nil
}

func (dao *SwapDao) GetDstTransactionByHash(hash string) (*models.DstTransaction, error) {
	return nil, nil
}

func (dao *SwapDao) WrapperTransactionCheckFee(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction) error {
	return nil
}

func (dao *SwapDao) FillTxSpecialChain(wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction, wrapperDetails []*models.WrapperDetail, polyDetails []*models.PolyDetail) (detailWrapperTxs []*models.WrapperTransaction, err error) {
	return
}

func (dao *SwapDao) GetLatestTx(chainId uint64) (string, string) {
	return "", ""
}