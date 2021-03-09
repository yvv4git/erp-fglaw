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

func TestClientTypesHandler(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
		{
			name: "Go to client-types page",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/client-types", nil)
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Page not found",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/client-types/unknown-page", nil)
				return
			},
			expectStatusCode: 404,
		},
		{
			name: "Find(read) client-types",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					ID:         1,
					ClientType: "Some",
					ActingAs:   "blah",
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("GET", "/client-types/read", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Create client-types",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					ClientType: "Some",
					ActingAs:   "blah",
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/client-types/create", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Update client-types",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					ID:         1,
					ClientType: "Some2",
					ActingAs:   "blah2",
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("PUT", "/client-types/update", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Delete client-types",
			request: func() (req *http.Request, err error) {
				form := forms.Clients{
					ID: 1,
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("DELETE", "/client-types/delete", bytes.NewBuffer(jsonForm))
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
