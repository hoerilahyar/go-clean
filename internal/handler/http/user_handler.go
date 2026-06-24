package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/usecase"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	var req request.GetUsersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	result, err := h.usecase.GetAll(c.Request.Context(), req)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "users retrieved", result)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	user, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.Created(c, "user created", user)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req request.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	user, err := h.usecase.Update(c.Request.Context(), req)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user updated", user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid id")
		return
	}
	deletedBy := uint64(1)

	err = h.usecase.Delete(c.Request.Context(), id, deletedBy)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user deleted", nil)
}
