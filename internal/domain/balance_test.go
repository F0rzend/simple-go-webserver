package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance_Total(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		usd      USD
		btc      BTC
		btcPrice BTCPrice
		expected USD
	}{
		{
			name:     "empty",
			usd:      USD{0},
			btc:      BTC{0},
			btcPrice: BTCPrice{price: USD{0}},
			expected: USD{0},
		},
		{
			name:     "usd only",
			usd:      USD{1_00},
			btc:      BTC{0},
			btcPrice: BTCPrice{price: USD{0}},
			expected: USD{1_00},
		},
		{
			name:     "btc only",
			usd:      USD{0},
			btc:      BTC{SatoshiInBitcoin},
			btcPrice: BTCPrice{price: USD{1_00}},
			expected: USD{1_00},
		},
		{
			name:     "usd and btc",
			usd:      USD{1_00},
			btc:      BTC{SatoshiInBitcoin},
			btcPrice: BTCPrice{price: USD{1_00}},
			expected: USD{2_00},
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
