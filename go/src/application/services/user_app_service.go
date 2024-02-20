package services

import (
	"main/domain/entities"
	"main/domain/repositories"
	"net/http"
)

type IUserAppService interface {
	GetUser(http.ResponseWriter, *http.Request) []entities.User
}

type UserAppService struct {
	repository repositories.IUserRepository
}

func NewUserAppService() IUserAppService {
	return &UserAppService{}
}

func (app_service *UserAppService) GetUser(w http.ResponseWriter, r *http.Request) []entities.User {
	// TODO: ここでrepositoryを呼ぶ
	// TODO: ここでuserエンティティをreturnする
	user := app_service.repository.Find()
	return user
}
