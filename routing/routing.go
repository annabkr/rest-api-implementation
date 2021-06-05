package routing

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
}

func (r Route) RegisterRoute(pattern string, handler HandlerFunc) {
	r.Pattern = pattern
	mux.Handle(pattern, handler)
}

func (r Route) WithRouter(router *Router) {
	router.routes[r.Pattern] = &r
}

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := router.routes[r.URL.String()]; ok {
		if r.Method != route.Method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
