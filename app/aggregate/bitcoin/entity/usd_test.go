package entity

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSD_ToFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		usd      USD
		expected *big.Float
	}{
		{
			name:     "zero",
			usd:      MustNewUSD(0),
			expected: big.NewFloat(0),
		},
		{
			name:     "cent",
			usd:      MustNewUSD(0.01),
			expected: big.NewFloat(0.01),
		},
		{
			name:     "dollar",
			usd:      MustNewUSD(1),
			expected: big.NewFloat(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, _ := tc.usd.ToFloat().Float64()
			expect, _ := tc.expected.Float64()

			assert.Equal(t, expect, actual)
		})
	}
}

func TestUSD_IsZero(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    USD
		expected bool
	}{
		{
			name:     "zero",
			input:    MustNewUSD(0),
			expected: true,
		},
		{
			name:     "not zero",
			input:    MustNewUSD(1),
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

func TestUSD_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    USD
		expected string
	}{
		{
			name:     "zero",
			input:    MustNewUSD(0),
			expected: "$0",
		},
		{
			name:     "cent",
			input:    MustNewUSD(0.01),
			expected: "$0.01",
		},
		{
			name:     "dollar",
			input:    MustNewUSD(1),
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

	initial := MustNewUSD(1)
	toAdd := MustNewUSD(2)
	expected := MustNewUSD(3)

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
		err        error
	}{
		{
			name:       "success",
			initial:    MustNewUSD(2),
			toSubtract: MustNewUSD(1),
			expected:   MustNewUSD(1),
			err:        nil,
		},
		{
			name:       "subtract more than available",
			initial:    MustNewUSD(1),
			toSubtract: MustNewUSD(2),
			expected:   USD{},
			err:        ErrSubtractMoreUSDThanHave,
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
