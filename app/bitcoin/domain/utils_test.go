package domain

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrecision(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    int
		expected int
		err      error
	}{
		{
			name:     "bitcoin",
			input:    100_000_000,
			expected: 8,
		},
		{
			name:     "usd",
			input:    100,
			expected: 2,
		},
		{
			name:  "negative",
			input: -100,
			err:   ErrPrecisionInvalidFormat,
		},
		{
			name:  "number one",
			input: 1,
			err:   ErrPrecisionInvalidFormat,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := CountPrecision(tc.input)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkCountPrecision(b *testing.B) {
	checkValue := 100_000_000
	b.Run("string parsing", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len(strconv.Itoa(checkValue))
		}
	})

	b.Run("log10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			math.Log10(float64(checkValue))
		}
	})
}
