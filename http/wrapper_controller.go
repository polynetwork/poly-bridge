package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"poly-bridge/models"
)

type WrapperController struct {
	web.Controller
}

func (c *WrapperController) WrapperCheck() {
	var wrapperCheckReq models.WrapperCheckReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &wrapperCheckReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	wrappers, ok := contractCheck[wrapperCheckReq.ChainId]
	if !ok {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("no this chain!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	wrapperCheckRsp := &models.WrapperCheckRsp{
		wrapperCheckReq.ChainId,
		wrappers,
	}
	c.Data["json"] = wrapperCheckRsp
	c.ServeJSON()
}
