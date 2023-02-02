package repo

import (
	"context"
	"fmt"
	"interviewDemo/internal/model"
	"sync"
)

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
	Create(ctx context.Context, symbol string, coin *model.CoinItem) error
	Read(ctx context.Context, symbol string) (*model.CoinItem, error)
	Update(ctx context.Context, symbol string, coin *model.CoinItem) error
	Delete(ctx context.Context, symbol string) (*model.CoinItem, error)
}

// Create a item of the coin list
func (cl *CoinList) Create(ctx context.Context, symbol string, coin *model.CoinItem) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; ok {
		return fmt.Errorf("%v: %w", symbol, model.ErrIsExist)
	}
	cl.m[symbol] = coin
	return nil
}

// Read a item of the coin list
func (cl *CoinList) Read(ctx context.Context, symbol string) (*model.CoinItem, error) {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; !ok {
		return &model.CoinItem{}, fmt.Errorf("%v: %w", symbol, model.ErrNotFound)
	}
	return cl.m[symbol], nil
}

// Update a item of the coin list by symbol
func (cl *CoinList) Update(ctx context.Context, symbol string, coin *model.CoinItem) error {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.m[symbol]; !ok {
		return fmt.Errorf("%v: %w", symbol, model.ErrNotFound)
	}
	cl.m[symbol] = coin
	return nil
}

// Delete a item of the coin list by symbol
func (cl *CoinList) Delete(ctx context.Context, symbol string) (*model.CoinItem, error) {
	cl.Lock()
	defer cl.Unlock()
	c, ok := cl.m[symbol]
	if !ok {
		return nil, fmt.Errorf("%v: %w", symbol, model.ErrNotFound)
	}
	delete(cl.m, symbol)
	return c, nil
}
