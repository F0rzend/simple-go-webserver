package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTC_ToUSD(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		btc      BTC
		btcPrice USD
		expected USD
	}{
		{
			name:     "normal usage",
			btc:      BTC(1),
			btcPrice: USD(100),
			expected: USD(100),
		},
		{
			name:     "empty btc amount",
			btc:      BTC(0),
			btcPrice: USD(100),
			expected: USD(0),
		},
		{
			name:     "empty btc price",
			btc:      BTC(1),
			btcPrice: USD(0),
			expected: USD(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.btc.ToUSD(tc.btcPrice)

			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestBTCFromSatoshi(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		satoshi  uint64
		expected BTC
	}{
		{
			name:     "normal usage",
			satoshi:  100,
			expected: BTC(100),
		},
		{
			name:     "zero satoshi",
			satoshi:  0,
			expected: BTC(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := BTCFromSatoshi(tc.satoshi)

			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestBTC_String(t *testing.T) {
	t.Parallel()

	btc := BTC(100)

	actual := btc.String()
	expected := "0.00000100 BTC"

	assert.Equal(t, expected, actual)
}

func TestBTCFromFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		float    float64
		expected BTC
	}{
		{
			name:     "normal usage",
			float:    100,
			expected: BTC(100),
		},
		{
			name:     "zero float",
			float:    0,
			expected: BTC(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := BTCFromFloat(tc.float)

			assert.Equal(t, tc.expected, actual)
		})
	}

}
