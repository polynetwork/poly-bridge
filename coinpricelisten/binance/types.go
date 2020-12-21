package binance

type Ticker struct {
	Symbol            string                  `json:"symbol"`
	Price             float64                 `json:"price,string"`
}
