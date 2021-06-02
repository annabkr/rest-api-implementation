package api

import (
	"net/http"

	r "github.com/annabkr/rest-api-implementation/routing"
)

func InitializeRoutes() {
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/", r.HandlerFunc(root))
}
