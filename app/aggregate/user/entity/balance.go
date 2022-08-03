package userentity

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type Balance struct {
	USD bitcoinentity.USD
	BTC bitcoinentity.BTC
}

func NewBalance(usd bitcoinentity.USD, btc bitcoinentity.BTC) Balance {
	return Balance{
		USD: usd,
		BTC: btc,
	}
}

func (b Balance) Total(bitcoinPrice bitcoinentity.BTCPrice) bitcoinentity.USD {
	return b.BTC.ToUSD(bitcoinPrice).Add(b.USD)
}
