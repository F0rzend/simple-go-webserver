package handlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var ErrInvalidEmail = common.NewApplicationError(
	http.StatusBadRequest,
	"Email is not valid",
)
