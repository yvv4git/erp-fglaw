package forms_test

import (
	"os"
	"testing"

	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/database"
	"github.com/yvv4git/erp-fglaw/tests"
	"gorm.io/gorm"

	"github.com/go-testfixtures/testfixtures/v3"

	_ "github.com/yvv4git/erp-fglaw/tests"
)

var (
	cfg      *config.Config
	db       *gorm.DB
	fixtures *testfixtures.Loader
)

// TestMain used for some logic before and after run tests.
func TestMain(m *testing.M) {
	cfg = tests.Config()
	db = tests.DB()
	fixtures = tests.Fixtures()

	// Run tests.
	exitVal := m.Run()

	// Step after run tests.
	database.CloseDB()

	os.Exit(exitVal)
}
