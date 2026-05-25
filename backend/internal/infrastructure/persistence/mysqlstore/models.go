package mysqlstore

import (
	"time"
)

// UserModel / CartoonModel 是持久化模型，仅基础设施层使用。

func (UserModel) TableName() string {
	return "users"
}

type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

func (CartoonModel) TableName() string {
	return "cartoons"
}

type CartoonModel struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
