package service

import (
	"interviewDemo/internal/model"
	"interviewDemo/internal/repo"
)

// CoinService interface for implementation CRUD
type CoinService interface {
	Create(symbol string, coin *model.CoinItem) error
	Read(symbol string) (*model.CoinItem, error)
	Update(symbol string, coin *model.CoinItem) error
	Delete(symbol string) error
}

// CoinListServise type of service
type CoinListServise struct {
	coinListRepo repo.CoinListRepo
}

// NewCoinListServise init a new service
func NewCoinListServise(coinListRepo repo.CoinListRepo) *CoinListServise {
	return &CoinListServise{
		coinListRepo: coinListRepo,
	}
}

// Create a item of the coin list
func (cl CoinListServise) Create(symbol string, coin *model.CoinItem) error {
	return cl.coinListRepo.Create(symbol, coin)
}

// Read a item of the coin list
func (cl CoinListServise) Read(symbol string) (*model.CoinItem, error) {
	return cl.coinListRepo.Read(symbol)
}

// Update a item of the coin list by symbol
func (cl CoinListServise) Update(symbol string, coin *model.CoinItem) error {
	return cl.coinListRepo.Update(symbol, coin)
}

// Delete a item of the coin list by symbol
func (cl CoinListServise) Delete(symbol string) error {
	return cl.coinListRepo.Delete(symbol)
}
