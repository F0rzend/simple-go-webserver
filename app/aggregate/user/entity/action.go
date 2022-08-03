package entity

type Action string

const (
	DepositUSDAction  Action = "deposit"
	WithdrawUSDAction Action = "withdraw"

	BuyBTCAction  Action = "buy"
	SellBTCAction Action = "sell"
)
