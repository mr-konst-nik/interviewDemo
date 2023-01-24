package service

import (
	"fmt"
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
	err := cl.coinListRepo.Create(symbol, coin)
	if err != nil {
		err = fmt.Errorf("failed create coin %w", err)
	}
	return err
}

// Read a item of the coin list
func (cl CoinListServise) Read(symbol string) (*model.CoinItem, error) {
	m, err := cl.coinListRepo.Read(symbol)
	if err != nil {
		err = fmt.Errorf("failed read coin %w", err)
	}
	return m, err
}

// Update a item of the coin list by symbol
func (cl CoinListServise) Update(symbol string, coin *model.CoinItem) error {
	err := cl.coinListRepo.Update(symbol, coin)
	if err != nil {
		err = fmt.Errorf("failed update coin %w", err)
	}
	return err
}

// Delete a item of the coin list by symbol
func (cl CoinListServise) Delete(symbol string) (*model.CoinItem, error) {
	c, err := cl.coinListRepo.Delete(symbol)
	if err != nil {
		err = fmt.Errorf("failed delete coin %w", err)
		return nil, err
	}
	return c, nil
}
