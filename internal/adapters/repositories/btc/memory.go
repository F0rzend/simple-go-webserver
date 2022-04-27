package btc

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

var _ domain.BTCRepository = &MemoryBTCRepository{}

type MemoryBTCRepository struct {
	bitcoin domain.BTCPrice
}

func NewMemoryBTCRepository(initialPrice domain.USD) (*MemoryBTCRepository, error) {
	btcPrice, err := domain.NewBTCPrice(initialPrice)
	if err != nil {
		return nil, err
	}

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}, nil
}

func (r *MemoryBTCRepository) Get() (domain.BTCPrice, error) {
	return r.bitcoin, nil
}

func (r *MemoryBTCRepository) SetPrice(price domain.USD) error {
	btcPrice, err := domain.NewBTCPrice(price)
	if err != nil {
		return err
	}

	r.bitcoin = btcPrice
	return nil
}
