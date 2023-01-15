package model

// CoinItem is type describe a cryptocurrency
type CoinItem struct {
	Name        string `json:"name"`
	TypeID      uint8  `json:"type"`
	ExchangesID []uint `json:"exchanges"`
}
