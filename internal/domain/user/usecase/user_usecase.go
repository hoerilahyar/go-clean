package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/user/repository"
	"github.com/hoerilahyar/go-clean/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetAll(ctx context.Context, req request.GetUsersRequest) (*response.GetUsersResponse, error) {

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	users, total, err := u.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	if req.UserID != nil {
		return &response.GetUsersResponse{
			Items: users,
		}, nil
	}

	return &response.GetUsersResponse{
		Items: users,
		Pagination: &utils.Pagination{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
			HasNext:    req.Page < totalPages,
			HasPrev:    req.Page > 1,
		},
	}, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *userUsecase) Create(ctx context.Context, req request.CreateUserRequest) (*entity.User, error) {

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		FullName:    req.FullName,
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: &req.PhoneNumber,
		Password:    string(hashedPassword),
		Status:      req.Status,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Update(ctx context.Context, req request.UpdateUserRequest) (*entity.User, error) {

	user, err := u.repo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	exists, err := u.repo.IsEmailExists(ctx, req.Email, req.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	exists, err = u.repo.IsUsernameExists(ctx, req.Username, req.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	exists, err = u.repo.IsPhoneNumberExists(ctx, req.PhoneNumber, req.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("phone number already exists")
	}

	user.FullName = req.FullName
	user.Username = req.Username
	user.Email = req.Email
	user.PhoneNumber = &req.PhoneNumber
	user.Status = req.Status

	if err := u.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, id uint64, deletedBy uint64) error {

	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if user.Status == "active" {
		// return errors.New("active user cannot be deleted")
	}

	return u.repo.SoftDelete(ctx, id, deletedBy)
}
