package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/usecase"

	"github.com/hoerilahyar/go-clean/pkg/httpx"
	"github.com/hoerilahyar/go-clean/pkg/response"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(
	usecase usecase.UserUsecase,
) *UserHandler {

	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAll(
	c *gin.Context,
) {

	var req request.GetUsersRequest

	req, err := httpx.BindQuery[request.GetUsersRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	result, err := h.usecase.GetAll(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result, "Users retrieved successfully")
}

func (h *UserHandler) Create(
	c *gin.Context,
) {

	var req request.CreateUserRequest

	req, err := httpx.BindQuery[request.CreateUserRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	*req.CreatedBy, err = httpx.UserID(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	user, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, user, "User created successfully")
}

func (h *UserHandler) Update(
	c *gin.Context,
) {

	var req request.UpdateUserRequest

	req, err := httpx.BindQuery[request.UpdateUserRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	id, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	req.ID = id
	*req.UpdatedBy, err = httpx.UserID(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	user, err := h.usecase.Update(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, user, "User updated successfully")
}

func (h *UserHandler) Delete(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	userID, err := httpx.UserID(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = h.usecase.Delete(c.Request.Context(), id, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User deleted successfully")
}
