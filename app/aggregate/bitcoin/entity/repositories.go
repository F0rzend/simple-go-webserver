package entity

//go:generate moq -out "../../../tests/bitcoin_repository.gen.go" -pkg tests . BTCRepository:MockBTCRepository
type BTCRepository interface {
	Get() BTCPrice
	SetPrice(price USD) error
}
