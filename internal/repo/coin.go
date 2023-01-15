package repo

import (
	"errors"
	"fmt"
	"interviewDemo/internal/model"
	"sync"
)

// ErrNotFound error if symbol not found
var ErrNotFound = errors.New("not found")

// ErrIsExist error is symbol already exists
var ErrIsExist = errors.New("already exists")

// CoinList consists all list of cryptocurrencies
// key is id in string format, like "BTC"
type CoinList struct {
	m map[string]*model.CoinItem
	sync.Mutex
}

// NewCoinList create a new empty list of coins
func NewCoinList() *CoinList {
	return &CoinList{
		m: map[string]*model.CoinItem{},
	}
}

// CoinListRepo interface for implementation CRUD
type CoinListRepo interface {
	Create(symbol string, coin *model.CoinItem) error
	Read(symbol string) (*model.CoinItem, error)
	Update(symbol string, coin *model.CoinItem) error
	Delete(symbol string) error
}

// Create a item of the coin list
func (cl *CoinList) Create(symbol string, coin *model.CoinItem) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; ok {
		return fmt.Errorf("%v: %w", symbol, ErrIsExist)
	}
	cl.m[symbol] = coin
	return nil
}

// Read a item of the coin list
func (cl *CoinList) Read(symbol string) (*model.CoinItem, error) {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; !ok {
		return &model.CoinItem{}, fmt.Errorf("%v: %w", symbol, ErrNotFound)
	}
	return cl.m[symbol], nil
}

// Update a item of the coin list by symbol
func (cl *CoinList) Update(symbol string, coin *model.CoinItem) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; !ok {
		return fmt.Errorf("%v: %w", symbol, ErrNotFound)
	}
	cl.m[symbol] = coin
	return nil
}

// Delete a item of the coin list by symbol
func (cl *CoinList) Delete(symbol string) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; !ok {
		return fmt.Errorf("%v: %w", symbol, ErrNotFound)
	}
	delete(cl.m, symbol)
	return nil
}
