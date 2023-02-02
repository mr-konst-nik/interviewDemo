package service

import (
	"context"
	"fmt"
	"interviewDemo/internal/model"
	"interviewDemo/internal/repo"
)

// CoinService interface for implementation CRUD
type CoinService interface {
	Create(ctx context.Context, symbol string, coin *model.CoinItem) error
	Read(ctx context.Context, symbol string) (*model.CoinItem, error)
	Update(ctx context.Context, symbol string, coin *model.CoinItem) error
	Delete(ctx context.Context, symbol string) (*model.CoinItem, error)
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
func (cl CoinListServise) Create(ctx context.Context, symbol string, coin *model.CoinItem) error {
	err := cl.coinListRepo.Create(ctx, symbol, coin)
	if err != nil {
		err = fmt.Errorf("failed create coin %w", err)
	}
	return err
}

// Read a item of the coin list
func (cl CoinListServise) Read(ctx context.Context, symbol string) (*model.CoinItem, error) {
	m, err := cl.coinListRepo.Read(ctx, symbol)
	if err != nil {
		err = fmt.Errorf("failed read coin %w", err)
	}
	return m, err
}

// Update a item of the coin list by symbol
func (cl CoinListServise) Update(ctx context.Context, symbol string, coin *model.CoinItem) error {
	err := cl.coinListRepo.Update(ctx, symbol, coin)
	if err != nil {
		err = fmt.Errorf("failed update coin %w", err)
	}
	return err
}

// Delete a item of the coin list by symbol
func (cl CoinListServise) Delete(ctx context.Context, symbol string) (*model.CoinItem, error) {
	c, err := cl.coinListRepo.Delete(ctx, symbol)
	if err != nil {
		err = fmt.Errorf("failed delete coin %w", err)
		return nil, err
	}
	return c, nil
}
