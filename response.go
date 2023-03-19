package main

import (
	"encoding/json"
	"net/http"
)

type serverErrorCoded interface {
	Error() string
	StatusCode() int
}

func handleResponseRaw(w http.ResponseWriter, code int, data interface{}) {
	// prepare the response
	resData := struct {
		Ok   bool        `json:"ok"`
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
		Err  string      `json:"err,omitempty"`
	}{true, code, data, ""}
	if code > http.StatusOK+99 {
		resData.Ok = false
		resData.Err = data.(string)
		resData.Data = nil
	}

	r, marshalErr := json.Marshal(resData)
	if marshalErr != nil {
		handleErrResponse(w, marshalErr)
		return
	}

	// send the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(r)
}

func handleResponse(w http.ResponseWriter, data interface{}) {
	handleResponseRaw(w, http.StatusOK, data)
}

func handleErrResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	if coder, ok := err.(serverErrorCoded); ok {
		code = coder.StatusCode()
	}

	msg := http.StatusText(code)

	// we want more data if we are in debug mode
	if GetConfig().Debug {
		msg = err.Error()
	}

	handleResponseRaw(w, code, msg)
}
