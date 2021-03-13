package handlers_test

import (
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/database"
	"github.com/yvv4git/erp-fglaw/internal/handlers"
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
	cfg, err = config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.GetInstance(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Init routes.
	webServer = handlers.SetupWebServer(*cfg, db)

	instanceDB, err := db.DB()
	fixtures, err = testfixtures.New(
		testfixtures.Database(instanceDB),
		testfixtures.Dialect("sqlite"),           // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("tests/fixtures"), // the directory containing the YAML files
	)
	if err != nil {
		panic(err)
	}

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
