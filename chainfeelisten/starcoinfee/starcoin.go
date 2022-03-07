package starcoinfee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type StarcoinFee struct {
	starcoinCfg *conf.FeeListenConfig
	starcoinSdk *chainsdk.StarcoinSdkPro
}

func NewStarcoinFee(starcoinCfg *conf.FeeListenConfig, feeUpdateSlot int64) *StarcoinFee {
	StarcoinFee := &StarcoinFee{}
	StarcoinFee.starcoinCfg = starcoinCfg
	urls := starcoinCfg.GetNodesUrl()
	sdk := chainsdk.NewStarcoinSdkPro(urls, uint64(feeUpdateSlot), starcoinCfg.ChainId)
	StarcoinFee.starcoinSdk = sdk
	return StarcoinFee
}

func (this *StarcoinFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	suggestGasPrice, err := this.starcoinSdk.GetGasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice := big.NewInt(int64(suggestGasPrice))

	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.starcoinCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.starcoinCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.starcoinCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *StarcoinFee) GetChainId() uint64 {
	return this.starcoinCfg.ChainId
}

func (this *StarcoinFee) Name() string {
	return this.starcoinCfg.ChainName
}
