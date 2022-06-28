package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type ChangeBTCBalance struct {
	UserID uint64
	Action string
	Amount float64
}

type ChangeBTCBalanceHandler struct {
	userRepository userEntity.UserRepository
	btcRepository  bitcoinEntity.BTCRepository
}

func NewChangeBTCBalanceCommand(
	userRepository userEntity.UserRepository,
	btcRepository bitcoinEntity.BTCRepository,
) (ChangeBTCBalanceHandler, error) {
	if userRepository == nil {
		return ChangeBTCBalanceHandler{}, ErrNilUserRepository
	}
	if btcRepository == nil {
		return ChangeBTCBalanceHandler{}, ErrNilBTCRepository
	}

	return ChangeBTCBalanceHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}, nil
}

func MustNewChangeBTCBalanceHandler(
	userRepository userEntity.UserRepository,
	btcRepository bitcoinEntity.BTCRepository,
) ChangeBTCBalanceHandler {
	cmd, err := NewChangeBTCBalanceCommand(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}
	return cmd
}

func (h ChangeBTCBalanceHandler) Handle(cmd ChangeBTCBalance) error {
	action, err := bitcoinEntity.NewBTCAction(cmd.Action)
	if err != nil {
		return err
	}

	return h.userRepository.Update(cmd.UserID, func(user *userEntity.User) (*userEntity.User, error) {
		btc, err := bitcoinEntity.NewBTC(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeBTCBalance(action, btc, h.btcRepository.Get()); err != nil {
			return nil, err
		}
		return user, nil
	})
}
