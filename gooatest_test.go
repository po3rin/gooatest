package gooatest_test

import (
	"context"
)

var (
	schemaPath    = "_test/schema/schema.yml"
	commonContext = context.Background()
)

// ToDo : to write test code.
// func TestValidateResponse(t *testing.T) {
//  httpReq, _ := http.NewRequest(http.MethodGet, "/users", nil)
// 	tests := []struct {
// 		name          string
// 		input         gooatest.Params
// 		expectedError bool
// 	}{
// 		{
// 			name: "",
// 			input: gooatest.Params{
// 				HTTPReq:    httpReq,
// 				URI:        "http://localhost:8080/users",
// 				SchemaPath: schemaPath,
// 				Context:    commonContext,
// 			},
// 			expectedError: false,
// 		},
// 		{
// 			name: "",
// 			input: gooatest.Params{
// 				URI:        "http://localhost:8080/users",
// 				SchemaPath: schemaPath,
// 				Context:    commonContext,
// 			},
// 			expectedError: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		v, err := gooatest.NewValidator(tt.input)
// 		if err != nil {
// 			t.Errorf("test %+v detects unexpected error. %v\n", tt.name, err)
// 		}
// 		err = v.ValidateResponse()
// 		if err != nil {
// 			t.Errorf("test %+v detects unexpected error. %v\n", tt.name, err)
// 		}
// 	}
// }
