package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BitcoinService struct {
	SetBTCPrice SetBTCPriceHandler
	GetBTCPrice GetBTCPriceHandler
}

func NewBitcoinService(
	bitcoinRepository entity.BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		SetBTCPrice: MustNewSetBTCPriceCommand(bitcoinRepository),
		GetBTCPrice: MustNewGetBTCCommand(bitcoinRepository),
	}
}
