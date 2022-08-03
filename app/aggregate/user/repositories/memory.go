package userrepositories

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type MemoryUserRepository struct {
	users map[uint64]*userentity.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uint64]*userentity.User),
	}
}

func (r *MemoryUserRepository) Save(user *userentity.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Get(id uint64) (*userentity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepository) Delete(id uint64) error {
	_, ok := r.users[id]
	if !ok {
		return ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}
