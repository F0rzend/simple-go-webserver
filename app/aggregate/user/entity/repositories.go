package userentity

//go:generate moq -out "../repositories/mock.gen.go" -pkg userrepositories . UserRepository:MockUserRepository
type UserRepository interface {
	Get(id uint64) (*User, error)
	Save(user *User) error
	Delete(id uint64) error
}
