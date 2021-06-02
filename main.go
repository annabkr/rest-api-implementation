package main

import (
	"log"

	"github.com/annabkr/rest-api-implementation/api"
	"github.com/annabkr/rest-api-implementation/server"
)

func main() {
	api.InitializeRoutes()
	onPort := ":3000"
	err := server.ListenAndServe(onPort)
	if err != nil {
		log.Fatal(err)
	}
}
