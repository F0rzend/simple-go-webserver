package domain

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/shopspring/decimal"
)

const (
	SatoshiInBitcoin = 100_000_000

	BTCSuffix = "BTC"
)

type BTC struct {
	amount decimal.Decimal
}

func NewBTC(amount float64) (BTC, error) {
	if amount < 0 {
		return BTC{}, ErrNegativeCurrency(amount)
	}
	return BTC{decimal.NewFromFloat(amount)}, nil
}

func MustNewBTC(amount float64) BTC {
	btc, err := NewBTC(amount)
	if err != nil {
		panic(err)
	}
	return btc
}

func (btc BTC) ToFloat() *big.Float {
	return btc.amount.BigFloat()
}

func (btc BTC) ToFloat64() float64 {
	amount, _ := btc.amount.Float64()
	return amount
}

func (btc BTC) ToUSD(price BTCPrice) USD {
	return USD{btc.amount.Mul(price.GetPrice().amount)}
}

func (btc BTC) IsZero() bool {
	return btc.amount.IsZero()
}

func (btc BTC) String() string {
	if btc.IsZero() {
		return fmt.Sprintf("0 %s", BTCSuffix)
	}
	if btc.amount.IsInteger() {
		return fmt.Sprintf("%d %s", btc.amount.BigInt(), BTCSuffix)
	}

	precision := MustCountPrecision(SatoshiInBitcoin)
	format := fmt.Sprintf("%%.%df %s", precision, BTCSuffix)
	return fmt.Sprintf(format, btc.ToFloat())
}

func (btc BTC) Add(toAdd BTC) BTC {
	return BTC{btc.amount.Add(toAdd.amount)}
}

var ErrSubtractMoreBTCThanHave = errors.New("can't subtract more btc than available")

func (btc BTC) Sub(toSubtract BTC) (BTC, error) {
	if toSubtract.amount.GreaterThan(btc.amount) {
		return BTC{}, ErrSubtractMoreBTCThanHave
	}
	return BTC{btc.amount.Sub(toSubtract.amount)}, nil
}

func (btc BTC) LessThan(other BTC) bool {
	return btc.amount.LessThan(other.amount)
}

func (btc BTC) Equal(other BTC) bool {
	return btc.amount.Equal(other.amount)
}

type BTCPrice struct {
	price     USD
	updatedAt time.Time
}

func NewBTCPrice(price USD) BTCPrice {
	return BTCPrice{
		price:     price,
		updatedAt: time.Now(),
	}
}

func (btcPrice BTCPrice) GetPrice() USD {
	return btcPrice.price
}

func (btcPrice BTCPrice) GetUpdatedAt() time.Time {
	return btcPrice.updatedAt
}
