package userservice

import (
	"fmt"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func getUserIDGenerator() func() uint64 {
	var id uint64
	return func() uint64 {
		id++
		return id
	}
}

func (us *UserService) CreateUser(name, username, email string) (uint64, error) {
	user, err := userentity.NewUser(
		us.userIDGenerator(),
		name,
		username,
		email,
		0,
		0,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	err = us.userRepository.Save(user)
	if err != nil {
		return 0, fmt.Errorf("error saving user: %w", err)
	}

	return user.ID, nil
}
