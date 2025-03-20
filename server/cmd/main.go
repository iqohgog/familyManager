package main

import (
	"fmt"
	"v1/familyManager/configs"
)

// func App() http.Handler {

// }

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
}
