package handlers

import (
	"fmt"
	"main/application/services"
	"main/domain/repositories"
	"main/usecases"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: ここでDIしているのは妥当なのか調べる
	user_repository := repositories.NewUserRepository()
	user_app_service := services.NewUserAppService(user_repository)
	usecase := usecases.NewUserUsecase(user_app_service)
	user, _ := usecase.GetUser(w, r)

	fmt.Printf("%s", user)
}
