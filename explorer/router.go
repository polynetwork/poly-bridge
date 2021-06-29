package explorer

import (
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSRouter("/getCrossTx", &ExplorerController{}, "get:GetCrossTx"),
		beego.NSRouter("/getcrosstxlist/", &ExplorerController{}, "post:GetCrossTxList"),
		beego.NSRouter("/getexplorerinfo/", &ExplorerController{}, "post:GetExplorerInfo"),
		beego.NSRouter("/gettokentxlist/", &ExplorerController{}, "post:GetTokenTxList"),
		beego.NSRouter("/getaddresstxlist/", &ExplorerController{}, "post:GetAddressTxList"),
	)
	beego.AddNamespace(ns)
}
