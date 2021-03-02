package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsHandler_Get(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
		{
			name: "Go to products page",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/products", nil)
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Page not found",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/products/unknown-page", nil)
				return
			},
			expectStatusCode: 404,
		},
	}

	for _, tc := range testCases {
		request, err := tc.request()
		if err != nil {
			t.Fatal(err)
		}
		result, err := webServer.Test(request)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tc.expectStatusCode, result.StatusCode, tc.description)
	}
}
