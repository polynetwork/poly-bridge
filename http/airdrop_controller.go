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

func (c *AddressController) AirDropOfAddress() {
	var addressReq models.AirDropReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &addressReq)
	if err != nil || len(addressReq.Users) == 0 || len(addressReq.Users) > 10 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	evmAddressHashes := make([]string, 0)
	ontAddressHashes := make([]string, 0)
	neoAddressHashes := make([]string, 0)
	neo3AddressHashes := make([]string, 0)
	starcoinAddressHashes := make([]string, 0)

	for _, v := range addressReq.Users {
		if len(v.Address) > 0 {
			if v.ChainId == basedef.ONT_CROSSCHAIN_ID {
				ontAddressHashes = append(ontAddressHashes, v.Address)
			} else if v.ChainId == basedef.NEO_CROSSCHAIN_ID {
				neoAddressHashes = append(neoAddressHashes, v.Address)
			} else if v.ChainId == basedef.NEO3_CROSSCHAIN_ID {
				neo3AddressHashes = append(neo3AddressHashes, v.Address)
			} else if v.ChainId == basedef.STARCOIN_CROSSCHAIN_ID {
				starcoinAddressHashes = append(starcoinAddressHashes, v.Address)
			} else if basedef.IsETHChain(v.ChainId) {
				evmAddressHashes = append(evmAddressHashes, v.Address)
			}
		}
	}
	airDropInfos := make([]*models.AirDropInfo, 0)

	if len(evmAddressHashes) > 0 {
		airDrops := make([]*models.AirDropInfo, 0)
		db.Where("user in ?", evmAddressHashes).
			Find(&airDrops)
		airDropInfos = append(airDropInfos, airDrops...)
	}
	if len(ontAddressHashes) > 0 {
		airDrops := make([]*models.AirDropInfo, 0)
		db.Where("ont_addr in ?", ontAddressHashes).
			Find(&airDrops)
		airDropInfos = append(airDropInfos, airDrops...)
	}
	if len(neoAddressHashes) > 0 {
		airDrops := make([]*models.AirDropInfo, 0)
		db.Where("neo_addr in ?", neoAddressHashes).
			Find(&airDrops)
		airDropInfos = append(airDropInfos, airDrops...)
	}
	if len(neo3AddressHashes) > 0 {
		airDrops := make([]*models.AirDropInfo, 0)
		db.Where("neo3_addr in ?", neo3AddressHashes).
			Find(&airDrops)
		airDropInfos = append(airDropInfos, airDrops...)
	}
	if len(starcoinAddressHashes) > 0 {
		airDrops := make([]*models.AirDropInfo, 0)
		db.Where("star_coin_addr in ?", starcoinAddressHashes).
			Find(&airDrops)
		airDropInfos = append(airDropInfos, airDrops...)
	}

	c.Data["json"] = models.MakeAirDropRsp(addressReq, airDropInfos)
	c.ServeJSON()
}
