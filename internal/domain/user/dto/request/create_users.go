package request

type CreateUserRequest struct {
	FullName    string `json:"full_name" binding:"required,max=255"`
	Username    string `json:"username" binding:"required,max=100"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required,min=8"`
	Status      string `json:"status"`
}
