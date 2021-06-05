package main

import (
	"log"
	"net/http"

	"github.com/annabkr/rest-api-implementation/api"
	r "github.com/annabkr/rest-api-implementation/routing"
)

func main() {
	mux := r.Initialize()
	api.Initialize()
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
