package bitcoinentity

//go:generate moq -out "../repositories/mock.gen.go" -pkg bitcoinrepositories . BTCRepository:MockBTCRepository
type BTCRepository interface {
	GetPrice() BTCPrice
	SetPrice(price USD) error
}
