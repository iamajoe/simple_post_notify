package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func reqNotifySend() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var raw struct {
			Msg string `json:"msg"`
			Mod string `json:"mod"`
		}
		err := json.NewDecoder(r.Body).Decode(&raw)
		if err != nil {
			handleErrResponse(w, err)
			return
		}

		_, err = sendNotifyMsg(raw.Msg, raw.Mod)
		if err != nil {
			handleErrResponse(w, err)
			return
		}

		handleResponse(w, true)
	}
}

func getNotifyEndpoints() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/send", reqNotifySend())
	}
}
