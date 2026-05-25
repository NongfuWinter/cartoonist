package user

import "context"

// Repository 为用户领域仓储端口，由基础设施层实现。
type Repository interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
}
