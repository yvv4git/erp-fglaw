package database

import (
	"log"
	"os"
	"path"
	"testing"

	"runtime"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

const (
	configFile = "../../config/main"
)

var (
	cfg *config.Config
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

	// Run tests.
	exitVal := m.Run()

	// Step after run tests.
	CloseDB()

	os.Exit(exitVal)
}

func TestSingletone(t *testing.T) {
	singletoneFirst, err := GetInstance(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	singletoneSecond, err := GetInstance(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, singletoneFirst, singletoneSecond, "DB instances non equal.")
}
