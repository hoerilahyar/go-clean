package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	"github.com/hoerilahyar/go-clean/pkg/response"
)

type RoleHandler struct {
	usecase usecase.RoleUsecase
}

func NewRoleHandler(
	uc usecase.RoleUsecase,
) *RoleHandler {

	return &RoleHandler{
		usecase: uc,
	}
}

func (h *RoleHandler) GetAll(
	c *gin.Context,
) {

	var req request.GetRolesRequest

	req, err := httpx.BindQuery[request.GetRolesRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	roles, err := h.usecase.GetAll(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(
		c,
		roles,
		"Roles retrieved successfully",
	)
}

func (h *RoleHandler) GetByID(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	role, err := h.usecase.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, role, "Role retrieved successfully")
}

func (h *RoleHandler) Create(
	c *gin.Context,
) {

	req, err := httpx.BindJSON[entity.Role](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = h.usecase.Create(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, req, "Role created successfully")
}

func (h *RoleHandler) Update(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	req, err := httpx.BindJSON[entity.Role](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	req.ID = id

	err = h.usecase.Update(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, req, "Role updated successfully")
}
func (h *RoleHandler) Delete(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	userID, err := httpx.UserID(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = h.usecase.Delete(
		c.Request.Context(),
		id,
		userID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Role deleted successfully")
}
