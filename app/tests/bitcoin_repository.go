package tests

import (
	"time"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func NewMockBitcoinRepository() bitcoinEntity.BTCRepository {
	return &MockBTCRepository{
		GetFunc: func() bitcoinEntity.BTCPrice {
			return bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(100), time.Now())
		},
		SetPriceFunc: func(price bitcoinEntity.USD) error {
			return nil
		},
	}
}
