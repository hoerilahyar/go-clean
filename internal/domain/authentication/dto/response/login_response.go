package response

type LoginResponse struct {
	User         UserResponse    `json:"user"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	TokenType    string          `json:"token_type"`
	ExpiresIn    int64           `json:"expires_in"`
	Session      SessionResponse `json:"session"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type SessionResponse struct {
	ID        uint64 `json:"id"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}
