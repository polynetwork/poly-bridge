package routers

import (
	"github.com/astaxie/beego"
	"poly-swap/conf"
	"poly-swap/controllers"
)

func init() {
	server := beego.AppConfig.String("server")
	if server == conf.SERVER_POLY_SWAP {
		ns := beego.NewNamespace("/v1",
			beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),
			beego.NSRouter("/token/", &controllers.TokenController{}, "post:Token"),
			beego.NSRouter("/tokens/", &controllers.TokenController{}, "post:Tokens"),
			beego.NSRouter("/tokenmap/", &controllers.TokenMapController{}, "post:TokenMap"),
			beego.NSRouter("/getfee/", &controllers.FeeController{}, "post:GetFee"),
			beego.NSRouter("/checkfee/", &controllers.FeeController{}, "post:CheckFee"),
			beego.NSRouter("/transactoins/", &controllers.TransactionController{}, "post:Transactions"),
			beego.NSRouter("/transactoinsofuser/", &controllers.TransactionController{}, "post:TransactoinsOfUser"),
		)
		beego.AddNamespace(ns)
		beego.Router("/", &controllers.InfoController{}, "*:Get")
	} else {
		ns := beego.NewNamespace("/v1",
			beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),
			beego.NSRouter("/address/", &controllers.AddressController{}, "post:Address"),
		)
		beego.AddNamespace(ns)
		beego.Router("/", &controllers.InfoController{}, "*:Get")
	}
}
