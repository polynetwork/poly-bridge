package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"poly-bridge/models"
)

type TokenAssetController struct {
	beego.Controller
}

func (c *TokenAssetController) Gettokenasset() {
	var expectTimeReq models.TokenAssetReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &expectTimeReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	//tokenBasics := make([]*models.TokenBasic, 0)
	//if len(expectTimeReq.NameOrHash)==40{
	//
	//}

	//resAssetDetails := make([]*AssetDetail, 0)
	//extraAssetDetails := make([]*AssetDetail, 0)
	//tokenBasics := make([]*models.TokenBasic, 0)
	//res := db.
	//	Where("property = ?", 1).
	//	Preload("Tokens").
	//	Find(&tokenBasics)
	//if res.Error != nil {
	//	return err
	//}

}
