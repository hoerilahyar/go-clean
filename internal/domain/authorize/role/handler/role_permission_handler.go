package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type RolePermissionHandler struct {
	usecase usecase.RolePermissionUsecase
}

func NewRolePermissionHandler(
	uc usecase.RolePermissionUsecase,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		usecase: uc,
	}
}

func (h *RolePermissionHandler) GetPermissions(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	permissions, err := h.usecase.GetPermissions(c.Request.Context(), roleID)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", permissions)
}

func (h *RolePermissionHandler) Assign(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var req request.AssignPermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.Assign(
		c.Request.Context(),
		roleID,
		req.PermissionIDs,
	); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "permissions assigned", nil)
}

func (h *RolePermissionHandler) Sync(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var req request.AssignPermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.Sync(
		c.Request.Context(),
		roleID,
		req.PermissionIDs,
	); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "permissions synchronized", nil)
}

func (h *RolePermissionHandler) Remove(c *gin.Context) {
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

	if err := h.usecase.Remove(
		c.Request.Context(),
		roleID,
		permissionID,
	); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "permission removed", nil)
}
