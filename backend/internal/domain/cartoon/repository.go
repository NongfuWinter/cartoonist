package cartoon

import "context"

// Repository 仓储端口。
type Repository interface {
	List(ctx context.Context) ([]Cartoon, error)
}
