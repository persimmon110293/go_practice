package repositories

import (
	entity "main/domain/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Find() *gorm.DB
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) Find() *gorm.DB {
	db := dbInit()

	user_entity := []entity.User{}
	return db.Find(&user_entity, 1) // TODO: ここでのgormの使い方を調べる
}

func dbInit() *gorm.DB {
	dsn := "root:password@tcp(mysql:3306)/practice?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
