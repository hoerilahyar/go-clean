package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

// POST /users/:id/permissions
func (h *AssignmentHandler) AssignPermissionByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	var req request.AssignUserPermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.AssignUserPermissions(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user permissions assigned", nil)
}

// PUT /users/:id/permissions
func (h *AssignmentHandler) ReplacePermissionByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	var req request.UpdateUserPermissionsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.ReplaceUserPermissions(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user permissions updated", nil)
}

// GET /users/:id/permissions
func (h *AssignmentHandler) GetPermissionByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	result, err := h.usecase.GetUserPermissions(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", result)
}

// REVOKE /users/:id/permissions/:permission_id
func (h *AssignmentHandler) RevokePermissionFromUser(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid permission id")
		return
	}

	if err := h.usecase.DeleteUserPermission(
		c.Request.Context(),
		userID,
		permissionID,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user permission deleted", nil)
}
