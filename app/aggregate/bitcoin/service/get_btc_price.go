package service

import (
	"errors"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type GetBTCPriceHandler struct {
	btcRepository entity.BTCRepository
}

var ErrNilBTCRepository = errors.New("btc repository is nil")

func NewGetBTCCommand(btcRepository entity.BTCRepository) (GetBTCPriceHandler, error) {
	if btcRepository == nil {
		return GetBTCPriceHandler{}, ErrNilBTCRepository
	}

	return GetBTCPriceHandler{
		btcRepository: btcRepository,
	}, nil
}

func MustNewGetBTCCommand(btcRepository entity.BTCRepository) GetBTCPriceHandler {
	cmd, err := NewGetBTCCommand(btcRepository)
	if err != nil {
		panic(err)
	}

	return cmd
}

func (h *GetBTCPriceHandler) Handle() entity.BTCPrice {
	return h.btcRepository.Get()
}
