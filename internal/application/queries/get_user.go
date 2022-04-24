package queries

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type GetUserQueryHandler struct {
	userRepository domain.UserRepository
}

func NewGetUserQuery(userRepository domain.UserRepository) (GetUserQueryHandler, error) {
	if userRepository == nil {
		return GetUserQueryHandler{}, ErrNilUserRepository
	}

	return GetUserQueryHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewGetUserQuery(userRepository domain.UserRepository) GetUserQueryHandler {
	handler, err := NewGetUserQuery(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *GetUserQueryHandler) Handle(id uint64) (*domain.User, error) {
	user, err := h.userRepository.Get(id)
	return user, err
}
