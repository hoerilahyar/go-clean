package request

type LoginRequest struct {
	Identity string `json:"identity" binding:"required"`
	Password string `json:"password" binding:"required"`
}
