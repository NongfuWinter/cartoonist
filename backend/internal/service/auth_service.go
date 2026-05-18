package service

import (
	"backend/internal/repository"
	"backend/pkg/utils"
)

func Login(username, password string) (bool, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	return utils.CheckPassword(password, user.Password), nil
}
