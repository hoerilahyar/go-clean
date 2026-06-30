package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/usecase"
	"github.com/hoerilahyar/go-clean/pkg/httpx"
	"github.com/hoerilahyar/go-clean/pkg/response"
)

type MenuHandler struct {
	usecase usecase.MenuUsecase
}

func NewMenuHandler(
	usecase usecase.MenuUsecase,
) *MenuHandler {

	return &MenuHandler{
		usecase: usecase,
	}
}

// GetAll handles get menus.
func (h *MenuHandler) GetAll(
	c *gin.Context,
) {

	req, err := httpx.BindQuery[request.GetMenusRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	menus, err := h.usecase.GetAll(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(
		c,
		menus,
		"Menus retrieved successfully",
	)
}

// GetByID handles get menu by ID.
func (h *MenuHandler) GetByID(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	menu, err := h.usecase.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(
		c,
		menu,
		"Menu retrieved successfully",
	)
}

// Create handles create menu.
func (h *MenuHandler) Create(
	c *gin.Context,
) {

	req, err := httpx.BindJSON[request.CreateMenuRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	menu, err := h.usecase.Create(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(
		c,
		menu,
		"Menu created successfully",
	)
}

// Update handles update menu.
func (h *MenuHandler) Update(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	req, err := httpx.BindJSON[request.UpdateMenuRequest](c)
	if err != nil {
		response.Error(c, err)
		return
	}

	req.ID = id

	// Ambil user login jika middleware auth sudah mengisinya.
	// Contoh:
	//
	// user := middleware.MustUser(c)
	// req.UpdatedBy = user.ID

	menu, err := h.usecase.Update(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(
		c,
		menu,
		"Menu updated successfully",
	)
}

// Delete handles delete menu.
func (h *MenuHandler) Delete(
	c *gin.Context,
) {

	id, err := httpx.ParamUint64(
		c,
		"id",
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Ambil user login dari middleware/auth.
	// Contoh (sesuaikan dengan project):
	//
	// claims := middleware.MustClaims(c)
	// deletedBy := claims.UserID
	//
	// atau
	//
	// user := middleware.MustUser(c)
	// deletedBy := user.ID

	var deletedBy uint64

	err = h.usecase.Delete(
		c.Request.Context(),
		id,
		deletedBy,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.NoContent(c)
}
