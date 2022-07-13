package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

//go:generate moq -out "./mock.gen.go" . BitcoinService:MockBitcoinService
type BitcoinService interface {
	GetBTCPrice() entity.BTCPrice
	SetBTCPrice(newPrice float64) error
}

type BitcoinServiceImpl struct {
	bitcoinRepository entity.BTCRepository
}

func NewBitcoinService(
	bitcoinRepository entity.BTCRepository,
) BitcoinService {
	return &BitcoinServiceImpl{
		bitcoinRepository: bitcoinRepository,
	}
}
