package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"poly-swap/models"
)

type AddressController struct {
	beego.Controller
}

func (c *AddressController) Address() {
	var addressReq models.AddressReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &addressReq); err != nil {
		panic(err)
	}

	c.Data["json"] = models.MakeAddressRsp(addressReq.AddressHash, addressReq.ChainId, address(addressReq.AddressHash, addressReq.ChainId))
	c.ServeJSON()
}

func address(hash string, chainId uint64) string {
	return hash
}