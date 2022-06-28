package entity

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type Balance struct {
	USD entity.USD
	BTC entity.BTC
}

func NewBalance(usd entity.USD, btc entity.BTC) Balance {
	return Balance{
		USD: usd,
		BTC: btc,
	}
}

func (b Balance) Total(bitcoinPrice entity.BTCPrice) entity.USD {
	return b.BTC.ToUSD(bitcoinPrice).Add(b.USD)
}
