package entity

import "time"

type User struct {
	ID              uint64     `json:"id"`
	FullName        string     `json:"full_name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	PhoneNumber     *string    `json:"phone_number"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-"`
	RememberToken   *string    `json:"-"`
	Avatar          *string    `json:"avatar"`
	Status          string     `json:"status"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	CreatedBy       *uint64    `json:"created_by"`
	UpdatedBy       *uint64    `json:"updated_by"`
	DeletedBy       *uint64    `json:"deleted_by"`
}
