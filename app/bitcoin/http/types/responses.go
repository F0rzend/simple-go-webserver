package types

import (
	"math/big"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
)

type BTCResponse struct {
	Price     *big.Float `json:"price"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func BTCToResponse(btc domain.BTCPrice) BTCResponse {
	return BTCResponse{
		Price:     btc.GetPrice().ToFloat(),
		UpdatedAt: btc.GetUpdatedAt(),
	}
}
