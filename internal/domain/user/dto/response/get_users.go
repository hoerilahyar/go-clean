package response

import (
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

type GetUsersResponse struct {
	Items      []*entity.User    `json:"items"`
	Pagination *utils.Pagination `json:"pagination,omitempty"`
}
