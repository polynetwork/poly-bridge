package explorer

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/nft_http/controllers"
	"testing"
)

var (
	configFile = "./../../conf/config_testnet.json"
)

func TestMain(m *testing.M) {
	cfg := conf.NewConfig(configFile)

	controllers.Initialize(cfg)
	m.Run()
}

// todo(fuk): these test case are debug only! delete them after test.

func Test_GetExplorerInfo(t *testing.T) {

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	c.GetExplorerInfo()
}

func Test_GetTokenTxList(t *testing.T) {
	req := &models.TokenTxListReq{
		PageSize:10,
		PageNo:0,
		Token:"2",
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetTokenTxList()
}

func Test_GetAddressTxList(t *testing.T) {
	req := &models.AddressTxListReq{
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetAddressTxList()
}

func Test_GetCrossTxList(t *testing.T) {
	req := &models.AddressTxListReq{
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetCrossTxList()
}

func Test_GetCrossTx(t *testing.T) {
	req := &models.AddressTxListReq{
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetCrossTx()
}

func Test_GetAssetStatistic(t *testing.T) {
	req := &models.AddressTxListReq{
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetAssetStatistic()
}

func Test_GetTransferStatistic(t *testing.T) {
	req := &models.AddressTxListReq{
	}

	c := new(ExplorerController)
	c.Controller = web.Controller{}
	assert.NoError(t, encodeContextParams(&c.Controller, req))

	c.GetTransferStatistic()
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
