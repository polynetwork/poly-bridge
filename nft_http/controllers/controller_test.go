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

func encodeContextParams(c *beego.Controller, request interface{}) error {
	c.Ctx = new(context.Context)
	c.Ctx.Input = new(context.BeegoInput)

	enc, err := json.Marshal(request)
	if err != nil {
		return err
	}
	c.Ctx.Input.RequestBody = enc
	return nil
}
