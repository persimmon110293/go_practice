package usecases

import (
	"main/application/services"
	"main/domain/entities"
	"net/http"
)

type IUserUsecase interface {
	GetUser(http.ResponseWriter, *http.Request) []entities.User
}

type UserUsecase struct {
	app_service services.IUserAppService
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (usecase *UserUsecase) GetUser(w http.ResponseWriter, r *http.Request) []entities.User {
	// TODO: ここでapp_serviceを呼ぶ
	return usecase.app_service.GetUser(w, r)
}
