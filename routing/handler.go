package routing

import (
	"fmt"
	"net/http"

	"github.com/annabkr/rest-api-implementation/utils/errors"
	log "github.com/annabkr/rest-api-implementation/utils/logger"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.String()
	route, ok := GetRouter().routes[pattern]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != route.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := h(w, r); err != nil {
		switch e := err.(type) {
		case *errors.Error:
			log.Err(fmt.Sprintf("returning %d for request %s %s: %s", e.StatusCode(), r.Method, pattern, e.Error()))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(e.StatusCode())
			_, err = w.Write(e.Json())
			if err != nil {
				log.Warn(fmt.Sprintf("failed to write: %s", err.Error()))
			}
		default:
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}
}
