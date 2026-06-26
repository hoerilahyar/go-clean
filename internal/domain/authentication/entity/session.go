package entity

import "time"

type Session struct {
	ID           uint64
	UserID       uint64
	RefreshToken string
	UserAgent    string
	IPAddress    string
	ExpiredAt    time.Time
	CreatedAt    time.Time
}
