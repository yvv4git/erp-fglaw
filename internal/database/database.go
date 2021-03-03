package database

import (
	"sync"

	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var mu sync.Mutex
var dbInstance *gorm.DB

// GetInstance is used for create or get instance.
// This is singleton.
func GetInstance(config config.Config) (*gorm.DB, error) {
	var err error

	if dbInstance == nil {
		mu.Lock()
		defer mu.Unlock()

		if dbInstance == nil {
			dbInstance, err = gorm.Open(sqlite.Open(config.DB.FileName), &gorm.Config{})
			if err != nil {
				return nil, err
			}
		}
	}

	return dbInstance, err
}
