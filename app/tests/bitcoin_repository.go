package tests

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func NewMockBitcoinRepository() entity.BTCRepository {
	return &MockBTCRepository{
		GetFunc: func() entity.BTCPrice {
			return entity.NewBTCPrice(entity.MustNewUSD(100), time.Now())
		},
		SetPriceFunc: func(price entity.USD) error {
			return nil
		},
	}
}
