package api

import (
	"fmt"
	"net/http"

	r "github.com/annabkr/rest-api-implementation/routing"
	"github.com/annabkr/rest-api-implementation/utils/errors"
)

func Initialize() {
	apiRouter := r.GetRouter()
	apiRouter.RegisterRoute("GET", "/", root)
	apiRouter.RegisterRoute("GET", "/record", getRecord)
}

func getRecord(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Get record")
	return errors.NewForbiddenError("forbidden", nil, "123")
}
