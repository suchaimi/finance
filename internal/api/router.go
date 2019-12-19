package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter provide a handler api service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
