package entity

import "time"

type PasswordReset struct {
	ID        uint64
	UserID    uint64
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
}
