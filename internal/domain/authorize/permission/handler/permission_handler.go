package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/usecase"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	"github.com/hoerilahyar/go-clean/pkg/response"
)

type PermissionHandler struct {
	usecase usecase.PermissionUsecase
}

func NewPermissionHandler(
	uc usecase.PermissionUsecase,
) *PermissionHandler {

	return &PermissionHandler{
		usecase: uc,
	}
}

func (h *PermissionHandler) GetAll(
	c *gin.Context,
) {

	var filter request.PermissionFilter

	filter, err := httpx.BindQuery[request.PermissionFilter](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	permissions, err := h.usecase.GetAll(
		c.Request.Context(),
		filter,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, permissions, "Permissions retrieved successfully")
}

func (h *PermissionHandler) GetByID(
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

	permission, err := h.usecase.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, permission, "Permission retrieved successfully")
}

func (h *PermissionHandler) GetByGroup(
	c *gin.Context,
) {

	group, err := httpx.ParamString(
		c,
		"group",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	permissions, err := h.usecase.GetByGroup(
		c.Request.Context(),
		group,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, permissions, "Permissions retrieved successfully")
}

func (h *PermissionHandler) Create(
	c *gin.Context,
) {

	req, err := httpx.BindJSON[entity.Permission](c)
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

	response.Created(c, req, "Permission created successfully")
}

func (h *PermissionHandler) Update(
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

	req, err := httpx.BindJSON[entity.Permission](c)
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

	response.Success(c, req, "Permission updated successfully")
}

func (h *PermissionHandler) Delete(
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

	err = h.usecase.Delete(
		c.Request.Context(),
		id,
		userID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Permission deleted successfully")
}
