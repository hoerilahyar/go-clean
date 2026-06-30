package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

// AssignUserPermissions assigns permissions to a user.
func (h *AssignmentHandler) AssignUserPermissions(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Bind request body.
	req, err := httpx.BindJSON[request.AssignUserPermissionRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Assign permissions.
	if err := h.usecase.AssignUserPermissions(
		c.Request.Context(),
		userID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User permissions assigned successfully")
}

// ReplaceUserPermissions replaces all permissions assigned to a user.
func (h *AssignmentHandler) ReplaceUserPermissions(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Bind request body.
	req, err := httpx.BindJSON[request.UpdateUserPermissionsRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Replace permissions.
	if err := h.usecase.ReplaceUserPermissions(
		c.Request.Context(),
		userID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User permissions updated successfully")
}

// GetUserPermissions returns permissions assigned to a user.
func (h *AssignmentHandler) GetUserPermissions(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Retrieve user permissions.
	res, err := h.usecase.GetUserPermissions(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}

// DeleteUserPermission removes a permission from a user.
func (h *AssignmentHandler) DeleteUserPermission(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Retrieve permission ID.
	permissionID, err := httpx.ParamUint64(c, "permission_id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Delete user permission.
	if err := h.usecase.DeleteUserPermission(
		c.Request.Context(),
		userID,
		permissionID,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User permission deleted successfully")
}
