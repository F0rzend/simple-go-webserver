package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSDFromCent(t *testing.T) {
	var cent uint64 = 1_00
	expected := USD{cent}
	actual := USDFromCent(cent)

	assert.Equal(t, expected, actual)
}

func TestUSDFromFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    float64
		expected USD
		err      error
	}{
		{
			name:     "zero",
			input:    0.0,
			expected: USD{0},
			err:      nil,
		},
		{
			name:     "cent",
			input:    0.01,
			expected: USD{1},
			err:      nil,
		},
		{
			name:     "dollar",
			input:    1.0,
			expected: USD{1_00},
			err:      nil,
		},
		{
			name:     "less than 1 cent",
			input:    0.001,
			expected: USD{0},
			err:      ErrUSDAmountTooSmall,
		},
		{
			name:     "negative",
			input:    -1.0,
			expected: USD{0},
			err:      ErrUSDAmountTooSmall,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := USDFromFloat(tc.input)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestUSD_ToFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		usd      USD
		expected float64
	}{
		{
			name:     "zero",
			usd:      USD{0},
			expected: 0.0,
		},
		{
			name:     "cent",
			usd:      USD{1},
			expected: 0.01,
		},
		{
			name:     "dollar",
			usd:      USD{1_00},
			expected: 1.0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.usd.ToFloat()

			assert.Equal(t, tc.expected, actual)
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
			input:    USD{0},
			expected: true,
		},
		{
			name:     "not zero",
			input:    USD{1},
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
			input:    USD{0},
			expected: "$0",
		},
		{
			name:     "cent",
			input:    USD{1},
			expected: "$0.01",
		},
		{
			name:     "dollar",
			input:    USD{1_00},
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

	initial := USD{1}
	toAdd := USD{2}
	expected := USD{3}

	actual, err := initial.Add(toAdd)

	assert.NoError(t, err)
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
			initial:    USD{2},
			toSubtract: USD{1},
			expected:   USD{1},
			err:        nil,
		},
		{
			name:       "subtract more than available",
			initial:    USD{1},
			toSubtract: USD{2},
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
