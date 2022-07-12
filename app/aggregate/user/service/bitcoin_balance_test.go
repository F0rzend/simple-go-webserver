package service

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func TestUserService_ChangeBitcoinBalance_InvalidAction(t *testing.T) {
	t.Parallel()

	const action = "invalid"

	userRepository := &userRepositories.MockUserRepository{}
	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

	service := NewUserService(userRepository, bitcoinRepository)

	expectError := common.NewServiceError(http.StatusBadRequest, fmt.Sprintf("Invalid action: %s", action))

	err := service.ChangeBitcoinBalance(ChangeBitcoinBalanceCommand{
		Action: action,
	})

	assert.Equal(t, expectError, err)

	assert.Len(t, userRepository.UpdateCalls(), 0)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 0)
}

func TestUserService_ChangeBitcoinBalance_UserNotFound(t *testing.T) {
	t.Parallel()

	const (
		userID = 42
		action = "buy"
	)

	userRepository := &userRepositories.MockUserRepository{
		UpdateFunc: func(_ uint64, _ func(*userEntity.User) (*userEntity.User, error)) error {
			return userRepositories.ErrUserNotFound
		},
	}
	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

	service := NewUserService(userRepository, bitcoinRepository)

	expectError := common.NewServiceError(http.StatusNotFound, fmt.Sprintf("User with id %d not found", userID))

	err := service.ChangeBitcoinBalance(ChangeBitcoinBalanceCommand{
		UserID: userID,
		Action: action,
	})

	assert.Equal(t, expectError, err)

	assert.Len(t, userRepository.UpdateCalls(), 1)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 0)
}

func TestUserService_ChangeBitcoinBalance_NegativeCurrency(t *testing.T) {
	t.Parallel()

	const (
		userID = 42
		action = "buy"
		amount = -1
	)

	userRepository := &userRepositories.MockUserRepository{
		UpdateFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
			_, err := updateFunc(&userEntity.User{})
			return err
		},
	}
	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

	service := NewUserService(userRepository, bitcoinRepository)

	expectError := common.NewServiceError(
		http.StatusBadRequest,
		"The amount of currency cannot be negative. Please pass a number greater than 0",
	)

	err := service.ChangeBitcoinBalance(ChangeBitcoinBalanceCommand{
		UserID: userID,
		Action: action,
		Amount: amount,
	})

	assert.Equal(t, expectError, err)

	assert.Len(t, userRepository.UpdateCalls(), 1)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 0)
}

func TestUserService_ChangeBitcoinBalance_ErrInsufficientFunds(t *testing.T) {
	t.Parallel()

	const (
		userID = 42
		amount = 1
	)

	testCases := []struct {
		name   string
		action string
	}{
		{
			name:   "selling",
			action: "sell",
		},
		{
			name:   "buying",
			action: "buy",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{
				UpdateFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
					zeroUSD, _ := bitcoinEntity.NewUSD(0)
					zeroBTC, _ := bitcoinEntity.NewBTC(0)

					_, err := updateFunc(
						&userEntity.User{
							Balance: userEntity.Balance{
								USD: zeroUSD,
								BTC: zeroBTC,
							},
						},
					)
					return err
				},
			}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{
				GetPriceFunc: func() bitcoinEntity.BTCPrice {
					price, _ := bitcoinEntity.NewUSD(1)
					return bitcoinEntity.NewBTCPrice(price, time.Now())
				},
			}

			service := NewUserService(userRepository, bitcoinRepository)

			expectError := common.NewServiceError(
				http.StatusBadRequest,
				fmt.Sprintf("The user does not have enough funds to %s BTC", tc.action),
			)

			err := service.ChangeBitcoinBalance(ChangeBitcoinBalanceCommand{
				UserID: userID,
				Action: tc.action,
				Amount: amount,
			})

			assert.Equal(t, expectError, err)

			assert.Len(t, userRepository.UpdateCalls(), 1)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
		})
	}
}

func TestUserService_ChangeBitcoinBalance_Success(t *testing.T) {
	t.Parallel()

	const (
		userID = 42
		action = "buy"
		amount = 1
	)

	userRepository := &userRepositories.MockUserRepository{
		UpdateFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
			userBalance, _ := bitcoinEntity.NewUSD(10)
			zeroBTC, _ := bitcoinEntity.NewBTC(0)

			_, err := updateFunc(
				&userEntity.User{
					Balance: userEntity.Balance{
						USD: userBalance,
						BTC: zeroBTC,
					},
				},
			)
			return err
		},
	}
	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{
		GetPriceFunc: func() bitcoinEntity.BTCPrice {
			price, _ := bitcoinEntity.NewUSD(1)
			return bitcoinEntity.NewBTCPrice(price, time.Now())
		},
	}

	service := NewUserService(userRepository, bitcoinRepository)

	err := service.ChangeBitcoinBalance(ChangeBitcoinBalanceCommand{
		UserID: userID,
		Action: action,
		Amount: amount,
	})

	assert.NoError(t, err)

	assert.Len(t, userRepository.UpdateCalls(), 1)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
}
