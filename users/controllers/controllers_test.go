package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Test create controller
func TestCreate(t *testing.T) {
	tt := []struct {
		name    string
		json    []byte
		status  int
		message string
		err     string
	}{
		{
			name: "Create with correct parameters",
			json: []byte(`{
					"name": "Ramiro",
					"email": "Ramiro@ramiro.com"
				}`),
			status:  http.StatusOK,
			message: "User created successfully",
			err:     "User should have been created",
		},
		{
			name:    "Create without parameters",
			json:    []byte(``),
			status:  http.StatusBadRequest,
			message: "",
			err:     "",
		},
		{
			name: "Create with invalid name",
			json: []byte(`{
				"name": "ra",
				"email": "Ramiro@ramiro.com"
			}`),
			status:  http.StatusBadRequest,
			message: "",
			err:     "",
		},
		{
			name: "Create with invalid email",
			json: []byte(`{
				"name": "Ramiro",
				"email": "Ramiroramiro.com"
			}`),
			status:  http.StatusBadRequest,
			message: "",
			err:     "",
		},
	}

	// 	var json = []byte(`{
	//     "name": "Ramiro",
	//     "email": "Ramiro@ramiro.com"
	// }`)

	// Run all sub tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := "/api/users/create"

			r := chi.NewMux()
			r.Post(url, Create)

			request, err := http.NewRequest("POST", url, bytes.NewBuffer(tc.json))
			if err != nil {
				t.Errorf("Expected nil, received %s", err.Error())
			}

			response := httptest.NewRecorder()

			r.ServeHTTP(response, request)

			// defer response.Body.Close()

			if response.Code != tc.status {
				t.Errorf("Expected %d, received %d", tc.status, response.Code)
			}

			// Obtain the response body as a string
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			bodyString := string(bodyBytes)
			// fmt.Println(bodyString)

			if !strings.Contains(bodyString, tc.message) {
				t.Errorf(tc.err)
			}
		})
	}
}

// Test read all controllers
func TestReadAll(t *testing.T) {
	url := "/api/users/readall"
	message := "users where fetched"

	r := chi.NewMux()
	r.Get(url, ReadAll)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	// defer response.Body.Close()

	if response.Code != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, response.Code)
	}

	// Obtain the response body as a string
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	if !strings.Contains(bodyString, message) {
		t.Errorf("Wrong response body")
	}
}

func TestFilter(t *testing.T) {
	tt := []struct {
		name    string
		json    []byte
		status  int
		message string
		err     string
	}{
		{
			name:    "Filter without parameters",
			json:    []byte(``),
			status:  http.StatusOK,
			message: "users where fetched",
			err:     "It should have fetch any user",
		},
		{
			name: "Filter with empty parameters",
			json: []byte(`{
					"name": "",
					"email": ""
				}`),
			status:  http.StatusOK,
			message: "users where fetched",
			err:     "It should have fetch any user",
		},
		{
			name: "Filter without name",
			json: []byte(`{
					"name": "",
					"email": "Ramiro"
				}`),
			status:  http.StatusOK,
			message: "users where fetched",
			err:     "It should have fetch any user",
		},
		{
			name: "Filter without email",
			json: []byte(`{
					"name": "Ramiro",
					"email": ""
				}`),
			status:  http.StatusOK,
			message: "users where fetched",
			err:     "It should have fetch any user",
		},
		{
			name: "Filter with both parameters",
			json: []byte(`{
					"name": "Ramiro",
					"email": "@ramiro.com"
				}`),
			status:  http.StatusOK,
			message: "users where fetched",
			err:     "It should have fetch any user",
		},
	}

	// 	var json = []byte(`{
	//     "name": "Ramiro",
	//     "email": "Ramiro@ramiro.com"
	// }`)

	// Run all sub tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := "/api/users/filter"

			r := chi.NewMux()
			r.Get(url, Filter)

			request, err := http.NewRequest("GET", url, bytes.NewBuffer(tc.json))
			if err != nil {
				t.Errorf("Expected nil, received %s", err.Error())
			}

			response := httptest.NewRecorder()

			r.ServeHTTP(response, request)

			// defer response.Body.Close()

			if response.Code != tc.status {
				t.Errorf("Expected %d, received %d", tc.status, response.Code)
			}

			// Obtain the response body as a string
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			bodyString := string(bodyBytes)
			// fmt.Println(bodyString)

			if !strings.Contains(bodyString, tc.message) {
				t.Errorf(tc.err)
			}
		})
	}
}
