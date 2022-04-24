package domain

import (
	"fmt"
	"math"
	"time"
)

const SatoshiInBitcoin = 100_000_000

type BTC uint64

func BTCFromFloat(amount float64) BTC {
	return BTC(math.Round(amount * SatoshiInBitcoin))
}

func (btc BTC) ToFloat() float64 {
	return float64(btc) / SatoshiInBitcoin
}

func (btc BTC) ToUSD(bitcoinPrice USD) USD {
	return USD(btc) / bitcoinPrice
}

func (btc BTC) String() string {
	precision := CountPrecision(SatoshiInBitcoin)
	format := fmt.Sprintf("%%.%df BTC", precision)
	return fmt.Sprintf(format, btc.ToFloat())
}

func (btc BTC) Add(other BTC) BTC {
	return btc + other
}

func (btc BTC) Sub(other BTC) BTC {
	return btc - other
}

func (btc BTC) IsNegative() bool {
	return btc < 0
}

type BTCPrice struct {
	Price     USD
	UpdatedAt time.Time
}

func NewBTCPrice(price USD) BTCPrice {
	return BTCPrice{
		Price:     price,
		UpdatedAt: time.Now(),
	}
}
