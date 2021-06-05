package api

import (
	"fmt"
	"net/http"
)

func Initialize() {
	//ApiRouter := r.NewRouter()
	//ApiRouter.RegisterRoute(ApiRoot, "/", root)
	// r.ApiRoot.RegisterRoute("/", r.HandlerFunc(root))
	// r.ApiGetRecord.RegisterRoute("/record", getRecord)
}

func getRecord(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Get record")
	return nil
}
