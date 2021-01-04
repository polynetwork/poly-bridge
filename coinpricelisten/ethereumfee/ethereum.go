package ethereumfee

import (
	"math/big"
	"poly-swap/chainsdk"
	"poly-swap/conf"
)

type EthereumFee struct {
	ethCfg *conf.FeeListenConfig
	ethSdk *chainsdk.EthereumSdkPro
}

func NewEthereumFee(ethCfg *conf.FeeListenConfig) *EthereumFee {
	ethereumFee := &EthereumFee{}
	ethereumFee.ethCfg = ethCfg
	//
	urls := ethCfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls)
	ethereumFee.ethSdk = sdk
	return ethereumFee
}

func (this *EthereumFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice, err := this.ethSdk.SuggestGasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(conf.PRICE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.ethCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.ethCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	return gasPrice, gasPrice, proxyFee, nil
}

func (this *EthereumFee) GetChainId() uint64 {
	return this.ethCfg.ChainId
}
