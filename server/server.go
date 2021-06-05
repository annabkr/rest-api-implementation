package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/pkg/errors"

	log "github.com/annabkr/rest-api-implementation/utils/logger"
)

func ListenAndServe(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "no available ports")
	}

	log.Info(fmt.Sprintf("Listening on port %s", port))

	err = http.Serve(listener, nil)
	if err != nil {
		return errors.Wrap(err, "unable to start server")
	}
	return nil
}
