package response

import "time"

type UserPermissionResponse struct {
	UserID       uint64    `json:"user_id"`
	PermissionID uint64    `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserPermissionDetailResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
}

type GetUserPermissionsResponse struct {
	UserID      uint64                         `json:"user_id"`
	Permissions []UserPermissionDetailResponse `json:"permissions"`
}
