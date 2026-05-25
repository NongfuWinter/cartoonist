package db

import (
	"backend/internal/infrastructure/persistence/mysqlstore"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open 连接 MySQL 并完成 AutoMigrate。
func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := mysqlstore.AutoMigrate(db); err != nil {
		return nil, err
	}
	return db, nil
}
