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
			usd:      MustNewUSD(0),
			btc:      MustNewBTC(0),
			btcPrice: BTCPrice{price: MustNewUSD(0)},
			expected: MustNewUSD(0),
		},
		{
			name:     "usd only",
			usd:      MustNewUSD(1),
			btc:      MustNewBTC(0),
			btcPrice: BTCPrice{price: MustNewUSD(1)},
			expected: MustNewUSD(1),
		},
		{
			name:     "btc only",
			usd:      MustNewUSD(0),
			btc:      MustNewBTC(1),
			btcPrice: BTCPrice{price: MustNewUSD(1)},
			expected: MustNewUSD(1),
		},
		{
			name:     "usd and btc",
			usd:      MustNewUSD(1),
			btc:      MustNewBTC(1),
			btcPrice: BTCPrice{price: MustNewUSD(1)},
			expected: MustNewUSD(2),
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
