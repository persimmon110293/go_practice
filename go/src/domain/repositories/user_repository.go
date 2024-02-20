package repositories

import (
	"main/domain/entities"
)

type IUserRepository interface {
	Find() []entities.User
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) Find() []entities.User {
	// TODO: ここでDBからデータを取り出す
	user := []entities.User{}
	return user
}
