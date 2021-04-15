package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/stretchr/testify/assert"
	"poly-bridge/conf"
	"testing"
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
		Size: 10,
	}

	c := new(InfoController)
	c.Controller = beego.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.Home()
}

func TestItemController_Items(t *testing.T) {
	req := ItemsOfAddressReq{
		ChainId:  79,
		Asset:    "66638F4970C2ae63773946906922c07a583b6069",
		Address:  "5Fb03EB21303D39967a1a119B32DD744a0fA8986",
		TokenId:  "",
		PageSize: 0,
		PageNo:   10,
	}

	c := new(ItemController)
	bc := new(beego.Controller)
	c.Controller = *bc
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.Items()
}

func encodeContextParams(c *beego.Controller, request interface{}) error {
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
