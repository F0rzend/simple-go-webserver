package domain

import (
	"fmt"
	"math"
	"testing"
)

func ExampleCountPrecision() {
	precision := CountPrecision(100_000_000)
	fmt.Println(precision)
	fmt.Printf(fmt.Sprintf("%%.%df", precision), 3.1415)
	// Output:
	// 8
	// 3.14150000
}

func BenchmarkCountPrecision(b *testing.B) {
	b.ReportAllocs()

	value := SatoshiInBitcoin

	b.Run("current realisation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountPrecision(value)
		}
	})

	b.Run("using log10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = int(math.Log10(float64(value)))
		}
	})

	b.Run("using string symbols counting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len(fmt.Sprintf("%d", value)) - 1
		}
	})

}
