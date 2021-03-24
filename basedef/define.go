package basedef

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
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
	SERVER_POLY_SWAP = "polyswap"
	SERVER_EXPLORER  = "explorer"
	SERVER_ADDRESS   = "address"
	SERVER_STAKE     = "stake"
)

const (
	ADDRESS_LENGTH = 64
)
