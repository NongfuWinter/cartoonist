package cartoon

import "time"

// Cartoon 领域实体。
type Cartoon struct {
	ID        uint
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
