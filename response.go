package hummingbot_go

type PriceResponse struct {
	Network   string  `json:"network"`
	Timestamp int64   `json:"timestamp"`
	Latency   float64 `json:"latency"`

	Base           string `json:"base"`
	Quote          string `json:"quote"`
	Amount         string `json:"amount"`
	RawAmount      string `json:"rawAmount"`
	ExpectedAmount string `json:"expectedAmount"`

	Price string `json:"price"`

	GasPrice      float64 `json:"gasPrice,omitempty"`
	GasPriceToken *string `json:"gasPriceToken,omitempty"`
	GasLimit      *int64  `json:"gasLimit,omitempty"`
	GasCost       *string `json:"gasCost,omitempty"`
}
