package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes)
}

func CheckPassword(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
