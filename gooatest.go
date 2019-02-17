package gooatest

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
)

// Validator impliments validate methods
// to validate HTTP Response & Request While against Open API Specification.
type Validator interface {
	ValidateRequest() error
	ValidateResponse() error
}

// OpenAPIValidator validates HTTP Response & Request While against Open API Specification.
type OpenAPIValidator struct {
	Context                 context.Context
	RequestValidationInput  *openapi3filter.RequestValidationInput
	ResponseValidationInput *openapi3filter.ResponseValidationInput
}

// Params is validator params
type Params struct {
	HTTPReq         *http.Request
	URI             string
	SchemaPath      string
	Context         context.Context
	ResponseRecoder *httptest.ResponseRecorder
}

// NewRouterFromYAML generates new router from YAML.
func newRouterFromYAML(path string) (*openapi3filter.Router, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromYAMLData(data)
	if err != nil {
		return nil, err
	}
	router := openapi3filter.NewRouter()
	router.AddSwagger(swagger)
	if err != nil {
		return nil, err
	}
	return router, nil
}

// NewValidator generates new Validator.
func NewValidator(p Params) (Validator, error) {
	u, err := url.Parse(p.URI)
	if err != nil {
		return nil, err
	}
	router, err := newRouterFromYAML(p.SchemaPath)
	if err != nil {
		return nil, err
	}
	route, _, err := router.FindRoute(p.HTTPReq.Method, u)
	if err != nil {
		return nil, err
	}
	// Validate response
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request: p.HTTPReq,
		Route:   route,
	}
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 p.ResponseRecoder.Code,
		Header:                 p.ResponseRecoder.HeaderMap,
	}
	if p.ResponseRecoder.Body != nil {
		body, err := ioutil.ReadAll(p.ResponseRecoder.Body)
		if err != nil {
			return nil, err
		}
		responseValidationInput.SetBodyBytes(body)
	}
	return OpenAPIValidator{
		Context:                 p.Context,
		RequestValidationInput:  requestValidationInput,
		ResponseValidationInput: responseValidationInput,
	}, nil
}

// ValidateRequest validates HTTP Request While against Open API Specification.
func (v OpenAPIValidator) ValidateRequest() error {
	if err := openapi3filter.ValidateRequest(v.Context, v.RequestValidationInput); err != nil {
		return err
	}
	return nil
}

// ValidateResponse validates HTTP Response & Request While against Open API Specification.
func (v OpenAPIValidator) ValidateResponse() error {
	if err := openapi3filter.ValidateResponse(v.Context, v.ResponseValidationInput); err != nil {
		return err
	}
	return nil
}
