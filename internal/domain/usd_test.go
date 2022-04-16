package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUSDFromCent(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		cent     uint64
		expected USD
	}{
		{
			name:     "normal usage",
			cent:     100,
			expected: USD(100),
		},
		{
			name:     "zero cent",
			cent:     0,
			expected: USD(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := USDFromCent(tc.cent)

			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestUSD_String(t *testing.T) {
	t.Parallel()

	usd := USD(100)

	actual := usd.String()
	expected := "$1.00"

	assert.Equal(t, expected, actual)
}

func TestUSDFromFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		float    float64
		expected USD
	}{
		{
			name:     "normal usage",
			float:    100,
			expected: USD(100),
		},
		{
			name:     "zero float",
			float:    0,
			expected: USD(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := USDFromFloat(tc.float)

			assert.Equal(t, tc.expected, actual)
		})
	}

}
