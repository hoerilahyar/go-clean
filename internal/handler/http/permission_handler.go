package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/permission/usecase"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type PermissionHandler struct {
	usecase usecase.PermissionUsecase
}

func NewPermissionHandler(uc usecase.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{
		usecase: uc,
	}
}

func (h *PermissionHandler) GetAll(c *gin.Context) {
	permissions, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "permissions retrieved", permissions)
}

func (h *PermissionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid permission id")
		return
	}

	permission, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		utils.NotFound(c, "permission not found")
		return
	}

	utils.OK(c, "permission retrieved", permission)
}

func (h *PermissionHandler) GetByGroup(c *gin.Context) {
	group := c.Param("group")

	permissions, err := h.usecase.GetByGroup(c.Request.Context(), group)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "permissions retrieved", permissions)
}

func (h *PermissionHandler) Create(c *gin.Context) {
	var permission entity.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.Create(c.Request.Context(), &permission); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Created(c, "permission created", permission)
}

func (h *PermissionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid permission id")
		return
	}

	var permission entity.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	permission.ID = id

	if err := h.usecase.Update(c.Request.Context(), &permission); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "permission updated", permission)
}

func (h *PermissionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid permission id")
		return
	}

	// nanti diambil dari JWT/Auth middleware
	var deletedBy uint64 = 1

	if err := h.usecase.Delete(c.Request.Context(), id, deletedBy); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "permission deleted", nil)
}
