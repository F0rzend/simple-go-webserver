package bitcoinrepositories

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

const minimalPrice = 1e-2

type MemoryBTCRepository struct {
	bitcoin bitcoinentity.BTCPrice
}

func NewMemoryBTCRepository() *MemoryBTCRepository {
	btcPrice := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(minimalPrice), time.Now())

	return &MemoryBTCRepository{
		bitcoin: btcPrice,
	}
}

func (r *MemoryBTCRepository) GetPrice() bitcoinentity.BTCPrice {
	return r.bitcoin
}

func (r *MemoryBTCRepository) SetPrice(price bitcoinentity.USD) error {
	r.bitcoin = bitcoinentity.NewBTCPrice(price, time.Now())

	return nil
}
