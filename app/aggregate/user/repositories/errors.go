package userrepositories

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var ErrUserNotFound = common.NewApplicationError(
	http.StatusNotFound,
	"User not found",
)
