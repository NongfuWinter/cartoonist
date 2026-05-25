package cartoonuc

import (
	"context"

	dom "backend/internal/domain/cartoon"
)

// ListCartoons 列出漫画用例。
type ListCartoons struct {
	repo dom.Repository
}

func NewListCartoons(repo dom.Repository) *ListCartoons {
	return &ListCartoons{repo: repo}
}

func (uc *ListCartoons) Execute(ctx context.Context) ([]dom.Cartoon, error) {
	return uc.repo.List(ctx)
}
