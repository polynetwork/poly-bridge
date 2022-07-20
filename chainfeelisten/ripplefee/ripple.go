package ripplefee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type RippleFee struct {
	rippleCfg *conf.FeeListenConfig
	rippleSdk *chainsdk.RippleSdkPro
}

func NewRippleFee(rippleCfg *conf.FeeListenConfig, feeUpdateSlot int64) *RippleFee {
	RippleFee := &RippleFee{}
	RippleFee.rippleCfg = rippleCfg
	urls := rippleCfg.GetNodesUrl()
	sdk := chainsdk.NewRippleSdkPro(urls, uint64(feeUpdateSlot), rippleCfg.ChainId)
	RippleFee.rippleSdk = sdk
	return RippleFee
}

func (this *RippleFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice, err := this.rippleSdk.GetFee()
	if err != nil {
		return nil, nil, nil, err
	}

	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.rippleCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.rippleCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.rippleCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *RippleFee) GetChainId() uint64 {
	return this.rippleCfg.ChainId
}

func (this *RippleFee) Name() string {
	return this.rippleCfg.ChainName
}
