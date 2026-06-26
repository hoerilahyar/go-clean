package request

type AssignRolePermissionRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required,min=1"`
}
