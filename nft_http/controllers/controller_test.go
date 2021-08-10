package controllers

import (
	"encoding/json"
	"poly-bridge/conf"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

var (
	configFile = "./../../conf/config_testnet.json"
)

func TestMain(m *testing.M) {
	cfg := conf.NewConfig(configFile)

	Initialize(cfg)
	m.Run()
}

// todo(fuk): these test case are debug only! delete them after test.

func TestInfoController_Home(t *testing.T) {
	req := &HomeReq{
		ChainId: 79,
		Size:    10,
	}

	c := new(InfoController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.Home()
}

func TestItemController_Items(t *testing.T) {
	req := ItemsOfAddressReq{
		ChainId:  2,
		Asset:    "a85c9fc8f2c9060d674e0ca97f703a0a30619305",
		Address:  "f1c7203ef81fb9663babd8516ebd30d33ee84ee8",
		TokenId:  "",
		PageSize: 12,
		PageNo:   0,
	}

	c := new(ItemController)
	bc := new(web.Controller)
	c.Controller = *bc
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.Items()
}

func encodeContextParams(c *web.Controller, request interface{}) error {
	c.Ctx = new(context.Context)
	c.Ctx.Input = new(context.BeegoInput)
	c.Ctx.Output = new(context.BeegoOutput)
	c.Data = make(map[interface{}]interface{})
	c.Ctx.ResponseWriter = new(context.Response)
	enc, err := json.Marshal(request)
	if err != nil {
		return err
	}
	c.Ctx.Input.RequestBody = enc
	return nil
}
