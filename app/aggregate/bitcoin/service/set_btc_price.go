package bitcoinservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (bs *BitcoinService) SetBTCPrice(newPrice float64) error {
	price, err := bitcoinentity.NewUSD(newPrice)
	if err != nil {
		return err
	}

	return bs.bitcoinRepository.SetPrice(price)
}
