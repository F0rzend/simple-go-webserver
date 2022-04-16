package queries

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type GetUserQueryHandler struct {
	repository domain.UserRepository
}

func NewGetUserQuery(repository domain.UserRepository) *GetUserQueryHandler {
	return &GetUserQueryHandler{repository: repository}
}

func (h *GetUserQueryHandler) Handle(id uint64) (*domain.User, error) {
	user, err := h.repository.Get(id)
	return user, err
}
