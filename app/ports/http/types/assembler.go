package types

import "github.com/F0rzend/simple-go-webserver/app/domain"

type Assembler struct{}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (a *Assembler) UserToResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email.Address,
		BTCBalance: user.Balance.BTC.ToFloat(),
		USDBalance: user.Balance.USD.ToFloat(),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func (a *Assembler) BTCToResponse(btc domain.BTCPrice) BTCResponse {
	return BTCResponse{
		Price:     btc.GetPrice().ToFloat(),
		UpdatedAt: btc.GetUpdatedAt(),
	}
}
