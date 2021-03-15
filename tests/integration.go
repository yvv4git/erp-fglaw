package tests

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/database"
	"github.com/yvv4git/erp-fglaw/internal/handlers"
	"gorm.io/gorm"
)

const (
	configFile = "config/tests"
)

var (
	cfg       *config.Config
	db        *gorm.DB
	fixtures  *testfixtures.Loader
	webServer *fiber.App
)

func init() {
	// Change pwd.
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
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

	db, err = database.GetInstance(*cfg)
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
}

// PrepareTestDatabase is used for reload data storage state.
func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}

// Config is used as getter for config.
func Config() *config.Config {
	return cfg
}

// DB is used as getter for db.
func DB() *gorm.DB {
	return db
}

// Fixtures is used as getter for fixtures.
func Fixtures() *testfixtures.Loader {
	return fixtures
}

// WebServer is used as getter for webServer.
func WebServer() *fiber.App {
	return webServer
}
