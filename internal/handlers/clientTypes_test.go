package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"github.com/yvv4git/erp-fglaw/tests"
)

func TestClientTypesHandler(t *testing.T) {
	// Anonymous struct.
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		check            func(body []byte)
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
			name: "Read client-types on first page",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					Pagination: forms.Pagination{
						Page:  0,
						Limit: 10,
					},
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("GET", "/client-types/read", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
			check: func(body []byte) {
				var result []domain.ClientTypes
				err := json.Unmarshal(body, &result)
				//t.Log(result)
				assert.Nil(t, err, "Unmarshal http body to client-types entities")
				assert.Equal(t, 10, len(result), "Count of client-types entities on first page")
			},
		},
		{
			name: "Read client-types on second page",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					Pagination: forms.Pagination{
						Page:  1,
						Limit: 10,
					},
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("GET", "/client-types/read", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
			check: func(body []byte) {
				var result []domain.ClientTypes
				err := json.Unmarshal(body, &result)
				assert.Nil(t, err, "Unmarshal http body to client-types entities")
				assert.Equal(t, 1, len(result), "Count of client-types entities on second page")
			},
		},
		{
			name: "Read client-types on third page, but empty result",
			request: func() (req *http.Request, err error) {
				form := forms.ClientTypes{
					Pagination: forms.Pagination{
						Page:  2,
						Limit: 10,
					},
				}
				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("GET", "/client-types/read", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
			check: func(body []byte) {
				var result []domain.ClientTypes
				err := json.Unmarshal(body, &result)
				assert.Nil(t, err, "Unmarshal http body to client-types entities")
				assert.Equal(t, 0, len(result), "Count of client-types entities on third page")
			},
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

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			if tc.check != nil {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				tc.check(body)
			}
		})
	}
}

func TestClientTypes_Create(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		check            func(body []byte)
		description      string
	}{
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
			check: func(body []byte) {
				//t.Log(string(body))
				assert.Equal(t, `{"success":true}`, string(body), "Status success")
			},
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

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			if tc.check != nil {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				tc.check(body)
			}
		})
	}
}

func TestClientTypes_Update(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		check            func(body []byte)
		description      string
	}{
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
			check: func(body []byte) {
				//t.Log(string(body))
				assert.Equal(t, `{"success":true}`, string(body), "Status success")
			},
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

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			if tc.check != nil {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				tc.check(body)
			}
		})
	}
}

func TestClientTypes_Delete(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		check            func(body []byte)
		description      string
	}{
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
			check: func(body []byte) {
				//t.Log(string(body))
				assert.Equal(t, `{"success":true}`, string(body), "Status success")
			},
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

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			if tc.check != nil {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				tc.check(body)
			}
		})
	}
}
