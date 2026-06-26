package response

type MeResponse struct {
	User        UserResponse         `json:"user"`
	Roles       []RoleResponse       `json:"roles"`
	Permissions []PermissionResponse `json:"permissions"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type RoleResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PermissionResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Group string `json:"group"`
}
