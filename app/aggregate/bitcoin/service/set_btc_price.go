package bitcoinservice

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (bs *BitcoinService) SetBTCPrice(newPrice float64) error {
	price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(newPrice), time.Now())
	if err != nil {
		return err
	}

	return bs.bitcoinRepository.SetPrice(price)
}
