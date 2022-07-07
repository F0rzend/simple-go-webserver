package tests

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
)

func NewMockUserRepository() entity.UserRepository {
	now := time.Now()
	users := map[uint64]*entity.User{
		1: must(entity.NewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		)),
		2: must(entity.NewUser(
			2, //nolint:gomnd
			"Jane",
			"Doe",
			"janedoe@mail.com",
			100, //nolint:gomnd
			100, //nolint:gomnd
			now,
			now,
		)),
	}
	return &MockUserRepository{
		CreateFunc: func(user *entity.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := entity.NewUser(
				user.ID,
				user.Name,
				user.Username,
				user.Email.Address,
				btc,
				usd,
				now,
				now,
			)
			return err
		},
		DeleteFunc: func(id uint64) error {
			if _, ok := users[id]; !ok {
				return repositories.ErrUserNotFound
			}
			return nil
		},
		GetFunc: func(id uint64) (*entity.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, repositories.ErrUserNotFound
			}
			return user, nil
		},
		UpdateFunc: func(id uint64, updFunc func(*entity.User) (*entity.User, error)) error {
			user, ok := users[id]
			if !ok {
				return repositories.ErrUserNotFound
			}
			userCopy := *user
			_, err := updFunc(&userCopy)
			return err
		},
	}
}
