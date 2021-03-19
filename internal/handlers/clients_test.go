package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"github.com/yvv4git/erp-fglaw/tests"
)

func TestClientsHandler(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
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
			tests.PrepareTestDatabase()

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

func TestClientsHandler_Read(t *testing.T) {
	type resultRead struct {
		Clients []domain.Clients `json:"clients"`
		Count   int              `json:"count"`
	}

	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		check            func(body io.Reader)
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
			name: "Read clients on first page",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "0")
				data.Set("limit", "10")
				url := "/clients/read?" + data.Encode()

				req, err = http.NewRequest("GET", url, nil)
				return
			},
			check: func(body io.Reader) {
				var resRead resultRead
				bodyBytes, err := ioutil.ReadAll(body)
				assert.Nil(t, err)
				err = json.Unmarshal(bodyBytes, &resRead)
				assert.Nil(t, err)
				assert.Equal(t, 10, len(resRead.Clients), "Check clients count")
				assert.Equal(t, 11, resRead.Count, "Check clients count without pagination")
			},
			expectStatusCode: 200,
		},
		{
			name: "Read clients on second page",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "1")
				data.Set("limit", "10")
				url := "/clients/read?" + data.Encode()

				req, err = http.NewRequest("GET", url, nil)
				return
			},
			check: func(body io.Reader) {
				var resRead resultRead
				bodyBytes, err := ioutil.ReadAll(body)
				assert.Nil(t, err)
				err = json.Unmarshal(bodyBytes, &resRead)
				assert.Nil(t, err)
				assert.Equal(t, 1, len(resRead.Clients), "Check clients count")
				assert.Equal(t, 11, resRead.Count, "Check clients count without pagination")
			},
			expectStatusCode: 200,
		},
		{
			name: "Filter clients by ID",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "0")
				data.Set("limit", "10")
				data.Set("id", "1")
				url := "/clients/read?" + data.Encode()

				req, err = http.NewRequest("GET", url, nil)
				return
			},
			check: func(body io.Reader) {
				var resRead resultRead
				bodyBytes, err := ioutil.ReadAll(body)
				assert.Nil(t, err)
				err = json.Unmarshal(bodyBytes, &resRead)
				assert.Nil(t, err)
				assert.Equal(t, 1, len(resRead.Clients), "Check clients count")
				assert.Equal(t, 1, resRead.Count, "Check clients count without pagination")
			},
			expectStatusCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase()

			// Create request.
			request, err := tc.request()
			if err != nil {
				t.Fatal(err)
			}

			// Start test web server.
			result, err := webServer.Test(request)
			if err != nil {
				t.Fatal(err)
			}
			defer result.Body.Close()

			// Check.
			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)
			if tc.check != nil {
				tc.check(result.Body)
			}
		})
	}
}
