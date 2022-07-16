package entity

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var (
	ErrNegativeCurrency = common.NewApplicationError(
		http.StatusBadRequest,
		"The amount of currency cannot be negative. Please pass a number greater than 0",
	)
	ErrInvalidUSDAction = common.NewApplicationError(
		http.StatusBadRequest,
		fmt.Sprintf(
			"You must specify a valid action. Available actions: %s",
			strings.Join(GetUSDActions(), ", "),
		),
	)
	ErrInvalidBTCAction = common.NewApplicationError(
		http.StatusBadRequest,
		fmt.Sprintf(
			"You must specify a valid action. Available actions: %s",
			strings.Join(GetBTCActions(), ", "),
		),
	)
)
