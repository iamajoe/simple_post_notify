package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_reqNotify(t *testing.T) {
	type args struct {
		msg  string
		mod  string
		body []byte
	}
	type testStruct struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}

	tests := []testStruct{
		func() testStruct {
			msg := fmt.Sprintf("%d", rand.Intn(10000))
			mod := fmt.Sprintf("%d", rand.Intn(10000))
			body := []byte(fmt.Sprintf(`{"msg":"%s", "mod":"%s"}`, msg, mod))

			return testStruct{
				name:     "errors if module does not exist",
				args:     args{msg, mod, body},
				wantCode: http.StatusBadRequest,
				wantBody: fmt.Sprintf(
					"{\"ok\":false,\"code\":%d,\"err\":\"%s\"}",
					http.StatusBadRequest,
					http.StatusText(http.StatusBadRequest),
				),
			}
		}(),
		func() testStruct {
			msg := fmt.Sprintf("%d", rand.Intn(10000))
			mod := "telegram"
			body := []byte(fmt.Sprintf(`{"msg":"%s", "mod":"%s"}`, msg, mod))

			return testStruct{
				name:     "runs",
				args:     args{msg, mod, body},
				wantCode: http.StatusOK,
				wantBody: fmt.Sprintf(
					"{\"ok\":true,\"code\":%d,\"data\":\"%s\"}",
					http.StatusOK,
					"true",
				),
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/send", bytes.NewBuffer(tt.args.body))
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(reqNotifySend())
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("wrong status code: got %v want %v", rec.Code, tt.wantCode)
			}

			if rec.Body.String() != tt.wantBody {
				t.Errorf("body = %v, want %v", rec.Body.String(), tt.wantBody)
			}

			// TODO: need to actually check if sent
		})
	}
}
