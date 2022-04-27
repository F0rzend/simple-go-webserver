package types

import (
	"math/big"
	"time"
)

type UserResponse struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	BTCBalance *big.Float `json:"btc_balance"`
	USDBalance *big.Float `json:"usd_balance"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type BTCResponse struct {
	Price     *big.Float `json:"price"`
	UpdatedAt time.Time  `json:"updated_at"`
}
