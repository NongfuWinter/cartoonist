package mysqlstore

import (
	"context"

	"backend/internal/domain/user"

	"gorm.io/gorm"
)

var _ user.Repository = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var m UserModel
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}
	return &user.User{
		ID:           m.ID,
		Username:     m.Username,
		PasswordHash: m.Password,
	}, nil
}
