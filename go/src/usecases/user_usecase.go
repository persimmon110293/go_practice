package usecases

import (
	service "main/application/services"
	"net/http"
)

type IUserUsecase interface {
	GetUser(http.ResponseWriter, *http.Request) ([]byte, error)
}

type UserUsecase struct {
	app_service service.IUserAppService
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (usecase *UserUsecase) GetUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	return usecase.app_service.GetUser(w, r)
}
