package queries

import (
	domain2 "github.com/F0rzend/simple-go-webserver/app/user/domain"
)

type GetUserQueryHandler struct {
	userRepository domain2.UserRepository
}

func NewGetUserQuery(userRepository domain2.UserRepository) (GetUserQueryHandler, error) {
	if userRepository == nil {
		return GetUserQueryHandler{}, ErrNilUserRepository
	}

	return GetUserQueryHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewGetUserQuery(userRepository domain2.UserRepository) GetUserQueryHandler {
	handler, err := NewGetUserQuery(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *GetUserQueryHandler) Handle(id uint64) (*domain2.User, error) {
	user, err := h.userRepository.Get(id)
	return user, err
}
