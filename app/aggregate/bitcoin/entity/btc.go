package bitcoinentity

import (
	"fmt"
	"math/big"
	"time"

	"github.com/shopspring/decimal"
)

const (
	SatoshiInBitcoin = 100_000_000
	BTCSuffix        = "BTC"
)

type BTC struct {
	amount decimal.Decimal
}

func NewBTC(amount float64) BTC {
	return BTC{decimal.NewFromFloat(amount)}
}

func (btc BTC) ToFloat() *big.Float {
	return btc.amount.BigFloat()
}

func (btc BTC) ToUSD(price BTCPrice) USD {
	return USD{btc.amount.Mul(price.GetPrice().amount)}
}

func (btc BTC) String() string {
	if btc.amount.IsZero() {
		return fmt.Sprintf("0 %s", BTCSuffix)
	}
	if btc.amount.IsInteger() {
		return fmt.Sprintf("%d %s", btc.amount.BigInt(), BTCSuffix)
	}

	precision := MustCountPrecision(SatoshiInBitcoin)
	format := fmt.Sprintf("%%.%df %s", precision, BTCSuffix)
	return fmt.Sprintf(format, btc.amount.BigFloat())
}

func (btc BTC) Add(toAdd BTC) BTC {
	return BTC{btc.amount.Add(toAdd.amount)}
}

func (btc BTC) Sub(toSubtract BTC) BTC {
	return BTC{btc.amount.Sub(toSubtract.amount)}
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

func NewBTCPrice(price USD, updatedAt time.Time) BTCPrice {
	return BTCPrice{
		price:     price,
		updatedAt: updatedAt,
	}
}

func (btcPrice BTCPrice) GetPrice() USD {
	return btcPrice.price
}

func (btcPrice BTCPrice) GetUpdatedAt() time.Time {
	return btcPrice.updatedAt
}
