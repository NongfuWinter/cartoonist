package mysqlstore

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&UserModel{}, &CartoonModel{})
}
