package bitcoinservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BitcoinService struct {
	bitcoinRepository bitcoinentity.BTCRepository
}

func NewBitcoinService(
	bitcoinRepository bitcoinentity.BTCRepository,
) *BitcoinService {
	return &BitcoinService{
		bitcoinRepository: bitcoinRepository,
	}
}
