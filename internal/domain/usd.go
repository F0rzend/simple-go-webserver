package domain

import (
	"fmt"
	"math"
)

const CentInUSDollar = 100

type USD uint64

func USDFromCent(cent uint64) USD {
	return USD(cent)
}

func NewUSD(amount float64) USD {
	return USD(math.Round(amount * CentInUSDollar))
}

func (usd USD) ToFloat() float64 {
	return float64(usd) / CentInUSDollar
}

func (usd USD) String() string {
	precision := CountPrecision(CentInUSDollar)
	format := fmt.Sprintf("$%%.%df", precision)
	return fmt.Sprintf(format, usd.ToFloat())
}
