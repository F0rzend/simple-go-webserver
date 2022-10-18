package bitcoinentity

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	CentInUSD = 100
	USDPrefix = "$"
)

type USD struct {
	amount decimal.Decimal
}

func NewUSD(amount float64) USD {
	return USD{decimal.NewFromFloat(amount)}
}

func (usd USD) ToFloat() *big.Float {
	return usd.amount.BigFloat()
}

func (usd USD) String() string {
	if usd.amount.IsZero() {
		return fmt.Sprintf("%s0", USDPrefix)
	}

	if usd.amount.IsInteger() {
		return fmt.Sprintf("%s%d", USDPrefix, usd.amount.BigInt())
	}

	precision := MustCountPrecision(CentInUSD)
	format := fmt.Sprintf("$%%.%df", precision)
	return fmt.Sprintf(format, usd.ToFloat())
}

func (usd USD) Add(toAdd USD) USD {
	return USD{usd.amount.Add(toAdd.amount)}
}

func (usd USD) Sub(toSubtract USD) USD {
	return USD{usd.amount.Sub(toSubtract.amount)}
}

func (usd USD) LessThan(toCompare USD) bool {
	return usd.amount.LessThan(toCompare.amount)
}

func (usd USD) Equal(toCompare USD) bool {
	return usd.amount.Equal(toCompare.amount)
}
