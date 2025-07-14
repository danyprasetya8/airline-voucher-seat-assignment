package database

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(filepath.Join("data", "vouchers.db")), &gorm.Config{})
}
