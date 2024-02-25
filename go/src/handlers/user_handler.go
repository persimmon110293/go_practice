package handlers

import (
	"main/usecases"
	"net/http"
)

type UserHandler struct{}

func (user_handler *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	usecase := usecases.NewUserUsecase()
	usecase.GetUser(w, r)
}
