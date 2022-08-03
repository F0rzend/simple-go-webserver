package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BitcoinService struct {
	bitcoinRepository entity.BTCRepository
}

func NewBitcoinService(
	bitcoinRepository entity.BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		bitcoinRepository: bitcoinRepository,
	}
}
