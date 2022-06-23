package domain

import (
	"errors"
	"math"
)

const MinimalPrecision = 10

var ErrPrecisionInvalidFormat = errors.New("invalid format to get precision")

// CountPrecision counts float numbers precision
// From integer in format like 1 (one) with some amount of zeroes
// It uses math.Log10 to count the number of digits
// It is used for formatting (satoshi -> btc, cent -> dollar)
// See BTC.String and USD.String for example of usage
func CountPrecision(value int) (int, error) {
	if value < MinimalPrecision {
		return 0, ErrPrecisionInvalidFormat
	}

	precision := math.Log10(float64(value))

	if int(math.Pow(10, precision)) != value { // nolint: gomnd
		return 0, ErrPrecisionInvalidFormat
	}
	return int(precision), nil
}

func MustCountPrecision(value int) int {
	precision, err := CountPrecision(value)
	if err != nil {
		panic(err)
	}
	return precision
}
