package response

import "time"

type RolePermissionResponse struct {
	RoleID       uint64    `json:"role_id"`
	PermissionID uint64    `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type RolePermissionDetailResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
}

type GetRolePermissionsResponse struct {
	RoleID      uint64                         `json:"role_id"`
	Permissions []RolePermissionDetailResponse `json:"permissions"`
}
