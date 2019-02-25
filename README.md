# gooatest

<img src="https://img.shields.io/badge/go-v1.11-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/gooatest.svg?style=shield)](https://circleci.com/gh/po3rin/gooatest)

## Overview
Package gooatest lets you validate HTTP Response & Request While against Open API Specification. This Package wraps getkin/kin-openapi package.

## Usage

```go
// ...

import (
    // ...
	"github.com/po3rin/gooatest"
)

func TestValidateResponse() {
    // prepare http.Request & http.Response.
    httpReq, _ := http.NewRequest(http.MethodGet, "/users", nil)
    responseRecoder:= &httptest.ResponseRecorder{
        Code: 200,
        HeaderMap: http.Header{
            "Content-Type": []string{"application/json"},
        },
        Body: bytes.NewBufferString(`{"users":[{"id":1,"name":"po3rin","added_at":"2018-12-01T00:00:00Z"}]}`),
    }
    // you able to get http.Response from httptest.ResponseRecorder using Result()
    httpRes := responseRecoder.Result()

    // prepare gooatest.Params to init validator.
    p := gooatest.Params{
        HTTPReq:         httpReq,
        HTTPRes:         httpRes,
        BaseURL:         "http://localhost:8080",
        SchemaPath:      "_test/schema/schema.yml",
        Context:         context.Background(),
    }

    // init validator using gooatest.Params.
    v, _ := gooatest.NewValidator(p)

    // exec to validate.
    err = v.ValidateResponse()
    if err != nil {
        // error handling ...
    }
}
```
