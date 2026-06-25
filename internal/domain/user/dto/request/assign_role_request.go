package request

type AssignRoleRequest struct {
	RoleIDs []uint64 `json:"role_ids" binding:"required,min=1,dive,gt=0"`
}
