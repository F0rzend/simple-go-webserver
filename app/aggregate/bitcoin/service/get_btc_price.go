package bitcoinservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (bs *BitcoinService) GetBTCPrice() bitcoinentity.BTCPrice {
	return bs.bitcoinRepository.GetPrice()
}
