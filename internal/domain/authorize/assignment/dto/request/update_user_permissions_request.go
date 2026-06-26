package request

type UpdateUserPermissionsRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required"`
}
