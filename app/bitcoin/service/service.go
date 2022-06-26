package service

import (
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/repositories"
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/service/commands"
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/service/queries"
)

type BitcoinService struct {
	Commands commands.Commands
	Queries  queries.Queries
}

func newBitcoinService(
	bitcoinRepository domain.BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		Commands: commands.Commands{
			SetBTCPrice: commands.MustNewSetBTCPriceCommand(bitcoinRepository),
		},
		Queries: queries.Queries{
			GetBTC: queries.MustNewGetBTCCommand(bitcoinRepository),
		},
	}
}

func MustBitcoinService() *BitcoinService {
	btcRepository, err := repositories.NewMemoryBTCRepository(domain.MustNewUSD(100))
	if err != nil {
		panic(err)
	}

	return newBitcoinService(btcRepository)
}

func NewComponentTestBitcoinService() (*BitcoinService, error) {
	bitcoinRepository := &repositories.MockBTCRepository{
		GetFunc: func() domain.BTCPrice {
			return domain.NewBTCPrice(domain.MustNewUSD(100))
		},
		SetPriceFunc: func(price domain.USD) error {
			return nil
		},
	}

	return newBitcoinService(bitcoinRepository), nil
}
