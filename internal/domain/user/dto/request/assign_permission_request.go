package request

type AssignPermissionRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required,min=1,dive,gt=0"`
}
