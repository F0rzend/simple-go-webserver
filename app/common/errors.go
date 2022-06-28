package common

import (
	"fmt"
)

// Repository errors.
// These types are used to identify the error cause
// in service layer when a repository fails.
type (
	ErrUserNotFound      uint64 // rises when a user with the given ID was not found
	ErrUserAlreadyExists uint64 // rises when a user with the given ID is already exists
)

// Domain errors.
// These types are used to identify the error cause.
type (
	ErrNegativeCurrency  float64 // rises when a user tries to pass a negative amount.
	ErrInsufficientFunds float64 // rises when a user has insufficient funds. The value is the amount of funds required.
)

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with id %d not found", e)
}

func (e ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("user with id %d already exists", e)
}

func (e ErrNegativeCurrency) Error() string {
	return fmt.Sprintf("negative currency: %f", e)
}

func (e ErrInsufficientFunds) Error() string {
	return fmt.Sprintf("insufficient funds: %f", e)
}
