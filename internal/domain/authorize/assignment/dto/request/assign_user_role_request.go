package request

type AssignUserRoleRequest struct {
	RoleIDs []uint64 `json:"role_ids" binding:"required,min=1"`
}
