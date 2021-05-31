package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pkg/errors"
)

//APIs basically have a few things:
//A route - e.g. '/users'
//A handler - function where the route ends up
//A server - to listen for requests that are directed to the handler

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(`{"message": "hello world"}`))
	return nil
}

func ListenAndServe() error {
	listener, err := net.Listen("tcp", ":0")
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
