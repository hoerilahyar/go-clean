package entity

type UserPermission struct {
	UserID       uint64 `db:"user_id"`
	PermissionID uint64 `db:"permission_id"`
}
