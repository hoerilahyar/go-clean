package usecase

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/user/repository"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	"github.com/hoerilahyar/go-clean/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(
	repository repository.UserRepository,
) UserUsecase {

	return &userUsecase{
		repository: repository,
	}
}

// validateCreate validates create user request.
func (u *userUsecase) validateCreate(
	ctx context.Context,
	req request.CreateUserRequest,
) error {

	exists, err := u.repository.IsEmailExists(
		ctx,
		req.Email,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Email already exists",
		)
	}

	exists, err = u.repository.IsUsernameExists(
		ctx,
		req.Username,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Username already exists",
		)
	}

	exists, err = u.repository.IsPhoneNumberExists(
		ctx,
		req.PhoneNumber,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Phone number already exists",
		)
	}

	return nil
}

// validateUpdate validates update user request.
func (u *userUsecase) validateUpdate(
	ctx context.Context,
	req request.UpdateUserRequest,
) error {

	exists, err := u.repository.IsUserExists(
		ctx,
		req.ID,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return apperror.NotFound(
				"User not found",
			)
		}

		return err
	}

	if !exists {
		return apperror.NotFound(
			"User not found",
		)
	}

	exists, err = u.repository.IsEmailExists(
		ctx,
		req.Email,
		req.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Email already exists",
		)
	}

	exists, err = u.repository.IsUsernameExists(
		ctx,
		req.Username,
		req.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Username already exists",
		)
	}

	exists, err = u.repository.IsPhoneNumberExists(
		ctx,
		req.PhoneNumber,
		req.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Phone number already exists",
		)
	}

	return nil
}

// findUserByID retrieves a user by ID.
func (u *userUsecase) findUserByID(
	ctx context.Context,
	id uint64,
) (*entity.User, error) {

	user, err := u.repository.FindByID(
		ctx,
		id,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NotFound(
				"User not found",
			)
		}

		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetAll(
	ctx context.Context,
	req request.GetUsersRequest,
) (*response.GetUsersResponse, error) {

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	users, total, err := u.repository.FindAll(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	if req.UserID != nil {

		return &response.GetUsersResponse{
			Items: users,
		}, nil
	}

	totalPages := int(
		math.Ceil(
			float64(total) / float64(req.Limit),
		),
	)

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

func (u *userUsecase) GetByID(
	ctx context.Context,
	id uint64,
) (*entity.User, error) {

	return u.findUserByID(
		ctx,
		id,
	)
}

func (u *userUsecase) Create(
	ctx context.Context,
	req request.CreateUserRequest,
) (*entity.User, error) {

	err := u.validateCreate(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to hash password",
			err,
		)
	}

	user := &entity.User{
		FullName:    req.FullName,
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: &req.PhoneNumber,
		Password:    string(hashedPassword),
		Avatar:      req.Avatar,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
	}

	err = u.repository.Create(
		ctx,
		user,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Update(
	ctx context.Context,
	req request.UpdateUserRequest,
) (*entity.User, error) {

	err := u.validateUpdate(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	user, err := u.findUserByID(
		ctx,
		req.ID,
	)
	if err != nil {
		return nil, err
	}

	user.FullName = req.FullName
	user.Username = req.Username
	user.Email = req.Email
	user.PhoneNumber = &req.PhoneNumber
	user.Avatar = req.Avatar
	user.Status = req.Status
	user.UpdatedBy = req.UpdatedBy

	err = u.repository.Update(
		ctx,
		user,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	user, err := u.findUserByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	if user.Status == "active" {
		return apperror.BadRequest(
			"Active user cannot be deleted",
		)
	}

	err = u.repository.SoftDelete(
		ctx,
		id,
		deletedBy,
	)
	if err != nil {
		return err
	}

	return nil
}
