package userservice

import (
	"fmt"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) UpdateUser(userID uint64, name, email *string) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if name == nil && email == nil {
		return nil
	}

	if name != nil {
		user.Name = *name
	}

	if email != nil {
		newEmail, err := userentity.ParseEmail(*email)
		if err != nil {
			return fmt.Errorf("failed to parse email: %w", err)
		}
		user.Email = newEmail
	}

	user.UpdatedAt = time.Now()

	err = us.userRepository.Save(user)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
