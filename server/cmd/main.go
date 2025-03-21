package main

import (
	"fmt"
	"v1/familyManager/configs"
	"v1/familyManager/pkg/db"
)

// func App() http.Handler {

// }

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
	db := db.New(conf)
	fmt.Println(db)
}
