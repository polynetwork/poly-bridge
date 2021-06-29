package explorer

import (
	"github.com/beego/beego/v2/server/web"
)

func GetRouter() web.LinkNamespace {
	ns := web.NSNamespace("/explorer",
		web.NSRouter("/getCrossTx", &ExplorerController{}, "get:GetCrossTx"),
		web.NSRouter("/getcrosstxlist/", &ExplorerController{}, "post:GetCrossTxList"),
		web.NSRouter("/getexplorerinfo/", &ExplorerController{}, "post:GetExplorerInfo"),
		web.NSRouter("/gettokentxlist/", &ExplorerController{}, "post:GetTokenTxList"),
		web.NSRouter("/getaddresstxlist/", &ExplorerController{}, "post:GetAddressTxList"),
	)
	return ns
}
