package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"poly-swap/models"
)

type TokenMapController struct {
	beego.Controller
}

func (c *TokenMapController) TokenMap() {
	var tokenMapReq models.TokenMapReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenMapReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokenMap := new(models.TokenMap)
	db.Where("src_token_hash = ?", tokenMapReq.Hash).Preload("SrcToken").Preload("DstToken").First(tokenMap)
	c.Data["json"] = models.MakeTokenMapRsp(tokenMap)
	c.ServeJSON()
}
