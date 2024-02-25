package services

import (
	"encoding/json"
	repository "main/domain/repositories"
	"net/http"
)

type IUserAppService interface {
	GetUser(http.ResponseWriter, *http.Request) ([]byte, error)
}

type UserAppService struct {
	repository repository.IUserRepository
}

func NewUserAppService() IUserAppService {
	return &UserAppService{}
}

func (app_service *UserAppService) GetUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	user := app_service.repository.Find()
	return json.Marshal(user)
}
