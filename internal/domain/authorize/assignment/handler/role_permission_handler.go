package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

// AssignRolePermissions assigns permissions to a role.
func (h *AssignmentHandler) AssignRolePermissions(c *gin.Context) {

	roleID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	req, err := httpx.BindJSON[request.AssignRolePermissionRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.usecase.AssignRolePermissions(
		c.Request.Context(),
		roleID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Role permissions assigned successfully")
}

// ReplaceRolePermissions replaces all permissions assigned to a role.
func (h *AssignmentHandler) ReplaceRolePermissions(c *gin.Context) {

	roleID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	req, err := httpx.BindJSON[request.UpdateRolePermissionsRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.usecase.ReplaceRolePermissions(
		c.Request.Context(),
		roleID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Role permissions updated successfully")
}

// GetRolePermissions returns permissions assigned to a role.
func (h *AssignmentHandler) GetRolePermissions(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	roleID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	result, err := h.usecase.GetRolePermissions(
		ctx,
		roleID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result)
}

// DeleteRolePermission removes a permission from a role.
func (h *AssignmentHandler) DeleteRolePermission(c *gin.Context) {

	roleID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	permissionID, err := httpx.ParamUint64(c, "permission_id")
	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.usecase.DeleteRolePermission(
		c.Request.Context(),
		roleID,
		permissionID,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Role permission deleted successfully")
}
