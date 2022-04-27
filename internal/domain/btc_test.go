package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTCFromFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    float64
		expected BTC
		err      error
	}{
		{
			name:     "zero",
			input:    0.0,
			expected: BTC{0},
			err:      nil,
		},
		{
			name:     "satoshi",
			input:    0.00000001,
			expected: BTC{1},
			err:      nil,
		},
		{
			name:     "bitcoin",
			input:    1.0,
			expected: BTC{SatoshiInBitcoin},
			err:      nil,
		},
		{
			name:     "less than 1 satoshi",
			input:    0.000000001,
			expected: BTC{0},
			err:      ErrBTCAmountTooSmall,
		},
		{
			name:     "negative",
			input:    -1.0,
			expected: BTC{0},
			err:      ErrBTCAmountTooSmall,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := BTCFromFloat(tc.input)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBTC_ToFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		btc      BTC
		expected float64
	}{
		{
			name:     "zero",
			btc:      BTC{0},
			expected: 0.0,
		},
		{
			name:     "satoshi",
			btc:      BTC{1},
			expected: 0.00_000_001,
		},
		{
			name:     "bitcoin",
			btc:      BTC{1_23456789},
			expected: 1.23456789,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.btc.ToFloat()

			assert.Equal(t, tc.expected, actual)
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
			input:    BTC{0},
			price:    BTCPrice{price: USD{0}},
			expected: USD{0},
		},
		{
			name:     "btc is 1 usd",
			input:    BTC{SatoshiInBitcoin},
			price:    BTCPrice{price: USD{1_00}},
			expected: USD{1_00},
		},
		{
			name:     "btc is 100 usd",
			input:    BTC{2_00000000},
			price:    BTCPrice{price: USD{10_00}},
			expected: USD{20_00},
		},
		{
			name:     "btc is 1 cent",
			input:    BTC{SatoshiInBitcoin},
			price:    BTCPrice{price: USD{1}},
			expected: USD{1},
		},
		{
			name:     "satoshi to dollar",
			input:    BTC{1},
			price:    BTCPrice{price: USD{SatoshiInBitcoin}},
			expected: USD{1},
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
			input:    BTC{0},
			expected: true,
		},
		{
			name:     "not zero",
			input:    BTC{1},
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
			input:    BTC{0},
			expected: "0 BTC",
		},
		{
			name:     "satoshi",
			input:    BTC{1},
			expected: "0.00000001 BTC",
		},
		{
			name:     "bitcoin",
			input:    BTC{SatoshiInBitcoin},
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

	initial := BTC{1}
	toAdd := BTC{2}
	expected := BTC{3}

	actual, err := initial.Add(toAdd)

	assert.NoError(t, err)
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
			initial:    BTC{2},
			toSubtract: BTC{1},
			expected:   BTC{1},
			err:        nil,
		},
		{
			name:       "subtract more than available",
			initial:    BTC{1},
			toSubtract: BTC{2},
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

func TestNewBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    USD
		expected USD
		err      error
	}{
		{
			name:     "zero",
			input:    USD{0},
			expected: USD{0},
			err:      nil,
		},
		{
			name:     "one cent",
			input:    USD{1},
			expected: USD{},
			err:      ErrBTCPriceTooSmall,
		},
		{
			name:     "one cent for satohi",
			input:    USD{SatoshiInBitcoin},
			expected: USD{SatoshiInBitcoin},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := NewBTCPrice(tc.input)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, actual.GetPrice())
		})
	}
}
