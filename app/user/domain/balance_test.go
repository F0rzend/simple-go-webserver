package domain

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"

	"github.com/stretchr/testify/assert"
)

func TestBalance_Total(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		usd      domain.USD
		btc      domain.BTC
		btcPrice domain.BTCPrice
		expected domain.USD
	}{
		{
			name:     "empty",
			usd:      domain.MustNewUSD(0),
			btc:      domain.MustNewBTC(0),
			btcPrice: domain.NewBTCPrice(domain.MustNewUSD(0)),
			expected: domain.MustNewUSD(0),
		},
		{
			name:     "usd only",
			usd:      domain.MustNewUSD(1),
			btc:      domain.MustNewBTC(0),
			btcPrice: domain.NewBTCPrice(domain.MustNewUSD(1)),
			expected: domain.MustNewUSD(1),
		},
		{
			name:     "btc only",
			usd:      domain.MustNewUSD(0),
			btc:      domain.MustNewBTC(1),
			btcPrice: domain.NewBTCPrice(domain.MustNewUSD(1)),
			expected: domain.MustNewUSD(1),
		},
		{
			name:     "usd and btc",
			usd:      domain.MustNewUSD(1),
			btc:      domain.MustNewBTC(1),
			btcPrice: domain.NewBTCPrice(domain.MustNewUSD(1)),
			expected: domain.MustNewUSD(2),
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
