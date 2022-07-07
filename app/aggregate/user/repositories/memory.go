package repositories

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type MemoryUserRepository struct {
	users map[uint64]*entity.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uint64]*entity.User),
	}
}

func (r *MemoryUserRepository) Create(user *entity.User) error {
	if _, ok := r.users[user.ID]; ok {
		return ErrUserAlreadyExists
	}

	for _, registeredUser := range r.users {
		if registeredUser.Email == user.Email {
			return ErrUserAlreadyExists
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Get(id uint64) (*entity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepository) Update(
	id uint64,
	updFunc func(*entity.User) (*entity.User, error),
) error {
	currentUser, ok := r.users[id]

	if !ok {
		return ErrUserNotFound
	}

	updatedUser, err := updFunc(currentUser)
	if err != nil {
		return err
	}
	r.users[id] = updatedUser

	return nil
}

func (r *MemoryUserRepository) Delete(id uint64) error {
	_, ok := r.users[id]
	if !ok {
		return ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}
