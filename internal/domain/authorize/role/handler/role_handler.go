package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type RoleHandler struct {
	usecase usecase.RoleUsecase
}

func NewRoleHandler(uc usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		usecase: uc,
	}
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	var req request.GetRolesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	roles, err := h.usecase.GetAll(c.Request.Context(), req)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", roles)
}

// func (h *RoleHandler) GetByID(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		utils.BadRequest(c, "invalid role id")
// 		return
// 	}

// 	role, err := h.usecase.GetByID(c.Request.Context(), id)
// 	if err != nil {
// 		utils.NotFound(c, err.Error())
// 		return
// 	}

// 	utils.OK(c, "success", role)
// }

func (h *RoleHandler) Create(c *gin.Context) {
	var role entity.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.Create(c.Request.Context(), &role); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Created(c, "role created", role)
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var role entity.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	role.ID = id

	if err := h.usecase.Update(c.Request.Context(), &role); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "role updated", role)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid role id")
		return
	}

	var deletedBy uint64 = 1

	if err := h.usecase.Delete(c.Request.Context(), id, deletedBy); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, "role deleted", nil)
}
