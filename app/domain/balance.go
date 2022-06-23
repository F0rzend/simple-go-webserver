package domain

type Balance struct {
	USD USD
	BTC BTC
}

func NewBalance(usd USD, btc BTC) Balance {
	return Balance{
		USD: usd,
		BTC: btc,
	}
}

func (b Balance) Total(bitcoinPrice BTCPrice) USD {
	return b.BTC.ToUSD(bitcoinPrice).Add(b.USD)
}
