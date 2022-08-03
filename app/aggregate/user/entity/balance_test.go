package userentity

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
		usd      bitcoinentity.USD
		btc      bitcoinentity.BTC
		btcPrice bitcoinentity.BTCPrice
		expected bitcoinentity.USD
	}{
		{
			name:     "empty",
			usd:      bitcoinentity.MustNewUSD(0),
			btc:      bitcoinentity.MustNewBTC(0),
			btcPrice: bitcoinentity.NewBTCPrice(bitcoinentity.MustNewUSD(0), now),
			expected: bitcoinentity.MustNewUSD(0),
		},
		{
			name:     "usd only",
			usd:      bitcoinentity.MustNewUSD(1),
			btc:      bitcoinentity.MustNewBTC(0),
			btcPrice: bitcoinentity.NewBTCPrice(bitcoinentity.MustNewUSD(1), now),
			expected: bitcoinentity.MustNewUSD(1),
		},
		{
			name:     "btc only",
			usd:      bitcoinentity.MustNewUSD(0),
			btc:      bitcoinentity.MustNewBTC(1),
			btcPrice: bitcoinentity.NewBTCPrice(bitcoinentity.MustNewUSD(1), now),
			expected: bitcoinentity.MustNewUSD(1),
		},
		{
			name:     "usd and btc",
			usd:      bitcoinentity.MustNewUSD(1),
			btc:      bitcoinentity.MustNewBTC(1),
			btcPrice: bitcoinentity.NewBTCPrice(bitcoinentity.MustNewUSD(1), now),
			expected: bitcoinentity.MustNewUSD(2),
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
