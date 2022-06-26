package commands

import (
	bitcoinDomain "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	userDomain "github.com/F0rzend/simple-go-webserver/app/user/domain"
)

type ChangeBTCBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

type ChangeBTCBalanceCommandHandler struct {
	userRepository userDomain.UserRepository
	btcRepository  bitcoinDomain.BTCRepository
}

func NewChangeBTCBalanceCommand(
	userRepository userDomain.UserRepository,
	btcRepository bitcoinDomain.BTCRepository,
) (ChangeBTCBalanceCommandHandler, error) {
	if userRepository == nil {
		return ChangeBTCBalanceCommandHandler{}, ErrNilUserRepository
	}
	if btcRepository == nil {
		return ChangeBTCBalanceCommandHandler{}, ErrNilBTCRepository
	}

	return ChangeBTCBalanceCommandHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}, nil
}

func MustNewChangeBTCBalanceCommand(
	userRepository userDomain.UserRepository,
	btcRepository bitcoinDomain.BTCRepository,
) ChangeBTCBalanceCommandHandler {
	cmd, err := NewChangeBTCBalanceCommand(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}
	return cmd
}

func (h ChangeBTCBalanceCommandHandler) Handle(cmd ChangeBTCBalanceCommand) error {
	action, err := bitcoinDomain.NewBTCAction(cmd.Action)
	if err != nil {
		return err
	}

	return h.userRepository.Update(cmd.UserID, func(user *userDomain.User) (*userDomain.User, error) {
		btc, err := bitcoinDomain.NewBTC(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeBTCBalance(action, btc, h.btcRepository.Get()); err != nil {
			return nil, err
		}
		return user, nil
	})
}
