package entity

type UserRole struct {
	UserID uint64 `db:"user_id"`
	RoleID uint64 `db:"role_id"`
}
