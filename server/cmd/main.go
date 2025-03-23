package main

import (
	"fmt"
	"net/http"
	"os"
	"v1/familyManager/configs"
	"v1/familyManager/internal/auth"
	"v1/familyManager/internal/user"
	"v1/familyManager/pkg/db"
	"v1/familyManager/pkg/middleware"
)

func App() http.Handler {
	// Configs
	conf := configs.LoadConfig()
	db := db.New(conf)
	router := http.NewServeMux()

	// Repository
	userRepository := user.NewUserRepository(db)

	// Service
	authServices := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authServices,
	})

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
	)
	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    string(":" + os.Getenv("SERVER_PORT")),
		Handler: app,
	}
	fmt.Println("Server is lestining on port", os.Getenv("SERVER_PORT"))
	server.ListenAndServe()
}
