package request

type GetRolesRequest struct {
	ID       *uint64 `form:"id"`
	Name     string  `form:"name"`
	Slug     string  `form:"slug"`
	IsActive *bool   `form:"is_active"`
}
