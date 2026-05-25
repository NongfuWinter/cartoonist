package mysqlstore

import (
	"context"

	"backend/internal/domain/cartoon"

	"gorm.io/gorm"
)

type CartoonRepository struct {
	db *gorm.DB
}

func NewCartoonRepository(db *gorm.DB) *CartoonRepository {
	return &CartoonRepository{db: db}
}

func (r *CartoonRepository) List(ctx context.Context) ([]cartoon.Cartoon, error) {
	var rows []CartoonModel
	if err := r.db.WithContext(ctx).Order("id DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]cartoon.Cartoon, len(rows))
	for i, m := range rows {
		out[i] = cartoon.Cartoon{
			ID:        m.ID,
			Title:     m.Title,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	}
	return out, nil
}
