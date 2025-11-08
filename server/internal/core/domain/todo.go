package domain

import "time"

type Todo struct {
	ID          uint
	Title       string
	Description string
	IsDone      bool
	CompletedAt *time.Time // nil if not completed
	CreatedAt   time.Time
}
