package domain

import (
	btc "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
)

type Balance struct {
	USD btc.USD
	BTC btc.BTC
}

func NewBalance(usd btc.USD, btc btc.BTC) Balance {
	return Balance{
		USD: usd,
		BTC: btc,
	}
}

func (b Balance) Total(bitcoinPrice btc.BTCPrice) btc.USD {
	return b.BTC.ToUSD(bitcoinPrice).Add(b.USD)
}
