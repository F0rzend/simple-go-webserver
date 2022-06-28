package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
)

type BitcoinService struct {
	SetBTCPrice SetBTCPriceHandler
	GetBTCPrice GetBTCPriceHandler
}

func newBitcoinService(
	bitcoinRepository entity.BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		SetBTCPrice: MustNewSetBTCPriceCommand(bitcoinRepository),
		GetBTCPrice: MustNewGetBTCCommand(bitcoinRepository),
	}
}

func MustBitcoinService() *BitcoinService {
	btcRepository, err := repositories.NewMemoryBTCRepository(entity.MustNewUSD(100))
	if err != nil {
		panic(err)
	}

	return newBitcoinService(btcRepository)
}

func NewComponentTestBitcoinService() (*BitcoinService, error) {
	bitcoinRepository := &repositories.MockBTCRepository{
		GetFunc: func() entity.BTCPrice {
			return entity.NewBTCPrice(entity.MustNewUSD(100))
		},
		SetPriceFunc: func(price entity.USD) error {
			return nil
		},
	}

	return newBitcoinService(bitcoinRepository), nil
}
