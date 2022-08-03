package bitcoinrepositories

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type MemoryBTCRepository struct {
	bitcoin bitcoinentity.BTCPrice
}

func NewMemoryBTCRepository(initialPrice bitcoinentity.USD) (*MemoryBTCRepository, error) {
	btcPrice := bitcoinentity.NewBTCPrice(initialPrice, time.Now())

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}, nil
}

func (r *MemoryBTCRepository) GetPrice() bitcoinentity.BTCPrice {
	return r.bitcoin
}

func (r *MemoryBTCRepository) SetPrice(price bitcoinentity.USD) error {
	btcPrice := bitcoinentity.NewBTCPrice(price, time.Now())

	r.bitcoin = btcPrice
	return nil
}
