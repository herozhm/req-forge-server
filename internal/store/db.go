package store

import (
	"req-forge-server/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// InitDB 打开 SQLite 并自动建表(glebarez/sqlite 纯 Go 实现,无需 CGO)。
func InitDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Decomposition{}); err != nil {
		return nil, err
	}
	return db, nil
}
