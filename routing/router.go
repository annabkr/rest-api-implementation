package routing

import "net/http"

var (
	mux    *http.ServeMux
	router *Router
)

type Router struct {
	routes map[string]*Route
}

func Initialize() *http.ServeMux {
	mux = http.NewServeMux()
	router = &Router{
		routes: make(map[string]*Route),
	}
	return mux
}
