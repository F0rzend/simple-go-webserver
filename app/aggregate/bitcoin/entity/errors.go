package bitcoinentity

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var ErrNegativeCurrency = common.NewApplicationError(
	http.StatusBadRequest,
	"The amount of currency cannot be negative. Please pass a number greater than 0",
)
