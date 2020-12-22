package chainlisten

import (
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/chainlisten/neolisten"
	"poly-swap/conf"
)

func StartChainListen(cfg *conf.ChainListenConfig, dbCfg *conf.DBConfig) {
	ethListen := ethereumlisten.NewEthereumChainListen(cfg.EthereumChainListenConfig, dbCfg)
	ethListen.Start()
	neoListen := neolisten.NewNeoChainListen(cfg.NeoChainListenConfig, dbCfg)
	neoListen.Start()
}
