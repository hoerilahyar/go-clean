package request

type PermissionFilter struct {
	ID        *uint64 `form:"id"`
	Name      string  `form:"name"`
	Slug      string  `form:"slug"`
	GroupName string  `form:"group_name"`
}
