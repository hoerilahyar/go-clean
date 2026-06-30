package request

type CreateMenuRequest struct {
	ParentID *uint64 `json:"parent_id"`

	Name string  `json:"name" binding:"required,max=100"`
	Slug string  `json:"slug" binding:"required,max=100"`
	Path string  `json:"path" binding:"required,max=255"`
	Icon *string `json:"icon"`

	SortOrder int  `json:"sort_order"`
	IsVisible bool `json:"is_visible"`
	IsActive  bool `json:"is_active"`

	CreatedBy uint64 `json:"-"`
}
