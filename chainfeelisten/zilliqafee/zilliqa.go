package zilliqafee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/utils/decimal"
)

type ZilliqaFee struct {
	zilCfg *conf.FeeListenConfig
	zilSdk *chainsdk.ZilliqaSdkPro
}

func NewZilliqaFee(zilCfg *conf.FeeListenConfig, feeUpdateSlot int64) *ZilliqaFee {
	zilFee := &ZilliqaFee{}
	zilFee.zilCfg = zilCfg
	sdk := chainsdk.NewZilliqaSdkPro(zilCfg.Nodes, uint64(feeUpdateSlot), zilCfg.ChainId)
	zilFee.zilSdk = sdk
	return zilFee
}

func (this *ZilliqaFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	//zil 10^-12
	gasPrice, err := this.zilSdk.GetMinimumGasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice_real, err := decimal.NewFromString(gasPrice)
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice_real = gasPrice_real.Div(decimal.New(1, 12))
	gasPrice_new := gasPrice_real.Mul(decimal.NewFromInt(basedef.FEE_PRECISION)).Mul(decimal.NewFromInt(this.zilCfg.GasLimit))
	proxyFee := gasPrice_new.Mul(decimal.NewFromInt(this.zilCfg.ProxyFee)).Div(decimal.NewFromInt32(100))
	minFee := gasPrice_new.Mul(decimal.NewFromInt(this.zilCfg.MinFee)).Div(decimal.NewFromInt32(100))
	return minFee.BigInt(), gasPrice_new.BigInt(), proxyFee.BigInt(), nil
}

func (this *ZilliqaFee) GetChainId() uint64 {
	return this.zilCfg.ChainId
}

func (this *ZilliqaFee) Name() string {
	return this.zilCfg.ChainName
}
