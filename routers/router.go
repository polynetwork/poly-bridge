package routers

import (
	"github.com/astaxie/beego"
	"poly-swap/controllers"
)

func init() {
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
}
