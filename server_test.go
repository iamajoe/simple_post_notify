package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
)

func testRequest(
	method string,
	endpoint string,
	body *bytes.Reader,
	router *chi.Mux,
) (int, string, error) {
	httpserver := httptest.NewServer(router)
	defer httpserver.Close()

	// default the body so it doesnt error unnecessarily
	parsedBody := body
	if parsedBody == nil {
		parsedBody = bytes.NewReader([]byte(""))
	}

	// start request test
	req, err := http.NewRequest(method, httpserver.URL+endpoint, parsedBody)
	if err != nil {
		return -1, "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, "", err
	}
	defer res.Body.Close()

	return res.StatusCode, string(resBody), nil
}
