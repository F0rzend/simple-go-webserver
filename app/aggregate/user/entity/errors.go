package entity

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var ErrInsufficientFunds = common.NewApplicationError(
	http.StatusBadRequest,
	"The user does not have enough funds",
)
