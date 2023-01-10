package basedef

import (
	"fmt"
)

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
	MARKET_COINCHECK     = "coincheck"
	MARKET_GATEIO        = "gateio"
	MARKET_SELF          = "self"
)

const (
	STATE_FINISHED = iota
	STATE_PENDDING
	STATE_SOURCE_DONE
	STATE_SOURCE_CONFIRMED
	STATE_POLY_CONFIRMED
	STATE_DESTINATION_DONE

	STATE_WITHOUT_WRAPPER = 11

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
	StatusOk = "OK"
)

const (
	Chain_Status_All_Nodes_No_Growth        = "All Nodes No Growth"
	Chain_Status_All_Nodes_Unavaiable       = "All Nodes Unavailable"
	Chain_Status_Too_Many_TXs_Stuck         = "Too Many TXs Stuck"
	Chain_Status_All_Relayer_Out_Of_Balance = "All Relayer Out Of Balance"
)

type LargeTx struct {
	Asset     string
	Type      string
	From      string
	To        string
	Amount    string
	USDAmount string
	Hash      string
	User      string
	Time      string
}

type NodeStatus struct {
	ChainId   uint64
	ChainName string
	Url       string
	Height    uint64
	Status    []string
	Time      int64
}

type ChainStatus struct {
	ChainId       uint64
	ChainName     string
	Height        uint64
	StatusTimeMap map[string]int64
	Health        bool
	Time          int64
}

type RelayerAccountStatus struct {
	ChainId   uint64
	ChainName string
	Address   string
	Balance   float64
	Threshold float64
	Status    string
	Time      int64
}

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
