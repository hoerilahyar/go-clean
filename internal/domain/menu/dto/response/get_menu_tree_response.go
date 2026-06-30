package response

type MenuTreeResponse struct {
	ID        uint64             `json:"id"`
	ParentID  *uint64            `json:"parent_id,omitempty"`
	Name      string             `json:"name"`
	Slug      string             `json:"slug"`
	Path      string             `json:"path"`
	Icon      *string            `json:"icon,omitempty"`
	SortOrder int                `json:"sort_order"`
	IsVisible bool               `json:"is_visible"`
	IsActive  bool               `json:"is_active"`
	Children  []MenuTreeResponse `json:"children,omitempty"`
}
