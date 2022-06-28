package userrepositories

import (
	domain2 "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

var _ domain2.UserRepository = &MemoryUserRepository{}

type MemoryUserRepository struct {
	users map[uint64]*domain2.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uint64]*domain2.User),
	}
}

func (r *MemoryUserRepository) Create(user *domain2.User) error {
	if _, ok := r.users[user.ID]; ok {
		return common.ErrUserAlreadyExists(user.ID)
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Get(id uint64) (*domain2.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, common.ErrUserNotFound(id)
	}
	return user, nil
}

func (r *MemoryUserRepository) Update(
	id uint64,
	updFunc func(*domain2.User) (*domain2.User, error),
) error {
	currentUser, ok := r.users[id]

	if !ok {
		return common.ErrUserNotFound(id)
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
		return common.ErrUserNotFound(id)
	}
	delete(r.users, id)
	return nil
}
