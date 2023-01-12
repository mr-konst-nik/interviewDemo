package model

import (
	"fmt"
	"sync"
)

// CoinItem is type describe a stock
type CoinItem struct {
	Name        string  `json:"name"`
	TypeID      uint8   `json:"type"`
	ExchangesID []uint8 `json:"exchanges"`
}

// CoinList consists all list of cryptocurrencies
// key is id in string format, like "BTC"
type CoinList struct {
	Value map[string]CoinItem
	sync.Mutex
}

// Init a coin list
func (cl *CoinList) Init() *CoinList {
	return &CoinList{}
}

// Create a item of the coin list
func (cl *CoinList) Create(symbol string, coin CoinItem) {
	cl.Lock()
	cl.Value[symbol] = coin
	cl.Unlock()
}

// Read a item of the coin list
func (cl *CoinList) Read(symbol string) (CoinItem, error) {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.Value[symbol]; !ok {
		return CoinItem{}, fmt.Errorf(symbol, "not found")
	}
	return cl.Value[symbol], nil
}

// Update a item of the coin list by symbol
func (cl *CoinList) Update(symbol string, coin CoinItem) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.Value[symbol]; !ok {
		return fmt.Errorf(symbol, "not found")
	}
	cl.Value[symbol] = coin
	return nil
}

// Delete a item of the coin list by symbol
func (cl *CoinList) Delete(symbol string) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.Value[symbol]; !ok {
		return fmt.Errorf(symbol, "not found")
	}
	delete(cl.Value, symbol)
	return nil
}
