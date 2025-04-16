package main

import (
	"fmt"
	"net/http"

	"purpleProject/configs"
	"purpleProject/internal/auth"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthConfigDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081...")
	server.ListenAndServe()
}
