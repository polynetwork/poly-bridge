package basedef

import "fmt"

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
	MARKET_COINCHECK     = "coincheck"
	MARKET_SELF          = "self"
)

const (
	STATE_FINISHED = iota
	STATE_PENDDING
	STATE_SOURCE_DONE
	STATE_SOURCE_CONFIRMED
	STATE_POLY_CONFIRMED
	STATE_DESTINATION_DONE

	STATE_WAIT = 100
	STATE_SKIP = 101
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

const (
	NATIVE_TOKEN = "0000000000000000000000000000000000000000"
)

func GetStateName(state int) string {
	switch state {
	case STATE_FINISHED:
		return "Finished"
	case STATE_PENDDING:
		return "Pending"
	case STATE_SOURCE_DONE:
		return "SrcDone"
	case STATE_SOURCE_CONFIRMED:
		return "SrcConfirmed"
	case STATE_POLY_CONFIRMED:
		return "PolyConfirmed"
	case STATE_DESTINATION_DONE:
		return "DestDone"
	case STATE_WAIT:
		return "WAIT"
	case STATE_SKIP:
		return "SKIP"
	default:
		return fmt.Sprintf("Unknown(%d)", state)
	}
}
