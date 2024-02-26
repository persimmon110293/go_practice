package usecases

import (
	"main/application/services"
	"net/http"
)

type IUserUsecase interface {
	GetUser(http.ResponseWriter, *http.Request) ([]byte, error)
}

type UserUsecase struct {
	app_service services.IUserAppService
}

func NewUserUsecase(user_app_service services.IUserAppService) *UserUsecase {
	return &UserUsecase{user_app_service}
}

func (usecase *UserUsecase) GetUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	return usecase.app_service.GetUser(w, r)
}
