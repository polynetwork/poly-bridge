package chainlisten

import (
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/conf"
)

func StartChainListen(cfg *conf.ChainListenConfig, dbCfg *conf.DBConfig) {
	ethListen := ethereumlisten.NewEthereumChainListen(cfg.EthereumChainListenConfig, dbCfg)
	ethListen.Start()
}
