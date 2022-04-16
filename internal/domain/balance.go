package domain

type Balance struct {
	BTC BTC
	USD USD
}

func NewBalance(btc BTC, usd USD) Balance {
	return Balance{
		BTC: btc,
		USD: usd,
	}
}

func (b Balance) Total(bitcoinPrice USD) USD {
	return b.BTC.ToUSD(bitcoinPrice) + b.USD
}
