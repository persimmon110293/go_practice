package handlers

import (
	"main/usecases"
	"net/http"
)

type UserHandler struct {
	usecase usecases.IUserUsecase
}

func NewUserHandler(usecase usecases.IUserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

func (user_handler *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user_handler.usecase.GetUser(w, r)
}
