package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
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
	addressHashes := make([]string, 0)
	for _, v := range addressReq.Users {
		if len(v.Address) > 0 {
			addressHashes = append(addressHashes, v.Address)
		}
	}

	if len(addressHashes) == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	airDropInfos := make([]*models.AirDropInfo, 0)
	db.Where("user in ?", addressHashes).
		Find(&airDropInfos)

	airDropRanks := make([]*models.AirDropRank, 0)
	for _, v := range airDropInfos {
		var amount int64
		db.Model(&models.AirDropInfo{}).Select("sum(amount)").Where("bind_addr = ?", v.User).
			Scan(&amount)
		var rank int64
		db.Raw("select count(*) from (select sum(amount) as sumamount, user from air_drop_infos group by bind_addr) z where z.sumamount >= ?", amount).
			Scan(&rank)
		airDropRanks = append(airDropRanks, &models.AirDropRank{
			v.User,
			v.ChainID,
			v.BindAddr,
			v.BindChainId,
			v.Amount,
			v.IsClaim,
			rank,
		})
	}

	c.Data["json"] = models.MakeAirDropRsp(addressReq, airDropRanks)
	c.ServeJSON()
}
