package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"github.com/yvv4git/erp-fglaw/tests"
)

func TestClientTypesHandler(t *testing.T) {
	// Anonymous struct.
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
			name: "Read client-types on first page",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "0")
				data.Set("limit", "10")
				url := "/client-types/read?" + data.Encode()

				req, err = http.NewRequest("GET", url, nil)
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Read client-types on second page",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "1")
				data.Set("limit", "10")

				url := "/client-types/read?" + data.Encode()
				req, err = http.NewRequest("GET", url, nil)
				// req.Header.Set("Content-Type", "application/json")
				// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Read client-types on third page, but empty result",
			request: func() (req *http.Request, err error) {
				data := url.Values{}
				data.Set("page", "2")
				data.Set("limit", "10")

				url := "/client-types/read?" + data.Encode()
				req, err = http.NewRequest("GET", url, nil)
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

			if result.StatusCode == 500 {
				body, err := ioutil.ReadAll(result.Body)
				assert.Nil(t, err)
				t.Log(string(body))
			}

			assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)

			/* body, err := ioutil.ReadAll(result.Body)
			assert.Nil(t, err)
			t.Log(body) */
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
