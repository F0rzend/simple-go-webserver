package bitcoinservice

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (bs *BitcoinService) SetBTCPrice(newPrice float64) error {
	if newPrice < 0 {
		return common.NewApplicationError(
			http.StatusBadRequest,
			"The price cannot be negative. Please pass a number greater than 0",
		)
	}

	return bs.bitcoinRepository.SetPrice(bitcoinentity.NewUSD(newPrice))
}
