package bitcoinservice

import (
	"fmt"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (bs *BitcoinService) SetBTCPrice(newPrice float64) error {
	price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(newPrice), time.Now())
	if err != nil {
		return fmt.Errorf("cannot create new btc price: %w", err)
	}

	err = bs.bitcoinRepository.SetPrice(price)
	if err != nil {
		return fmt.Errorf("cannot set btc price: %w", err)
	}

	return nil
}
