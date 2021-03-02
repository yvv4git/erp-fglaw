package handlers

import (
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

const (
	configFile = "../../config/main"
)

var (
	cfg       *config.Config
	webServer *fiber.App
)

func TestMain(m *testing.M) {
	// Pre-setup.
	// Change pwd.
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	// Step before run tests.
	// Init config.
	cfg, err := config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Init routes.
	webServer = SetupWebServer(*cfg)

	// Run tests.
	exitVal := m.Run()

	// Step after run tests.
	// ...
	os.Exit(exitVal)
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
