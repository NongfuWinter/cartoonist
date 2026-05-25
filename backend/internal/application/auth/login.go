package auth

import (
	"context"

	"backend/internal/domain/user"
	"backend/pkg/utils"
)

// Login 用户登录用例。
type Login struct {
	users user.Repository
}

func NewLogin(users user.Repository) *Login {
	return &Login{users: users}
}

func (uc *Login) Execute(ctx context.Context, username, password string) (bool, error) {
	u, err := uc.users.FindByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	return utils.CheckPassword(password, u.PasswordHash), nil
}
