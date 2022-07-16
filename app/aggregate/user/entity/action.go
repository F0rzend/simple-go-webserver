package entity

type Action string

var (
	DepositUSDAction  Action = "deposit"
	WithdrawUSDAction Action = "withdraw"

	BuyBTCAction  Action = "buy"
	SellBTCAction Action = "sell"
)
