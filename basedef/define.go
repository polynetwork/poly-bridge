package basedef

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
	MARKET_SELF          = "self"
)

const (
	STATE_FINISHED = iota
	STATE_PENDDING
	STATE_SOURCE_DONE
	STATE_SOURCE_CONFIRMED
	STATE_POLY_CONFIRMED
	STATE_DESTINATION_DONE
)

const (
	SERVER_POLY_BRIDGE = "polybridge"
	SERVER_POLY_SWAP   = "polyswap"
	SERVER_EXPLORER    = "explorer"
	SERVER_ADDRESS     = "address"
	SERVER_STAKE       = "stake"
)

const (
	MAINNET = "mainnet"
	TESTNET = "testnet"
	DEVNET  = "devnet"
)

const (
	ADDRESS_LENGTH = 64
)

const (
	SWAP_SWAP = iota
	SWAP_ADDLIQUIDITY
	SWAP_REMOVELIQUIDITY
	SWAP_ROLLBACK
)

var (
	POLY_CROSSCHAIN_ID     = uint64(0)
	BTC_CROSSCHAIN_ID      = uint64(1)
	ETHEREUM_CROSSCHAIN_ID = uint64(2)
	ONT_CROSSCHAIN_ID      = uint64(3)
	NEO_CROSSCHAIN_ID      = uint64(4)
	BSC_CROSSCHAIN_ID      = uint64(6)
	HECO_CROSSCHAIN_ID     = uint64(7)
	O3_CROSSCHAIN_ID       = uint64(80)
	NEO3_CROSSCHAIN_ID     = uint64(88)
	OK_CROSSCHAIN_ID       = uint64(90)
	MATIC_CROSSCHAIN_ID    = uint64(13)
	SWITCHEO_CROSSCHAIN_ID = uint64(1000) // No testnet for cosmos

	ENV = "devnet"
)

func Init() {
	if Environment == "testnet" {
		logs.Info("this is testnet")
		fmt.Println("this is testnet")
		POLY_CROSSCHAIN_ID = uint64(0)
		ETHEREUM_CROSSCHAIN_ID = uint64(2)
		ONT_CROSSCHAIN_ID = uint64(3)
		NEO_CROSSCHAIN_ID = uint64(5)
		HECO_CROSSCHAIN_ID = uint64(7)
		BSC_CROSSCHAIN_ID = uint64(79)
		O3_CROSSCHAIN_ID = uint64(82)
		NEO3_CROSSCHAIN_ID = uint64(88)
		OK_CROSSCHAIN_ID = uint64(200)
		MATIC_CROSSCHAIN_ID = uint64(202)

		ENV = "testnet"
	} else if Environment == "mainnet" {
		logs.Info("this is mainnet")
		fmt.Println("this is mainnet")
		POLY_CROSSCHAIN_ID = uint64(0)
		BTC_CROSSCHAIN_ID = uint64(1)
		ETHEREUM_CROSSCHAIN_ID = uint64(2)
		ONT_CROSSCHAIN_ID = uint64(3)
		NEO_CROSSCHAIN_ID = uint64(4)
		SWITCHEO_CROSSCHAIN_ID = uint64(5)
		BSC_CROSSCHAIN_ID = uint64(6)
		HECO_CROSSCHAIN_ID = uint64(7)
		O3_CROSSCHAIN_ID = uint64(10)
		OK_CROSSCHAIN_ID = uint64(12)
		MATIC_CROSSCHAIN_ID = uint64(17)
		NEO3_CROSSCHAIN_ID = uint64(88)

		ENV = "mainnet"
	}
}
