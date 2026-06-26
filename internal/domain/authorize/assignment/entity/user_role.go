package entity

import "time"

type UserRole struct {
	UserID uint64 `db:"user_id" json:"user_id"`
	RoleID uint64 `db:"role_id" json:"role_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	CreatedBy *uint64 `db:"created_by" json:"created_by"`
}
