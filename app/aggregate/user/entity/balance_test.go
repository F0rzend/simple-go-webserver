package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestBalance_Total(t *testing.T) {
	t.Parallel()

	now := time.Now()

	testCases := []struct {
		name     string
		usd      entity.USD
		btc      entity.BTC
		btcPrice entity.BTCPrice
		expected entity.USD
	}{
		{
			name:     "empty",
			usd:      entity.MustNewUSD(0),
			btc:      entity.MustNewBTC(0),
			btcPrice: entity.NewBTCPrice(entity.MustNewUSD(0), now),
			expected: entity.MustNewUSD(0),
		},
		{
			name:     "usd only",
			usd:      entity.MustNewUSD(1),
			btc:      entity.MustNewBTC(0),
			btcPrice: entity.NewBTCPrice(entity.MustNewUSD(1), now),
			expected: entity.MustNewUSD(1),
		},
		{
			name:     "btc only",
			usd:      entity.MustNewUSD(0),
			btc:      entity.MustNewBTC(1),
			btcPrice: entity.NewBTCPrice(entity.MustNewUSD(1), now),
			expected: entity.MustNewUSD(1),
		},
		{
			name:     "usd and btc",
			usd:      entity.MustNewUSD(1),
			btc:      entity.MustNewBTC(1),
			btcPrice: entity.NewBTCPrice(entity.MustNewUSD(1), now),
			expected: entity.MustNewUSD(2),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			balance := NewBalance(tc.usd, tc.btc)

			actual := balance.Total(tc.btcPrice)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
