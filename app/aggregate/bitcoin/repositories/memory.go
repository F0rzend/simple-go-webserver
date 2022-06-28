package repositories

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type MemoryBTCRepository struct {
	bitcoin entity.BTCPrice
}

func NewMemoryBTCRepository(initialPrice entity.USD) (*MemoryBTCRepository, error) {
	btcPrice := entity.NewBTCPrice(initialPrice, time.Now())

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}, nil
}

func (r *MemoryBTCRepository) Get() entity.BTCPrice {
	return r.bitcoin
}

func (r *MemoryBTCRepository) SetPrice(price entity.USD) error {
	btcPrice := entity.NewBTCPrice(price, time.Now())

	r.bitcoin = btcPrice
	return nil
}
