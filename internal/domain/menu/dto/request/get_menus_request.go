package request

type GetMenusRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`

	ID       *uint64 `form:"id"`
	ParentID *uint64 `form:"parent_id"`

	Name string `form:"name"`
	Slug string `form:"slug"`
	Path string `form:"path"`

	IsVisible *bool `form:"is_visible"`
	IsActive  *bool `form:"is_active"`

	Search string `form:"search"`

	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
}
