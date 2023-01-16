package aptosfee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type AptosFee struct {
	aptosCfg *conf.FeeListenConfig
	aptosSdk *chainsdk.AptosSdkPro
}

func NewAptosFee(aptosCfg *conf.FeeListenConfig, feeUpdateSlot int64) *AptosFee {
	aptosFee := &AptosFee{}
	aptosFee.aptosCfg = aptosCfg
	sdk := chainsdk.NewAptosSdkPro(aptosCfg.Nodes, uint64(feeUpdateSlot), aptosCfg.ChainId)
	aptosFee.aptosSdk = sdk
	return aptosFee
}

func (this *AptosFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	suggestGasPrice, err := this.aptosSdk.GetGasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice := big.NewInt(int64(suggestGasPrice))

	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.aptosCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.aptosCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.aptosCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *AptosFee) GetChainId() uint64 {
	return this.aptosCfg.ChainId
}

func (this *AptosFee) Name() string {
	return this.aptosCfg.ChainName
}
