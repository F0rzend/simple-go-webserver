package domain

import (
	"fmt"
	"math"
)

const SatoshiInBitcoin = 100_000_000

type BTC uint64

func BTCFromSatoshi(satoshi uint64) BTC {
	return BTC(satoshi)
}

func BTCFromFloat(amount float64) BTC {
	return BTC(math.Round(amount * SatoshiInBitcoin))
}

func (btc BTC) ToFloat() float64 {
	return float64(btc) / SatoshiInBitcoin
}

func (btc BTC) ToUSD(bitcoinPrice USD) USD {
	return USD(btc) * bitcoinPrice
}

func (btc BTC) String() string {
	precision := CountPrecision(SatoshiInBitcoin)
	format := fmt.Sprintf("%%.%df BTC", precision)
	return fmt.Sprintf(format, btc.ToFloat())
}
