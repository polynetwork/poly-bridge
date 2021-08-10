package http

import (
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/common"
	"poly-bridge/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type TokenAssetController struct {
	web.Controller
}

func (c *TokenAssetController) Gettokenasset() {
	var tokenAssetReq models.TokenAssetReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenAssetReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	if len(tokenAssetReq.NameOrHash) == 40 {
		tokens := make([]*models.Token, 0)
		err := db.Where("hash=? and property = ?", tokenAssetReq.NameOrHash, 1).
			Find(&tokens).Error
		if err != nil {
			c.Data["json"] = ""
			c.ServeJSON()
		}
		for _, token := range tokens {
			if token.Name == "" {
				continue
			}
			tokenBasic := new(models.TokenBasic)
			err = db.Where("name = ? and property = ?", token.TokenBasicName, 1).
				Preload("Tokens").
				First(tokenBasic).Error
			if err != nil {
				continue
			}
			tokenBasics = append(tokenBasics, tokenBasic)
		}
	} else {
		tokenBasic := new(models.TokenBasic)
		err = db.Where("name = ? and property = ?", tokenAssetReq.NameOrHash, 1).
			Preload("Tokens").
			First(tokenBasic).Error
		if err != nil {
			c.Data["json"] = ""
			c.ServeJSON()
		}
		tokenBasics = append(tokenBasics, tokenBasic)
	}

	assetDetails := make([]*models.AssetDetailRes, 0)
	for _, basic := range tokenBasics {
		dstChainAssets := make([]*models.DstChainAssetRes, 0)
		for _, token := range basic.Tokens {
			chainAsset := new(models.DstChainAssetRes)
			chainFee := new(models.ChainFee)
			err = db.Where("chain_id = ?", token.ChainId).
				First(chainFee).Error
			if err != nil {
				chainAsset.ChainName = ""
			} else {
				chainAsset.ChainName = chainFee.TokenBasicName
			}
			chainAsset.Hash = token.Hash
			balance, err := common.GetBalance(token.ChainId, token.Hash)
			if err != nil {
				chainAsset.ErrReason = err.Error()
				logs.Info("chainId: %v, Hash: %v, err:%v", token.ChainId, token.Hash, err)
				balance = big.NewInt(-1)
			}
			chainAsset.Balance = balance
			time.Sleep(time.Second)
			totalSupply, _ := common.GetTotalSupply(token.ChainId, token.Hash)
			if err != nil {
				chainAsset.ErrReason = err.Error()
				totalSupply = big.NewInt(-1)
				logs.Info("chainId: %v, Hash: %v, err:%v ", token.ChainId, token.Hash, err)
			}
			chainAsset.TotalSupply = totalSupply
			dstChainAssets = append(dstChainAssets, chainAsset)
		}
		assetDetail := new(models.AssetDetailRes)
		assetDetail.TokenAsset = dstChainAssets
		assetDetail.BasicName = basic.Name
		assetDetail.Precision = basic.Precision
		assetDetails = append(assetDetails, assetDetail)
	}
	c.Data["json"] = assetDetails
	c.ServeJSON()
}
