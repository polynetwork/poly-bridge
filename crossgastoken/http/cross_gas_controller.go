package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/crossgastoken/gasmodels"
	"poly-bridge/models"
)

type CrossNativeController struct {
	web.Controller
}

func (c *CrossNativeController) result400(message string) {
	c.Data["json"] = message
	c.Ctx.ResponseWriter.WriteHeader(400)
}

func (c *CrossNativeController) getchain() {
	crossNativeMaps:=make([]*gasmodels.CrossGasMap,0)
	err:=db.Where("property = ?",1).
		Order("src_chain_id").
		Preload("Token").
		Find(crossNativeMaps).Error
	if err!=nil{
		c.result400("hasn't cross token")
		return
	}
	c.Data["json"] = MakeGetChainResp(crossNativeMaps)
	c.ServeJSON()
}

func (c *CrossNativeController) getfee() {
	var gasFeeReq gasmodels.GasFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &gasFeeReq); err != nil {
		c.result400("request parameter is invalid!")
		return
	}
	token := new(gasmodels.GasToken)
	res := db.Where("hash = ? and chain_id = ?", gasFeeReq.SrcTokenHash, gasFeeReq.SrcChainId).First(token)
	if res.RowsAffected == 0 {
		c.result400(fmt.Sprintf("chain: %d does not have token: %s", gasFeeReq.SrcChainId, gasFeeReq.SrcTokenHash))
		return
	}
	//todo
	dstfee:=big.NewInt(0)


}