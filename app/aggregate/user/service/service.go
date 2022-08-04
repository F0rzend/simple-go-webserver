package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserService struct {
	userRepository UserRepository
	priceGetter    BTCPriceGetter

	userIDGenerator func() uint64
}

//go:generate moq -out "mock_user_repository.gen.go" . UserRepository:MockUserRepository
type UserRepository interface {
	Get(id uint64) (*userentity.User, error)
	Save(user *userentity.User) error
	Delete(id uint64) error
}

//go:generate moq -out "mock_btc_price_getter.gen.go" . BTCPriceGetter:MockBTCPriceGetter
type BTCPriceGetter interface {
	GetPrice() bitcoinentity.BTCPrice
}

func NewUserService(
	userRepository UserRepository,
	bitcoinRepository BTCPriceGetter,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		priceGetter:    bitcoinRepository,

		userIDGenerator: getUserIDGenerator(),
	}
}
