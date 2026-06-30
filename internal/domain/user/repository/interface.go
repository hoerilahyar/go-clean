package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context, req request.GetUsersRequest) ([]*entity.User, int64, error)
	FindByID(ctx context.Context, id uint64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	IsUserExists(ctx context.Context, id uint64) (bool, error)
	IsEmailExists(ctx context.Context, email string, excludeID uint64) (bool, error)
	IsUsernameExists(ctx context.Context, username string, excludeID uint64) (bool, error)
	IsPhoneNumberExists(ctx context.Context, phoneNumber string, excludeID uint64) (bool, error)

	Create(ctx context.Context, user *entity.User) error

	Update(ctx context.Context, user *entity.User) error

	SoftDelete(ctx context.Context, id uint64, deletedBy uint64) error
	HardDelete(ctx context.Context, id uint64) error
	Restore(ctx context.Context, id uint64) error
}
