package request

type CreatePermissionRequest struct {
	Name        string  `json:"name" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	GroupName   string  `json:"group_name" binding:"required"`
	Description *string `json:"description"`
}
