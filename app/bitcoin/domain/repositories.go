package domain

//go:generate moq -out "../repositories/mock.gen.go" -pkg btcrepositories . BTCRepository:MockBTCRepository
type BTCRepository interface {
	Get() BTCPrice
	SetPrice(price USD) error
}
