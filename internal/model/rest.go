package model

// CoinItemRequestBody is type CoinItem for receiving through REST
type CoinItemRequestBody struct {
	SymbolID    string `json:"symbol"`
	Name        string `json:"name"`
	TypeID      uint8  `json:"type"`
	ExchangesID []uint `json:"exchanges"`
}

// MessageErr is positive answer in REST
//
// status is false
type MessageErr struct {
	Status  bool   `json:"status" default:"false"`
	Message string `json:"message"`
}

// MessageOk is positive answer in REST
//
// status is true
type MessageOk struct {
	Status  bool      `json:"status" default:"true"`
	Message string    `json:"message"`
	Data    *CoinItem `json:"data"`
}
