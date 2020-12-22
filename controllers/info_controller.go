package controllers

import (
	"github.com/astaxie/beego"
	"poly-swap/models"
)

type InfoController struct {
	beego.Controller
}

func (c *InfoController) Get() {
	explorer := &models.PolySwapResp{
		Version: "v1",
		URL:     "http://localhost:8080/v1",
	}
	c.Data["json"] = explorer
	c.ServeJSON()
}
