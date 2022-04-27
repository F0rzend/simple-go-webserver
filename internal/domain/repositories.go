package domain

import "fmt"

// Repository errors.
// These types are used to identify the error cause
// in application layer when a repository fails.
type (
	ErrUserNotFound      uint64 // rises when a user with the given ID was not found
	ErrUserAlreadyExists uint64 // rises when a user with the given ID is already exists
)

func (err ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with id %d not found", err)
}

func (err ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("user with id %d already exists", err)
}

type UserRepository interface {
	Create(user *User) error
	Get(id uint64) (*User, error)
	// Update accept id of user, get it and pass in updFunc.
	// updFunc should update the user and return the updated user
	Update(
		id uint64, updFunc func(*User) (*User, error),
	) error
	Delete(id uint64) error
}

type BTCRepository interface {
	Get() (BTCPrice, error)
	SetPrice(price USD) error
}
