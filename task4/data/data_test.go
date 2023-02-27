package data

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestsOutput struct {
	name     string
	server   *httptest.Server
	response Output
	expError error
}

type TestsInput struct {
	name     string
	server   *httptest.Server
	response Input
	expError error
}

func TestGetTelephones(t *testing.T) {
	input := &Input{
		Id:        1,
		FirstName: "Ognjen",
		LastName:  "Bogdanovic",
		Telephone: "0638314856",
	}
	var output []Input
	output = append(output, *input)
	tests := []TestsOutput{
		{
			name: "request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`"results": [
					{
						"id": 1,
						"firstname": "Ognjen",
						"lastname": "Bogdanovic",
						"telephone": "0638314856"
					},`))
			})),
			response: Output{
				Results: output,
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := GetTelephones("test.json")
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expError, err)
			}
		})
	}
}

func TestGetTelephonesById(t *testing.T) {
	input := Input{
		Id:        1,
		FirstName: "Ognjen",
		LastName:  "Bogdanovic",
		Telephone: "0638314856",
	}
	tests := []TestsInput{
		{
			name: "request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`
						"id": 1,
						"firstname": "Ognjen",
						"lastname": "Bogdanovic",
						"telephone": "0638314856"
					`))
			})),
			response: input,
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := GetTelephoneById(1, "test.json")
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expError, err)
			}
		})
	}
}

func TestDeleteTelephoneByID(t *testing.T) {
	tests := []TestsOutput{
		{
			name: "request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`"results": nil`))
			})),
			response: Output{
				Results: []Input{},
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := DeleteTelephone(1, "test.json")
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expError, err)
			}
		})
	}
}

func TestUploadTelephone(t *testing.T) {
	input := &Input{
		Id:        1,
		FirstName: "Ognjen",
		LastName:  "Bogdanovic",
		Telephone: "0638314856",
	}
	var output []Input
	output = append(output, *input)
	tests := []TestsOutput{
		{
			name: "request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`"results": [
					{
						"id": 1,
						"firstname": "Ognjen",
						"lastname": "Bogdanovic",
						"telephone": "0638314856"
					},`))
			})),
			response: Output{
				Results: output,
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := UploadTelephone(*input, "test.json")
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expError, err)
			}
		})
	}
}
