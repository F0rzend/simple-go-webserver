package domain

type UserRepository interface {
	Create(user *User) error
	Get(id uint64) (*User, error)
	Update(
		id uint64, updFunc func(currentUser *User) (updatedUser *User, err error),
	) error
	Delete(id uint64) error
}
