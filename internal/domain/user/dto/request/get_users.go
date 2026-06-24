package request

type GetUsersRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`

	UserID *uint64 `json:"user_id"`
	Search string  `json:"search"`
	Status string  `json:"status"`

	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}
