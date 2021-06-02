package api

import "net/http"

func root(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "hello world"}`))
	return nil
}
