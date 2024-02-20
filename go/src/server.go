package main

import (
	"encoding/json"
	"net/http"

	"main/handlers"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		message := struct {
			Message string `json:"message"`
		}{
			Message: "pong",
		}
		response, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/user", handlers.UserHandler.GetUserHandler)
	http.ListenAndServe(":8080", nil)
}
