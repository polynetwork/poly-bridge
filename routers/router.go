package routers

import (
	"poly-swap/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),
		beego.NSRouter("/token/", &controllers.TokenController{}, "post:Token"),
		beego.NSRouter("/tokens/", &controllers.TokenController{}, "post:Tokens"),
		beego.NSRouter("/tokenmap/", &controllers.TokenMapController{}, "post:TokenMap"),
		beego.NSRouter("/getfee/", &controllers.FeeController{}, "post:GetFee"),
		beego.NSRouter("/checkfee/", &controllers.FeeController{}, "post:CheckFee"),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.InfoController{}, "*:Get")
}
