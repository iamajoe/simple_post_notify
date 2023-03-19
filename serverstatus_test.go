package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_reqStatus(t *testing.T) {
	type args struct {
	}
	type testStruct struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}

	tests := []testStruct{
		func() testStruct {
			return testStruct{
				name:     "runs",
				args:     args{},
				wantCode: http.StatusOK,
				wantBody: "{\"ok\":true,\"code\":200,\"data\":true}",
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/status", nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(reqStatus())
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("wrong status code: got %v want %v", rec.Code, tt.wantCode)
			}

			if rec.Body.String() != tt.wantBody {
				t.Errorf("body = %v, want %v", rec.Body.String(), tt.wantBody)
			}
		})
	}
}

func Test_getStatusEndpoints(t *testing.T) {
	type testStruct struct {
		name     string
		endpoint string
		method   string
		body     *bytes.Reader
		wantCode int
		wantBody string
	}

	tests := []testStruct{
		func() testStruct {
			return testStruct{
				name:     "/status : runs",
				endpoint: "/status",
				method:   "GET",
				body:     nil,
				wantCode: http.StatusOK,
				wantBody: "{\"ok\":true,\"code\":200,\"data\":true}",
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := getRouter()
			resCode, resBody, err := testRequest(tt.method, tt.endpoint, tt.body, router)
			if err != nil {
				t.Fatal(err)
				return
			}

			// test the response
			if resCode != tt.wantCode {
				t.Fatalf("wrong status code: got %v want %v", resCode, tt.wantCode)
			}

			if string(resBody) != tt.wantBody {
				t.Errorf("body = %v, want %v", string(resBody), tt.wantBody)
			}
		})
	}
}
