package commands

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type ChangeBTCBalanceCommand struct {
	UserId uint64
	Action string
	Amount float64
}

type ChangeBTCBalanceCommandHandler struct {
	userRepository domain.UserRepository
	btcRepository  domain.BTCRepository
}

func NewChangeBTCBalanceCommand(
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
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
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
) ChangeBTCBalanceCommandHandler {
	cmd, err := NewChangeBTCBalanceCommand(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}
	return cmd
}

func (h ChangeBTCBalanceCommandHandler) Handle(cmd ChangeBTCBalanceCommand) error {
	action, err := domain.NewBTCAction(cmd.Action)
	if err != nil {
		return err
	}

	btc, err := h.btcRepository.Get()
	if err != nil {
		return err
	}

	return h.userRepository.Update(cmd.UserId, func(user *domain.User) (*domain.User, error) {
		if err := user.ChangeBTCBalance(action, domain.BTCFromFloat(cmd.Amount), btc.Price); err != nil {
			return nil, err
		}
		return user, nil
	})
}
