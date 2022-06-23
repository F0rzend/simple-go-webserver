package domain

import (
	"errors"
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

func NewUSD(amount float64) (USD, error) {
	if amount < 0 {
		return USD{}, ErrNegativeCurrency(amount)
	}
	return USD{decimal.NewFromFloat(amount)}, nil
}

func MustNewUSD(amount float64) USD {
	usd, err := NewUSD(amount)
	if err != nil {
		panic(err)
	}
	return usd
}

func (usd USD) ToFloat() *big.Float {
	return usd.amount.BigFloat()
}

func (usd USD) ToFloat64() float64 {
	amount, _ := usd.ToFloat().Float64()
	return amount
}

func (usd USD) IsZero() bool {
	return usd.amount.IsZero()
}

func (usd USD) String() string {
	if usd.IsZero() {
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

var ErrSubtractMoreUSDThanHave = errors.New("can't subtract more usd than available")

func (usd USD) Sub(toSubtract USD) (USD, error) {
	if toSubtract.amount.GreaterThan(usd.amount) {
		return USD{}, ErrSubtractMoreUSDThanHave
	}

	return USD{usd.amount.Sub(toSubtract.amount)}, nil
}

func (usd USD) LessThan(toCompare USD) bool {
	return usd.amount.LessThan(toCompare.amount)
}

func (usd USD) Equal(toCompare USD) bool {
	return usd.amount.Equal(toCompare.amount)
}
