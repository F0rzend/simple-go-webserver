package domain

import (
	"errors"
	"fmt"
)

const (
	CentInUSD  = 100
	MinimalUSD = 0.01

	USDPrefix = "$"
)

type USD struct {
	cent uint64
}

func USDFromCent(cent uint64) USD {
	return USD{cent}
}

var ErrUSDAmountTooSmall = fmt.Errorf(
	"usd amount cant be less than %f (%f cent)",
	MinimalUSD,
	MinimalUSD*CentInUSD,
)

func USDFromFloat(amount float64) (USD, error) {
	if amount < MinimalUSD && amount != 0 {
		return USD{}, ErrUSDAmountTooSmall
	}

	return USD{uint64(amount * CentInUSD)}, nil
}

func (usd USD) ToFloat() float64 {
	return float64(usd.GetCent()) / CentInUSD
}

func (usd USD) IsZero() bool {
	return usd.cent == 0
}

func (usd USD) String() string {
	if usd.IsZero() {
		return fmt.Sprintf("%s0", USDPrefix)
	}
	if usd.GetCent()%CentInUSD == 0 {
		return fmt.Sprintf("%s%d", USDPrefix, usd.GetCent()/CentInUSD)
	}
	precision := MustCountPrecision(CentInUSD)
	format := fmt.Sprintf("$%%.%df", precision)
	return fmt.Sprintf(format, usd.ToFloat())
}

func (usd USD) Add(other USD) (USD, error) {
	return USD{usd.GetCent() + other.GetCent()}, nil
}

var ErrSubtractMoreUSDThanHave = errors.New("can't subtract more usd than available")

func (usd USD) Sub(toSubtract USD) (USD, error) {
	if toSubtract.GetCent() > usd.GetCent() {
		return USD{}, ErrSubtractMoreUSDThanHave
	}

	return USD{usd.GetCent() - toSubtract.GetCent()}, nil
}

func (usd USD) GetCent() uint64 {
	return usd.cent
}
