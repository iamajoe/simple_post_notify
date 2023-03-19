package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func setEndpoints(mux *chi.Mux) {
	mux.Route("/status", getStatusEndpoints())
	mux.Route("/notify", getNotifyEndpoints())
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	// if origin == "http://example.com" {
	// 	return true
	// }
	// return false
	return true
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.NoCache)
	r.Use(middleware.Timeout(60 * time.Second))

	// TODO: should limit cors a bit more
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:    AllowOriginFunc,
		AllowedMethods:     []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // []string{"*"}, //
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		MaxAge:             300, // Maximum value not ignored by any of major browsers
		Debug:              false,
	}))

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		handleErrResponse(w, NewError(http.StatusMethodNotAllowed, errors.New("not allowed")))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		handleErrResponse(w, NewError(http.StatusNotFound, errors.New("not allowed")))
	})

	setEndpoints(r)

	return r
}

func InitServer(address string) {
	log.Printf("listening at %s \n", address)
	err := http.ListenAndServe(address, getRouter())
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
