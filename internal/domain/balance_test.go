package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalance_Total(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		balance  Balance
		btcPrice USD
		expected USD
	}{
		{
			name: "normal usage",
			balance: Balance{
				BTC: BTC(1),
				USD: USD(100),
			},
			btcPrice: USD(100),
			expected: USD(200),
		},
		{
			name: "empty btc balance",
			balance: Balance{
				BTC: BTC(0),
				USD: USD(100),
			},
			btcPrice: USD(100),
			expected: USD(100),
		},
		{
			name: "empty btc price",
			balance: Balance{
				BTC: BTC(100),
				USD: USD(100),
			},
			btcPrice: USD(0),
			expected: USD(100),
		},
		{
			name: "empty usd balance",
			balance: Balance{
				BTC: BTC(10),
				USD: USD(0),
			},
			btcPrice: USD(100),
			expected: USD(1_000),
		},
		{
			name: "empty balance",
			balance: Balance{
				BTC: BTC(0),
				USD: USD(0),
			},
			btcPrice: USD(100),
			expected: USD(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.balance.Total(tc.btcPrice)

			assert.Equal(t, tc.expected, actual)
		})
	}

}
