package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"poly-swap/models"
)

type TokenController struct {
	beego.Controller
}

func (c *TokenController) Tokens() {
	var tokensReq models.TokensReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokensReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokens := make([]*models.Token, 0)
	db.Where("chain_id = ?", tokensReq.ChainId).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").Find(&tokens)
	c.Data["json"] = models.MakeTokensRsp(tokens)
	c.ServeJSON()
}

func (c *TokenController) Token() {
	var tokenReq models.TokenReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenReq); err != nil {
		panic(err)
	}
	db := newDB()
	token := new(models.Token)
	db.Where("hash = ?", tokenReq.Hash).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").First(token)
	c.Data["json"] = models.MakeTokenRsp(token)
	c.ServeJSON()
}
