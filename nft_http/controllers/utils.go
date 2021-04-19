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

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/nft_http/meta"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db           *gorm.DB
	txCounter    *TransactionCounter
	sdks         = make(map[uint64]*chainsdk.EthereumSdkPro)
	assets       = make([]*models.Token, 0)
	wrapperAddrs = make(map[uint64]common.Address)
	fetcher      *meta.StoreFetcher
	feeTokens    = make(map[uint64]*models.Token)
	lruDB        *lru.ARCCache
)

func NewDB(cfg *conf.DBConfig) *gorm.DB {
	user := cfg.User
	password := cfg.Password
	url := cfg.URL
	scheme := cfg.Scheme
	Logger := logger.Default
	if cfg.Debug {
		Logger = Logger.LogMode(logger.Info)
	}
	format := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, url, scheme)
	db, err := gorm.Open(mysql.Open(format), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	db.Where("standard = ? and property=?", models.TokenTypeErc721, 1).
		Preload("TokenBasic").
		Find(&assets)
	for _, v := range assets {
		logs.Info("load asset %s, chainid %d, hash %s", v.TokenBasicName, v.ChainId, v.Hash)
	}

	feeTokenList := make([]*models.Token, 0)
	db.Where("hash=?", nativeHash).
		Preload("TokenBasic").
		Find(&feeTokenList)
	for _, v := range feeTokenList {
		feeTokens[v.ChainId] = v
		logs.Info("load chainid %d feeToken %s", v.ChainId, v.TokenBasicName)
	}
	return db
}

func Initialize(c *conf.Config) {
	db = NewDB(c.DBConfig)

	for _, v := range c.ChainListenConfig {
		urls := v.GetNodesUrl()
		if len(urls) > 0 {
			pro := chainsdk.NewEthereumSdkPro(v.GetNodesUrl(), v.ListenSlot, v.ChainId)
			sdks[v.ChainId] = pro
			wrapperAddrs[v.ChainId] = common.HexToAddress(v.NFTWrapperContract)
			logs.Info("load chain id %d, contract %s", v.ChainId, v.NFTWrapperContract)
		} else {
			logs.Warn("chain %s node is empty", v.ChainName)
		}
	}

	arcLRU, err := lru.NewARC(5000)
	if err != nil {
		panic(err)
	}
	lruDB = arcLRU

	fetcher = meta.NewStoreFetcher(db)
	for _, asset := range assets {
		if asset.TokenBasic == nil {
			continue
		}
		fetcherTyp := meta.FetcherType(asset.TokenBasic.MetaFetcherType)
		baseUri := asset.TokenBasic.Meta
		assetName := asset.TokenBasic.Name
		fetcher.Register(fetcherTyp, assetName, baseUri)
	}

	txCounter = NewTransactionCounter()
}

type TransactionCounter struct {
	Count    int64
	LastTime int64
}

func NewTransactionCounter() *TransactionCounter {
	s := new(TransactionCounter)
	s.refresh()
	return s
}

func (s *TransactionCounter) refresh() {
	db.Model(&models.WrapperTransaction{}).
		Where("standard = ?", models.TokenTypeErc721).
		Count(&s.Count)

	s.LastTime = time.Now().Unix()
}

func (s *TransactionCounter) Number() int64 {
	now := time.Now().Unix()
	if now-s.LastTime < 120 {
		return s.Count
	}

	s.refresh()
	return s.Count
}

func selectNode(chainID uint64) *chainsdk.EthereumSdkPro {
	pro, ok := sdks[chainID]
	if !ok {
		return nil
	}
	return pro
}

var emptyAddr = common.Address{}

func selectWrapper(chainID uint64) common.Address {
	addr, ok := wrapperAddrs[chainID]
	if !ok {
		return emptyAddr
	}
	return addr
}

func selectNFTAsset(addr string) *models.Token {
	for _, v := range assets {
		origin := common.HexToAddress(v.Hash)
		src := common.HexToAddress(addr)
		if bytes.Equal(origin.Bytes(), src.Bytes()) {
			return v
		}
	}
	return nil
}

func selectAssetsByChainId(chainId uint64) []*models.Token {
	res := make([]*models.Token, 0)
	for _, v := range assets {
		if v.ChainId == chainId {
			res = append(res, v)
		}
	}
	return res
}

const (
	ErrCodeRequest     int = 400
	ErrCodeNotExist    int = 404
	ErrCodeNodeInvalid int = 500
)

var errMap = map[int]string{
	ErrCodeRequest:     "request parameter is invalid!",
	ErrCodeNotExist:    "not found",
	ErrCodeNodeInvalid: "blockchain node exception",
}

func input(c *beego.Controller, req interface{}) bool {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		code := ErrCodeRequest
		customInput(c, code, errMap[code])
		return false
	} else {
		return true
	}
}

func customInput(c *beego.Controller, code int, msg string) {
	c.Data["json"] = models.MakeErrorRsp(msg)
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func notExist(c *beego.Controller) {
	code := ErrCodeNotExist
	c.Data["json"] = models.MakeErrorRsp(errMap[code])
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func nodeInvalid(c *beego.Controller) {
	code := ErrCodeNodeInvalid
	c.Data["json"] = models.MakeErrorRsp(errMap[code])
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func output(c *beego.Controller, data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
}

func customOutput(c *beego.Controller, code int, msg string) {
	c.Data["json"] = models.MakeErrorRsp(msg)
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func getPageNo(totalNo, pageSize int) int {
	return (int(totalNo) + pageSize - 1) / pageSize
}

func findFeeToken(cid uint64, hash string) *models.Token {
	feeTokens := make([]*models.Token, 0)
	db.Model(&models.Token{}).
		Where("hash = ?", nativeHash).
		Preload("TokenBasic").
		Find(&feeTokens)

	for _, v := range feeTokens {
		if cid == v.ChainId && hash == v.Hash {
			return v
		}
	}
	return nil
}

func findAsset(cid uint64, hash string) *models.Token {
	for _, v := range assets {
		if v.ChainId == cid && hash == v.Hash {
			return v
		}
	}
	return nil
}
