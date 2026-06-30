package response

import (
	"github.com/hoerilahyar/go-clean/internal/domain/menu/entity"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type GetMenusResponse struct {
	Items      []entity.Menu     `json:"items"`
	Pagination *utils.Pagination `json:"pagination,omitempty"`
}
