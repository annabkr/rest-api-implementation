package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/annabkr/rest-api-implementation/utils"

	"github.com/pkg/errors"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

var routes = map[string]bool{"/": true, "/users": true}

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := routes[r.URL.Path]; !ok {
		err := utils.NewNotFoundError("page not found", nil, "")
		http.Error(w, err.Error(), err.StatusCode)
		return
	}

	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.Handle("/", HandlerFunc(root))
	err := ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "hello world"}`))
	return nil
}

func ListenAndServe() error {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return errors.Wrap(err, "no available ports")
	}

	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("Using port: %+v\n", port)

	err = http.Serve(listener, nil)
	if err != nil {
		return errors.Wrap(err, "unable to start server")
	}
	return nil
}
