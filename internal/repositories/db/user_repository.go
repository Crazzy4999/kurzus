package db

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
