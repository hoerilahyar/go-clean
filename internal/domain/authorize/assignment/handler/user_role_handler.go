package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

// AssignUserRoles assigns roles to a user.
func (h *AssignmentHandler) AssignUserRoles(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Bind and validate request body.
	req, err := httpx.BindJSON[request.AssignUserRoleRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Assign roles.
	if err := h.usecase.AssignUserRoles(
		c.Request.Context(),
		userID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(
		c,
		nil,
		"User roles assigned successfully",
	)
}

// ReplaceUserRoles replaces all roles assigned to a user.
func (h *AssignmentHandler) ReplaceUserRoles(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Bind and validate request body.
	req, err := httpx.BindJSON[request.UpdateUserRolesRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Replace roles.
	if err := h.usecase.ReplaceUserRoles(
		c.Request.Context(),
		userID,
		req,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User roles updated successfully")
}

// GetUserRoles returns roles assigned to a user.
func (h *AssignmentHandler) GetUserRoles(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Retrieve user roles.
	res, err := h.usecase.GetUserRoles(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}

// DeleteUserRole removes a role from a user.
func (h *AssignmentHandler) DeleteUserRole(c *gin.Context) {

	// Retrieve user ID.
	userID, err := httpx.ParamUint64(c, "id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Retrieve role ID.
	roleID, err := httpx.ParamUint64(c, "role_id")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Delete user role.
	if err := h.usecase.DeleteUserRole(
		c.Request.Context(),
		userID,
		roleID,
	); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil, "User role deleted successfully")
}
