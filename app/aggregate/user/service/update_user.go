package service

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) UpdateUser(userID uint64, name, email *string) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	if name == nil && email == nil {
		return nil
	}

	if name != nil {
		user.Name = *name
	}

	if email != nil {
		newEmail, err := entity.ParseEmail(*email)
		if err != nil {
			return err
		}
		user.Email = newEmail
	}

	user.UpdatedAt = time.Now()

	return us.userRepository.Save(user)
}
