package entity

//go:generate moq -out "../repositories/mock.gen.go" -pkg repositories . BTCRepository:MockBTCRepository
type BTCRepository interface {
	GetPrice() BTCPrice
	SetPrice(price USD) error
}
