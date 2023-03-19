package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func reqStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: need to get the git commit and so on
		ok := true

		handleResponse(w, ok)
	}
}

func getStatusEndpoints() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", reqStatus())
	}
}
