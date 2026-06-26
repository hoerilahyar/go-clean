package entity

import "time"

type RolePermission struct {
	RoleID       uint64 `db:"role_id" json:"role_id"`
	PermissionID uint64 `db:"permission_id" json:"permission_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	CreatedBy *uint64 `db:"created_by" json:"created_by"`
}
