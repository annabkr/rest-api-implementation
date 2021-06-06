package routing

import (
	"net/http"
)

var (
	mux    *http.ServeMux
	router *Router
)

func Initialize() {
	mux = http.NewServeMux()
	router = &Router{
		routes: make(map[string]*Route),
	}
}
