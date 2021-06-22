package switcheofee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type SwitcheoFee struct {
	swthCfg *conf.FeeListenConfig
	swthSdk *chainsdk.SwitcheoSdkPro
}

func NewSwitcheoFee(swthCfg *conf.FeeListenConfig, feeUpdateSlot int64) *SwitcheoFee {
	switcheoFee := &SwitcheoFee{}
	switcheoFee.swthCfg = swthCfg
	urls := swthCfg.GetNodesUrl()
	sdk := chainsdk.NewSwitcheoSdkPro(urls, uint64(feeUpdateSlot), swthCfg.ChainId)
	switcheoFee.swthSdk = sdk
	return switcheoFee
}

func (this *SwitcheoFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice := new(big.Int).SetUint64(0)
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.swthCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.swthCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.swthCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *SwitcheoFee) GetChainId() uint64 {
	return this.swthCfg.ChainId
}

func (this *SwitcheoFee) Name() string {
	return this.swthCfg.ChainName
}
