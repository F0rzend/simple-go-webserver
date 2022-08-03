package userhandlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var (
	ErrEmptyAction = common.NewApplicationError(
		http.StatusBadRequest,
		"Action cannot be empty",
	)
	ErrZeroAmount = common.NewApplicationError(
		http.StatusBadRequest,
		"Amount can't be zero",
	)
)
