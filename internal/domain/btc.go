package domain

import (
	"errors"
	"fmt"
	"time"
)

const (
	SatoshiInBitcoin = 100_000_000
	MinimalBitcoin   = 0.00_000_001

	BTCSuffix = "BTC"
)

type BTC struct {
	satoshi uint64
}

var ErrBTCAmountTooSmall = fmt.Errorf(
	"bitcoin amount cant be less than %f (%f satoshi)",
	MinimalBitcoin,
	MinimalBitcoin*SatoshiInBitcoin,
)

func BTCFromFloat(amount float64) (BTC, error) {
	if amount < MinimalBitcoin && amount != 0 {
		return BTC{}, ErrBTCAmountTooSmall
	}
	return BTC{uint64(amount * SatoshiInBitcoin)}, nil
}

func (btc BTC) ToFloat() float64 {
	return float64(btc.GetSatoshi()) / SatoshiInBitcoin
}

func (btc BTC) ToUSD(price BTCPrice) USD {
	if price.GetPrice().GetCent() == 0 {
		return USD{0}
	}
	return USD{uint64(btc.ToFloat() * float64(price.GetPrice().GetCent()))}
}

func (btc BTC) IsZero() bool {
	return btc.satoshi == 0
}

func (btc BTC) String() string {
	if btc.IsZero() {
		return fmt.Sprintf("0 %s", BTCSuffix)
	}
	if btc.GetSatoshi()%SatoshiInBitcoin == 0 {
		return fmt.Sprintf("%d %s", btc.GetSatoshi()/SatoshiInBitcoin, BTCSuffix)
	}

	precision := MustCountPrecision(SatoshiInBitcoin)
	format := fmt.Sprintf("%%.%df %s", precision, BTCSuffix)
	return fmt.Sprintf(format, btc.ToFloat())
}

func (btc BTC) Add(other BTC) (BTC, error) {
	return BTC{btc.GetSatoshi() + other.GetSatoshi()}, nil
}

var ErrSubtractMoreBTCThanHave = errors.New("can't subtract more btc than available")

func (btc BTC) Sub(toSubtract BTC) (BTC, error) {
	if toSubtract.GetSatoshi() > btc.GetSatoshi() {
		return BTC{}, ErrSubtractMoreBTCThanHave
	}
	return BTC{btc.GetSatoshi() - toSubtract.GetSatoshi()}, nil
}

func (btc BTC) GetSatoshi() uint64 {
	return btc.satoshi
}

type BTCPrice struct {
	price     USD
	updatedAt time.Time
}

var ErrBTCPriceTooSmall = errors.New("bitcoin price cant be less than 1 satoshi by 1 cent")

func NewBTCPrice(price USD) (BTCPrice, error) {
	if price.GetCent() < SatoshiInBitcoin && price.GetCent() != 0 {
		return BTCPrice{}, ErrBTCPriceTooSmall
	}
	return BTCPrice{
		price:     price,
		updatedAt: time.Now(),
	}, nil
}

func (btcPrice BTCPrice) GetPrice() USD {
	return btcPrice.price
}

func (btcPrice BTCPrice) GetUpdatedAt() time.Time {
	return btcPrice.updatedAt
}
