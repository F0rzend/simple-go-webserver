package bitcoinentity

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTC_ToFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		btc      BTC
		expected *big.Float
	}{
		{
			name:     "zero",
			btc:      MustNewBTC(0),
			expected: big.NewFloat(0),
		},
		{
			name:     "satoshi",
			btc:      MustNewBTC(1e-8),
			expected: big.NewFloat(1e-8),
		},
		{
			name:     "bitcoin",
			btc:      MustNewBTC(1),
			expected: big.NewFloat(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, _ := tc.btc.ToFloat().Float64()
			expect, _ := tc.expected.Float64()

			assert.Equal(t, expect, actual)
		})
	}
}

func TestBTC_ToUSD(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    BTC
		price    BTCPrice
		expected USD
	}{
		{
			name:     "zero",
			input:    MustNewBTC(0),
			price:    BTCPrice{price: MustNewUSD(0)},
			expected: MustNewUSD(0),
		},
		{
			name:     "btc is 1 usd",
			input:    MustNewBTC(1),
			price:    BTCPrice{price: MustNewUSD(1)},
			expected: MustNewUSD(1),
		},
		{
			name:     "btc is 100 usd",
			input:    MustNewBTC(1),
			price:    BTCPrice{price: MustNewUSD(10)},
			expected: MustNewUSD(10),
		},
		{
			name:     "btc is 1 cent",
			input:    MustNewBTC(100),
			price:    BTCPrice{price: MustNewUSD(0.01)},
			expected: MustNewUSD(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.input.ToUSD(tc.price)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBTC_IsZero(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    BTC
		expected bool
	}{
		{
			name:     "zero",
			input:    MustNewBTC(0),
			expected: true,
		},
		{
			name:     "not zero",
			input:    MustNewBTC(1),
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.input.IsZero()

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBTC_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    BTC
		expected string
	}{
		{
			name:     "zero",
			input:    MustNewBTC(0),
			expected: "0 BTC",
		},
		{
			name:     "satoshi",
			input:    MustNewBTC(1e-8),
			expected: "0.00000001 BTC",
		},
		{
			name:     "bitcoin",
			input:    MustNewBTC(1),
			expected: "1 BTC",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.input.String()

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBTC_Add(t *testing.T) {
	t.Parallel()

	initial := MustNewBTC(1)
	toAdd := MustNewBTC(2)
	expected := MustNewBTC(3)

	actual := initial.Add(toAdd)

	assert.Equal(t, expected, actual)
}

func TestBTC_Sub(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		initial    BTC
		toSubtract BTC
		expected   BTC
		err        error
	}{
		{
			name:       "success",
			initial:    MustNewBTC(2),
			toSubtract: MustNewBTC(1),
			expected:   MustNewBTC(1),
			err:        nil,
		},
		{
			name:       "subtract more than available",
			initial:    MustNewBTC(1),
			toSubtract: MustNewBTC(2),
			expected:   BTC{},
			err:        ErrSubtractMoreBTCThanHave,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.initial.Sub(tc.toSubtract)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
