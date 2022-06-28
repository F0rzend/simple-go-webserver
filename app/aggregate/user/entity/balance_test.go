package entity

import (
	"testing"

	domain2 "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	"github.com/stretchr/testify/assert"
)

func TestBalance_Total(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		usd      domain2.USD
		btc      domain2.BTC
		btcPrice domain2.BTCPrice
		expected domain2.USD
	}{
		{
			name:     "empty",
			usd:      domain2.MustNewUSD(0),
			btc:      domain2.MustNewBTC(0),
			btcPrice: domain2.NewBTCPrice(domain2.MustNewUSD(0)),
			expected: domain2.MustNewUSD(0),
		},
		{
			name:     "usd only",
			usd:      domain2.MustNewUSD(1),
			btc:      domain2.MustNewBTC(0),
			btcPrice: domain2.NewBTCPrice(domain2.MustNewUSD(1)),
			expected: domain2.MustNewUSD(1),
		},
		{
			name:     "btc only",
			usd:      domain2.MustNewUSD(0),
			btc:      domain2.MustNewBTC(1),
			btcPrice: domain2.NewBTCPrice(domain2.MustNewUSD(1)),
			expected: domain2.MustNewUSD(1),
		},
		{
			name:     "usd and btc",
			usd:      domain2.MustNewUSD(1),
			btc:      domain2.MustNewBTC(1),
			btcPrice: domain2.NewBTCPrice(domain2.MustNewUSD(1)),
			expected: domain2.MustNewUSD(2),
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
