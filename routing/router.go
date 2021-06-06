package routing

import "net/http"

type Route struct {
	Method  string
	Pattern string
}

type Router struct {
	routes map[string]*Route
}

func GetMux() *http.ServeMux {
	return mux
}

func GetRouter() *Router {
	return router
}

func (r Router) RegisterRoute(method string, pattern string, handler HandlerFunc) {
	GetMux().Handle(pattern, handler)
	GetRouter().routes[pattern] = &Route{
		Method:  method,
		Pattern: pattern,
	}
}
