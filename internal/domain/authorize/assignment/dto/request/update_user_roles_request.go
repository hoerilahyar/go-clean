package request

type UpdateUserRolesRequest struct {
	RoleIDs []uint64 `json:"role_ids" binding:"required"`
}
