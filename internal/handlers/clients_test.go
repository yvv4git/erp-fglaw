package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/forms"
)

func TestClientsHandler(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
		{
			name: "Go to clients page",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/clients", nil)
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Page not found",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/clients/unknown-page", nil)
				return
			},
			expectStatusCode: 404,
		},
		{
			name: "Find(read) clients",
			request: func() (req *http.Request, err error) {
				form := forms.Clients{
					ID:           1,
					Number:       1,
					Address:      "San Fancisco Main street 88",
					CuitCustomer: "Some customer",
					ClientPhone:  "123456789012",
					ClientTypeID: 1,
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("GET", "/clients/read", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Create clients",
			request: func() (req *http.Request, err error) {
				form := forms.Clients{
					Address:      "San Fancisco Main street 88",
					CuitCustomer: "Some customer",
					ClientPhone:  "123456789012",
					ClientTypeID: 1,
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/clients/create", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Update clients",
			request: func() (req *http.Request, err error) {
				form := forms.Clients{
					ID:           1,
					Number:       2,
					Address:      "San Fancisco Main street 88",
					CuitCustomer: "Some customer",
					ClientPhone:  "123456789012",
					ClientTypeID: 1,
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("PUT", "/clients/update", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Delete clients",
			request: func() (req *http.Request, err error) {
				form := forms.Clients{
					ID: 1,
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("DELETE", "/clients/delete", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request, err := tc.request()
			if err != nil {
				t.Fatal(err)
			}
			result, err := webServer.Test(request)
			if err != nil {
				t.Fatal(err)
			}
			defer result.Body.Close()

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}
		})
	}
}
