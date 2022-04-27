package types

import (
	"time"
)

type UserResponse struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	BTCBalance float64   `json:"btc_balance"`
	USDBalance float64   `json:"usd_balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BTCResponse struct {
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}
