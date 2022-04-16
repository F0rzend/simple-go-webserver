package userRepository

import (
	"errors"

	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

var (
	_ domain.UserRepository = &MemoryUserRepository{}

	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type MemoryUserRepository struct {
	users map[uint64]*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uint64]*domain.User),
	}
}

func (r *MemoryUserRepository) Create(user *domain.User) error {
	if _, ok := r.users[user.ID]; ok {
		return ErrUserAlreadyExists
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Get(id uint64) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepository) Update(
	id uint64,
	updFunc func(*domain.User) (*domain.User, error),
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
