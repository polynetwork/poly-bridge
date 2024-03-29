package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"poly-bridge/basedef"
	"poly-bridge/models"
)

type AirDropController struct {
	web.Controller
}

func (c *AirDropController) AirDropOfAddress() {
	var addressReq models.AirDropReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &addressReq)
	if err != nil || len(addressReq.Users) == 0 || len(addressReq.Users) > 10 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	ethAddrs := make([]string, 0)
	otherUsers := make([]string, 0)
	for _, v := range addressReq.Users {
		if len(v.Address) > 0 {
			if basedef.IsETHChain(v.ChainId) {
				ethAddrs = append(ethAddrs, v.Address)
			} else {
				otherUsers = append(otherUsers, v.Address)
			}
		}
	}

	if len(ethAddrs) == 0 && len(otherUsers) == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	airDropInfosUser := make([]*models.AirDropInfo, 0)
	db.Where("user in ?", otherUsers).
		Find(&airDropInfosUser)
	airDropInfosAddr := make([]*models.AirDropInfo, 0)
	db.Where("bind_addr in ?", ethAddrs).
		Find(&airDropInfosAddr)
	airDropInfos := make([]*models.AirDropInfo, 0)
	airDropInfos = append(airDropInfos, airDropInfosUser...)
	airDropInfos = append(airDropInfos, airDropInfosAddr...)

	bindAddrs := make([]string, 0)
	for _, v := range airDropInfos {
		bindAddrs = append(bindAddrs, v.BindAddr)
	}

	airDropRanks := make([]*models.AirDropRank, 0)
	db.Table("(?) as b", db.Table("(select @curRank := 0) as r, (?) as t",
		db.Model(&models.AirDropInfo{}).Select("sum(amount) as sum_amount, bind_addr").Group("bind_addr").Order("sum_amount desc, bind_addr")).Select("t.sum_amount as amount,t.bind_addr,@curRank := @curRank + 1 as rank")).
		Where("b.bind_addr in ?", bindAddrs).
		Find(&airDropRanks)

	c.Data["json"] = models.MakeAirDropRsp(addressReq, airDropInfos, airDropRanks)
	c.ServeJSON()
}

func (c *AirDropController) AirDropClaim() {
	var addressReq models.AirDropClaimReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &addressReq)
	if err != nil || len(addressReq.AirDropAddrs) == 0 || len(addressReq.AirDropAddrs) > 10 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	airDropNfts := make([]*models.AirDropNft, 0)
	chain := basedef.ETHEREUM_CROSSCHAIN_ID
	if basedef.ENV == basedef.TESTNET {
		chain = basedef.RINKEBY_CROSSCHAIN_ID
	}
	db.Where("bind_addr in ? and bind_chain_id = ? and rank <=1000 ", addressReq.AirDropAddrs, chain).
		Find(&airDropNfts)
	airDropClaimRsp, claimFlag := models.MakeAirDropClaimRsp(airDropNfts)
	if len(claimFlag) > 0 {
		for i, v := range airDropNfts {
			if claimFlag[i] {
				db.Save(v)
			}
		}
	}
	c.Data["json"] = airDropClaimRsp
	c.ServeJSON()
}
