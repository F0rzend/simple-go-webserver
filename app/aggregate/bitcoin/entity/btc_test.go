package bitcoinentity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
			input:    NewBTC(0),
			price:    BTCPrice{price: NewUSD(0)},
			expected: NewUSD(0),
		},
		{
			name:     "btc is 1 usd",
			input:    NewBTC(1),
			price:    BTCPrice{price: NewUSD(1)},
			expected: NewUSD(1),
		},
		{
			name:     "btc is 100 usd",
			input:    NewBTC(1),
			price:    BTCPrice{price: NewUSD(10)},
			expected: NewUSD(10),
		},
		{
			name:     "btc is 1 cent",
			input:    NewBTC(100),
			price:    BTCPrice{price: NewUSD(0.01)},
			expected: NewUSD(1),
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

func TestBTC_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    BTC
		expected string
	}{
		{
			name:     "zero",
			input:    NewBTC(0),
			expected: "0 BTC",
		},
		{
			name:     "satoshi",
			input:    NewBTC(1e-8),
			expected: "0.00000001 BTC",
		},
		{
			name:     "bitcoin",
			input:    NewBTC(1),
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

	initial := NewBTC(1)
	toAdd := NewBTC(2)
	expected := NewBTC(3)

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
	}{
		{
			name:       "success",
			initial:    NewBTC(2),
			toSubtract: NewBTC(1),
			expected:   NewBTC(1),
		},
		{
			name:       "subtract more than available",
			initial:    NewBTC(1),
			toSubtract: NewBTC(2),
			expected:   NewBTC(-1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.initial.Sub(tc.toSubtract)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBTCComparativeTransactions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		a             BTC
		b             BTC
		aLessThanB    bool
		aEqualToB     bool
		aGreaterThanB bool
	}{
		{
			name:          "a less than b",
			a:             NewBTC(1),
			b:             NewBTC(2),
			aLessThanB:    true,
			aEqualToB:     false,
			aGreaterThanB: false,
		},
		{
			name:          "a equal to b",
			a:             NewBTC(1),
			b:             NewBTC(1),
			aLessThanB:    false,
			aEqualToB:     true,
			aGreaterThanB: false,
		},
		{
			name:          "a greater than b",
			a:             NewBTC(2),
			b:             NewBTC(1),
			aLessThanB:    false,
			aEqualToB:     false,
			aGreaterThanB: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.aLessThanB, tc.a.LessThan(tc.b))
			assert.Equal(t, tc.aEqualToB, tc.a.Equal(tc.b))
			assert.Equal(t, tc.aGreaterThanB, tc.b.LessThan(tc.a))
		})
	}
}

func TestNewBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		price         USD
		expectedPrice USD
		err           error
	}{
		{
			name:          "success",
			price:         NewUSD(100.0),
			expectedPrice: NewUSD(100),
			err:           nil,
		},
		{
			name:          "negative price",
			price:         NewUSD(-1.0),
			expectedPrice: USD{},
			err:           ErrNegativePrice,
		},
	}

	now := time.Now()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			price, err := NewBTCPrice(tc.price, now)

			assert.ErrorIs(t, tc.err, err)
			assert.Equal(t, tc.expectedPrice, price.GetPrice())
		})
	}
}
