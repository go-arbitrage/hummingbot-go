package hummingbot_go

type PriceRequest struct {
	Chain     string `json:"chain"`
	Network   string `json:"network"`
	Connector string `json:"connector"`
	Base      string `json:"base"`
	Quote     string `json:"quote"`
	Amount    string `json:"amount"`
	Side      string `json:"side"`
}
