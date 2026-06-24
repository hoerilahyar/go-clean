package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/role/usecase"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type RoleHandler struct {
	usecase usecase.RoleUsecase
}

func NewRoleHandler(uc usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{usecase: uc}
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	roles, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "roles retrieved", roles)
}

func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "invalid id")
		return
	}
	role, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		utils.NotFound(c, "role not found")
		return
	}
	utils.OK(c, "role retrieved", role)
}
