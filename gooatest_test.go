package gooatest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/po3rin/gooatest"
)

var (
	schemaPath        = "_test/schema/schema.yml"
	commonContext     = context.Background()
	commonContentType = "application/json"
)

type user struct {
	Name string `json:"name"`
}

type requestBody struct {
	User user `json:"user"`
}

func TestGetValidateResponse(t *testing.T) {
	httpReq, _ := http.NewRequest(http.MethodGet, "/users", nil)
	tests := []struct {
		name            string
		httpReq         *http.Request
		responseRecoder *httptest.ResponseRecorder
		expectedError   bool
	}{
		{
			name:    "valid_response",
			httpReq: httpReq,
			responseRecoder: &httptest.ResponseRecorder{
				Code: 200,
				HeaderMap: http.Header{
					"Content-Type": []string{
						commonContentType,
					},
				},
				Body: bytes.NewBufferString(`{"users":[{"id":1,"name":"po3rin","added_at":"2018-12-01T00:00:00Z"}]}`),
			},
			expectedError: false,
		},
		{
			name:    "invalid_response",
			httpReq: httpReq,
			responseRecoder: &httptest.ResponseRecorder{
				Code: 200,
				HeaderMap: http.Header{
					"Content-Type": []string{
						commonContentType,
					},
				},
				Body: bytes.NewBufferString(`{"users":[{"id":1}]}`),
			},
			expectedError: true,
		},
		{
			name:    "no_headermap",
			httpReq: httpReq,
			responseRecoder: &httptest.ResponseRecorder{
				Code: 200,
				Body: bytes.NewBufferString(`{"users":[{"id":1,"name":"po3rin","added_at":"2018-12-01T00:00:00Z"}]}`),
			},
			expectedError: true,
		},
		{
			name:    "no_code",
			httpReq: httpReq,
			responseRecoder: &httptest.ResponseRecorder{
				HeaderMap: http.Header{
					"Content-Type": []string{
						commonContentType,
					},
				},
				Body: bytes.NewBufferString(`{"users":[{"id":1,"name":"po3rin","added_at":"2018-12-01T00:00:00Z"}]}`),
			},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		p := gooatest.Params{
			HTTPReq:         tt.httpReq,
			BaseURL:         "http://localhost:8080",
			SchemaPath:      schemaPath,
			Context:         commonContext,
			ResponseRecoder: tt.responseRecoder,
		}
		v, err := gooatest.NewValidator(p)
		if err != nil {
			t.Errorf("test %+v detects unexpected error. %v\n", tt.name, err)
		}
		err = v.ValidateResponse()
		if err != nil && !tt.expectedError {
			t.Errorf("test %+v detects unexpected error in ValidateResponse. %v\n", tt.name, err)
		} else if err == nil && tt.expectedError {
			t.Errorf("test %+v expected error in ValidateResponse.", tt.name)
		}
	}
}
func TestPostValidateResponse(t *testing.T) {
	tests := []struct {
		name                       string
		ReqBody                    interface{}
		responseRecoder            *httptest.ResponseRecorder
		expectedErrorInReqValidate bool
		expectedErrorInResValidate bool
	}{
		{
			name: "valid_response",
			ReqBody: requestBody{
				User: user{
					Name: "po3rin",
				},
			},
			responseRecoder: &httptest.ResponseRecorder{
				Code: 200,
				HeaderMap: http.Header{
					"Content-Type": []string{
						commonContentType,
					},
				},
				Body: bytes.NewBufferString(`{"users":[{"id":1,"name":"po3rin","added_at":"2018-12-01T00:00:00Z"}]}`),
			},
			expectedErrorInReqValidate: false,
			expectedErrorInResValidate: false,
		},
	}
	for _, tt := range tests {
		d, err := json.Marshal(tt.ReqBody)
		if err != nil {
			log.Fatalf("json.Marshal() failed with '%s'\n", err)
		}
		reqBody := bytes.NewBuffer(d)
		fmt.Println(reqBody)
		httpReq, err := http.NewRequest(http.MethodGet, "/users", reqBody)
		if err != nil {
			t.Errorf("test %+v detects unexpected error. %v\n", tt.name, err)
		}
		p := gooatest.Params{
			HTTPReq:         httpReq,
			BaseURL:         "http://localhost:8080",
			SchemaPath:      schemaPath,
			Context:         commonContext,
			ResponseRecoder: tt.responseRecoder,
		}
		v, err := gooatest.NewValidator(p)
		if err != nil {
			t.Errorf("test %+v detects unexpected error. %v\n", tt.name, err)
		}
		err = v.ValidateRequest()
		if err != nil && !tt.expectedErrorInReqValidate {
			t.Errorf("test %+v detects unexpected error in ValidateRequest. %v\n", tt.name, err)
		} else if err == nil && tt.expectedErrorInReqValidate {
			t.Errorf("test %+v expected error in ValidateRequest.", tt.name)
		}
		err = v.ValidateResponse()
		if err != nil && !tt.expectedErrorInResValidate {
			t.Errorf("test %+v detects unexpected error in ValidateResponse. %v\n", tt.name, err)
		} else if err == nil && tt.expectedErrorInResValidate {
			t.Errorf("test %+v expected error in ValidateResponse.", tt.name)
		}
	}
}
