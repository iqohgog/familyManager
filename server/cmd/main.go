package main

import (
	"fmt"
	"net/http"
	"os"
	"v1/familyManager/configs"
	"v1/familyManager/pkg/db"
	"v1/familyManager/pkg/middleware"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.New(conf)
	fmt.Println(db)

	router := http.NewServeMux()
	// authServices := auth.NewAuthService(userRepository)

	stack := middleware.Chain(
		middleware.CORS,
	)
	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: app,
	}
	fmt.Println("Server is lestining on port 8081")
	server.ListenAndServe()
}
