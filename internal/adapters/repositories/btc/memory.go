package btc

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

var (
	_ domain.BTCRepository = &MemoryBTCRepository{}
)

type MemoryBTCRepository struct {
	bitcoin domain.BTCPrice
}

func NewMemoryBTCRepository(initialPrice domain.USD) *MemoryBTCRepository {
	return &MemoryBTCRepository{
		bitcoin: domain.NewBTCPrice(initialPrice),
	}
}

func (r *MemoryBTCRepository) Get() (domain.BTCPrice, error) {
	return r.bitcoin, nil
}

func (r *MemoryBTCRepository) SetPrice(price domain.USD) error {
	r.bitcoin = domain.NewBTCPrice(price)
	return nil
}
