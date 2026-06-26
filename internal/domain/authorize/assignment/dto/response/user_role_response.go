package response

import "time"

type UserRoleResponse struct {
	UserID    uint64    `json:"user_id"`
	RoleID    uint64    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRoleDetailResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type GetUserRolesResponse struct {
	UserID uint64                   `json:"user_id"`
	Roles  []UserRoleDetailResponse `json:"roles"`
}
