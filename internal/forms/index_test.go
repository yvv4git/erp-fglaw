package forms_test

import (
	"log"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/database"
	"gorm.io/gorm"

	"github.com/go-testfixtures/testfixtures/v3"
)

const (
	configFile = "config/tests"
)

var (
	cfg      *config.Config
	db       *gorm.DB
	fixtures *testfixtures.Loader
)

// TestMain used for some logic before and after run tests.
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

	db, err = database.GetInstance(*cfg)
	if err != nil {
		log.Fatal(err)
	}

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
	database.CloseDB()

	os.Exit(exitVal)
}
