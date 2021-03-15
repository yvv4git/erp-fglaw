package handlers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/tests"
)

const (
	configFile = "config/tests"
)

var (
	cfg       *config.Config
	webServer *fiber.App
	fixtures  *testfixtures.Loader
)

func TestMain(m *testing.M) {
	cfg = tests.Config()
	webServer = tests.WebServer()
	fixtures = tests.Fixtures()

	// Run tests.
	exitVal := m.Run()

	// Step after run tests.
	// ...
	os.Exit(exitVal)
}

func TestDevConfig(t *testing.T) {
	assert.Equal(t, "db/test.db", cfg.DB.FileName, "Check testing file name.")
	assert.Equal(t, "localhost", cfg.DB.Host, "Check db server host.")
	assert.Equal(t, int32(3306), cfg.DB.Port, "Check db server port.")
	assert.Equal(t, "localhost", cfg.WebSrv.Host, "Check web server host.")
	assert.Equal(t, int32(3005), cfg.WebSrv.Port, "Check web server port.")
}

func TestIndexHandler_Get(t *testing.T) {
	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
		{
			name: "Go to main page",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/", nil)
				return
			},
			expectStatusCode: 200,
		},
		{
			name: "Page not found",
			request: func() (req *http.Request, err error) {
				req, err = http.NewRequest("GET", "/unknown-page", nil)
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
