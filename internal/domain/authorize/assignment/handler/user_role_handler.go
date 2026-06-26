package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

// POST /users/:id/roles
func (h *AssignmentHandler) AssignRoleByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	var req request.AssignUserRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.AssignUserRoles(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user roles assigned", nil)
}

// PUT /users/:id/roles
func (h *AssignmentHandler) ReplaceRoleByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	var req request.UpdateUserRolesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.ReplaceUserRoles(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user roles updated", nil)
}

// GET /users/:id/roles
func (h *AssignmentHandler) GetRoleByUserID(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	result, err := h.usecase.GetUserRoles(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", result)
}

// REVOKE /users/:id/roles/:role_id
func (h *AssignmentHandler) RevokeRoleFromUser(c *gin.Context) {

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	if err := h.usecase.DeleteUserRole(
		c.Request.Context(),
		userID,
		roleID,
	); err != nil {

		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "user role deleted", nil)
}
