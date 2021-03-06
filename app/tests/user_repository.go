package tests

import (
	"time"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func NewMockUserRepository() userEntity.UserRepository {
	now := time.Now()
	users := map[uint64]*userEntity.User{
		1: userEntity.MustNewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		),
		2: userEntity.MustNewUser(
			2,
			"Jane",
			"Doe",
			"janedoe@mail.com",
			100,
			100,
			now,
			now,
		),
	}
	return &MockUserRepository{
		CreateFunc: func(user *userEntity.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := userEntity.NewUser(
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
				return common.ErrUserNotFound(id)
			}
			return nil
		},
		GetFunc: func(id uint64) (*userEntity.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, common.ErrUserNotFound(id)
			}
			return user, nil
		},
		UpdateFunc: func(id uint64, updFunc func(*userEntity.User) (*userEntity.User, error)) error {
			user, ok := users[id]
			if !ok {
				return common.ErrUserNotFound(id)
			}
			userCopy := *user
			_, err := updFunc(&userCopy)
			return err
		},
	}
}
