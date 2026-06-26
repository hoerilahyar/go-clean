package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

// POST /roles/:id/permissions
func (h *AssignmentHandler) AssignPermissionByRoleID(c *gin.Context) {

	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var req request.AssignRolePermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.AssignRolePermissions(c.Request.Context(), roleID, req); err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "role permissions assigned", nil)
}

// PUT /roles/:id/permissions
func (h *AssignmentHandler) ReplacePermissionByRoleID(c *gin.Context) {

	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var req request.UpdateRolePermissionsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.ReplaceRolePermissions(c.Request.Context(), roleID, req); err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "role permissions updated", nil)
}

// GET /roles/:id/permissions
func (h *AssignmentHandler) GetPermissionByRoleID(c *gin.Context) {

	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	result, err := h.usecase.GetRolePermissions(
		c.Request.Context(),
		roleID,
	)

	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", result)
}

// REVOKE /roles/:id/permissions/:permission_id
func (h *AssignmentHandler) RevokePermissionFromRole(c *gin.Context) {

	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid permission id")
		return
	}

	if err := h.usecase.DeleteRolePermission(
		c.Request.Context(),
		roleID,
		permissionID,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "role permission deleted", nil)
}
