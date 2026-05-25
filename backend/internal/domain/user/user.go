package user

// User 领域实体（不依赖 ORM / Web）。
type User struct {
	ID           uint
	Username     string
	PasswordHash string
}
