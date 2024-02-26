package services

import (
	"encoding/json"
	"main/domain/repositories"
	"net/http"
)

type IUserAppService interface {
	GetUser(http.ResponseWriter, *http.Request) ([]byte, error)
}

type UserAppService struct {
	repository repositories.IUserRepository
}

func NewUserAppService(repository repositories.IUserRepository) IUserAppService {
	return &UserAppService{repository}
}

func (app_service *UserAppService) GetUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	user := app_service.repository.Find()
	return json.Marshal(user)
}
