package domain

import "math"

// CountPrecision from integer like 1 (one) with some amount of zeroes
// It uses math.Log10 to count the number of digits
// It is used for formatting (satoshi -> bitcoin, cent -> dollar)
func CountPrecision(value int) int {
	return int(math.Log10(float64(value)))
}
