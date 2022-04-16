package types

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type Assembler struct{}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (a *Assembler) ResponseFromDomain(user domain.User) UserResponse {
	return UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Username:      user.Username,
		Email:         user.Email.Address,
		BitcoinAmount: user.Balance.BTC.ToFloat(),
		UsdBalance:    user.Balance.USD.ToFloat(),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
