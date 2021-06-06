package main

import (
	"fmt"
	"net/http"

	"github.com/annabkr/rest-api-implementation/api"
	r "github.com/annabkr/rest-api-implementation/routing"
	log "github.com/annabkr/rest-api-implementation/utils/logger"
)

func main() {
	r.Initialize()
	mux := r.GetMux()
	api.Initialize()

	port := ":3000"
	log.Info(fmt.Sprintf("Listening on port %s", port))
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}
