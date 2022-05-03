package btcrepositories

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

var _ domain.BTCRepository = &MemoryBTCRepository{}

type MemoryBTCRepository struct {
	bitcoin domain.BTCPrice
}

func NewMemoryBTCRepository(initialPrice domain.USD) (*MemoryBTCRepository, error) {
	btcPrice := domain.NewBTCPrice(initialPrice)

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}, nil
}

func (r *MemoryBTCRepository) Get() domain.BTCPrice {
	return r.bitcoin
}

func (r *MemoryBTCRepository) SetPrice(price domain.USD) error {
	btcPrice := domain.NewBTCPrice(price)

	r.bitcoin = btcPrice
	return nil
}
