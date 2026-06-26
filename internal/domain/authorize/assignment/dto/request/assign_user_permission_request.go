package request

type AssignUserPermissionRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required,min=1"`
}
