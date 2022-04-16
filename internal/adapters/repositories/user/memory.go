package userRepository

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

var (
	_ domain.UserRepository = &MemoryUserRepository{}
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
		return domain.ErrUserAlreadyExists(user.ID)
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Get(id uint64) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, domain.ErrUserNotFound(id)
	}
	return user, nil
}

func (r *MemoryUserRepository) Update(
	id uint64,
	updFunc func(*domain.User) (*domain.User, error),
) error {
	currentUser, ok := r.users[id]

	if !ok {
		return domain.ErrUserNotFound(id)
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
		return domain.ErrUserNotFound(id)
	}
	delete(r.users, id)
	return nil
}
