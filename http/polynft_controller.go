package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"poly-bridge/models"
)

type PolyNftController struct {
	web.Controller
}

func (c *PolyNftController) GetNftSign() {
	var nftSignReq models.NftSignReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftSignReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	if nftSignReq.Address[:2] == "0x" || nftSignReq.Address[:2] == "0X" {
		nftSignReq.Address = nftSignReq.Address[2:]
	}
	nftUsers := make([]*models.NftUser, 0)
	colUser := new(models.NftUser)
	res := db.Where("col_address = ? ", nftSignReq.Address).
		First(colUser)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("%v does not exist", nftSignReq.Address))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	nftUsers = append(nftUsers, colUser)
	dfUser := new(models.NftUser)
	res = db.Where("df_address = ? ", nftSignReq.Address).
		First(dfUser)
	if res.RowsAffected > 0 {
		nftUsers = append(nftUsers, dfUser)
	}
	c.Data["json"] = nftUsers
	c.ServeJSON()
}