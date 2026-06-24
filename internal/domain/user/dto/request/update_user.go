package request

type UpdateUserRequest struct {
	ID          uint64 `json:"id" binding:"required"`
	FullName    string `json:"full_name" binding:"required,max=255"`
	Username    string `json:"username" binding:"required,max=100"`
	Email       string `json:"email" binding:"required,email,max=255"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive blocked"`
}
