package queries

type Queries struct {
	GetUser        GetUserQueryHandler
	GetUserBalance GetUserBalanceQueryHandler

	GetBTC GetBTCQueryHandler
}
