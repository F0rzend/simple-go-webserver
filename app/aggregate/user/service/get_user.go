package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type GetUserHandler struct {
	userRepository entity.UserRepository
}

func NewGetUserQuery(userRepository entity.UserRepository) (GetUserHandler, error) {
	if userRepository == nil {
		return GetUserHandler{}, ErrNilUserRepository
	}

	return GetUserHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewGetUserHandler(userRepository entity.UserRepository) GetUserHandler {
	handler, err := NewGetUserQuery(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *GetUserHandler) Handle(id uint64) (*entity.User, error) {
	user, err := h.userRepository.Get(id)
	return user, err
}
