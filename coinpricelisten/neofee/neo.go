package neofee

import (
	"math/big"
	"poly-swap/chainsdk"
	"poly-swap/conf"
)

type NeoFee struct {
	neoCfg *conf.FeeListenConfig
	neoSdk *chainsdk.NeoSdkPro
}

func NewNeoFee(neoCfg *conf.FeeListenConfig) *NeoFee {
	neoFee := &NeoFee{}
	neoFee.neoCfg = neoCfg
	//
	urls := neoCfg.GetNodesUrl()
	sdk := chainsdk.NewNeoSdkPro(urls)
	neoFee.neoSdk = sdk
	return neoFee
}

func (this *NeoFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	return big.NewInt(1000000000), big.NewInt(1000000000), big.NewInt(1000000000), nil
}

func (this *NeoFee) GetChainId() uint64 {
	return this.neoCfg.ChainId
}
