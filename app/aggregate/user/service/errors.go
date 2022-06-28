package service

import "errors"

var (
	ErrNilUserRepository = errors.New("user repository is nil")
	ErrNilBTCRepository  = errors.New("btc repository is nil")
)
