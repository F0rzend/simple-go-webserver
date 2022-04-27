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
	return USD{b.BTC.ToUSD(bitcoinPrice).GetCent() + b.USD.GetCent()}
}
