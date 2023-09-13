package bitcoinentity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSD_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    USD
		expected string
	}{
		{
			name:     "zero",
			input:    NewUSD(0),
			expected: "$0",
		},
		{
			name:     "cent",
			input:    NewUSD(0.01),
			expected: "$0.01",
		},
		{
			name:     "dollar",
			input:    NewUSD(1),
			expected: "$1",
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

func TestUSD_Add(t *testing.T) {
	t.Parallel()

	initial := NewUSD(1)
	toAdd := NewUSD(2)
	expected := NewUSD(3)

	actual := initial.Add(toAdd)

	assert.Equal(t, expected, actual)
}

func TestUSD_Sub(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		initial    USD
		toSubtract USD
		expected   USD
	}{
		{
			name:       "success",
			initial:    NewUSD(2),
			toSubtract: NewUSD(1),
			expected:   NewUSD(1),
		},
		{
			name:       "subtract more than available",
			initial:    NewUSD(1),
			toSubtract: NewUSD(2),
			expected:   NewUSD(-1),
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

func TestUSDComparativeTransactions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		a             USD
		b             USD
		aLessThanB    bool
		aEqualToB     bool
		aGreaterThanB bool
	}{
		{
			name:          "a less than b",
			a:             NewUSD(1),
			b:             NewUSD(2),
			aLessThanB:    true,
			aEqualToB:     false,
			aGreaterThanB: false,
		},
		{
			name:          "a equal to b",
			a:             NewUSD(1),
			b:             NewUSD(1),
			aLessThanB:    false,
			aEqualToB:     true,
			aGreaterThanB: false,
		},
		{
			name:          "a greater than b",
			a:             NewUSD(2),
			b:             NewUSD(1),
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
