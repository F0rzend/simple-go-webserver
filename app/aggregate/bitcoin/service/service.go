package bitcoinservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BitcoinService struct {
	bitcoinRepository BTCRepository
}

//go:generate moq -out "mock_btc_repository.gen.go" . BTCRepository:MockBTCRepository
type BTCRepository interface {
	GetPrice() bitcoinentity.BTCPrice
	SetPrice(price bitcoinentity.BTCPrice) error
}

func NewBitcoinService(
	bitcoinRepository BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		bitcoinRepository: bitcoinRepository,
	}
}
