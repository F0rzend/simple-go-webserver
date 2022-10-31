package bitcoinrepositories

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type MemoryBTCRepository struct {
	bitcoin bitcoinentity.BTCPrice
}

func NewMemoryBTCRepository() *MemoryBTCRepository {
	const defaultPrice = 1e-2

	btcPrice, _ := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(defaultPrice), time.Now())

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}
}

func (r *MemoryBTCRepository) GetPrice() bitcoinentity.BTCPrice {
	return r.bitcoin
}

func (r *MemoryBTCRepository) SetPrice(price bitcoinentity.BTCPrice) error {
	r.bitcoin = price

	return nil
}
