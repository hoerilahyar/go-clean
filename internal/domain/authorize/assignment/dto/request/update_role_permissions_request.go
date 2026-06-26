package request

type UpdateRolePermissionsRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required"`
}
