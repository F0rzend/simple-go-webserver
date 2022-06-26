package types

import (
	"errors"
	"net/http"
)

var ErrBadRequest = errors.New("validation error")

type SetBTCPriceRequest struct {
	Price float64 `json:"price"`
}

func (r SetBTCPriceRequest) Bind(_ *http.Request) error {
	if r.Price <= 0 {
		return ErrBadRequest
	}
	return nil
}
