package domain

import (
	"fmt"
)

// Domain errors.
// These types are used to identify the error cause.
type (
	ErrNegativeCurrency float64 // rises when a user tries to pass a negative amount.
)

func (e ErrNegativeCurrency) Error() string {
	return fmt.Sprintf("negative currency: %f", e)
}
