package api

import "net/http"

import "github.com/gorilla/mux"

// NewRouter provide a handler api service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	return router, nil
}
