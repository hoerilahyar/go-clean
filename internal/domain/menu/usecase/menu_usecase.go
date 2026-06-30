package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/repository"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

type menuUsecase struct {
	repository repository.MenuRepository
}

func NewMenuUsecase(
	repository repository.MenuRepository,
) MenuUsecase {

	return &menuUsecase{
		repository: repository,
	}
}

// validateCreate validates create menu request.
func (u *menuUsecase) validateCreate(
	ctx context.Context,
	req request.CreateMenuRequest,
) error {

	exists, err := u.repository.IsNameExists(
		ctx,
		req.Name,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Menu name already exists",
		)
	}

	exists, err = u.repository.IsSlugExists(
		ctx,
		req.Slug,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Menu slug already exists",
		)
	}

	return nil
}

// validateUpdate validates update menu request.
func (u *menuUsecase) validateUpdate(
	ctx context.Context,
	req request.UpdateMenuRequest,
) error {

	exists, err := u.repository.IsMenuExists(
		ctx,
		req.ID,
	)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.NotFound(
			"Menu not found",
		)
	}

	exists, err = u.repository.IsNameExists(
		ctx,
		req.Name,
		req.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Menu name already exists",
		)
	}

	exists, err = u.repository.IsSlugExists(
		ctx,
		req.Slug,
		req.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Menu slug already exists",
		)
	}

	return nil
}

// findMenuByID retrieves a menu by ID.
func (u *menuUsecase) findMenuByID(
	ctx context.Context,
	id uint64,
) (*entity.Menu, error) {

	menu, err := u.repository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
func (u *menuUsecase) GetAll(
	ctx context.Context,
	req request.GetMenusRequest,
) (*response.GetMenusResponse, error) {

	menus, err := u.repository.FindAll(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	return &response.GetMenusResponse{
		Items: menus,
	}, nil
}

func (u *menuUsecase) GetByID(
	ctx context.Context,
	id uint64,
) (*entity.Menu, error) {

	return u.findMenuByID(
		ctx,
		id,
	)
}

func (u *menuUsecase) Create(
	ctx context.Context,
	req request.CreateMenuRequest,
) (*entity.Menu, error) {

	err := u.validateCreate(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	menu := &entity.Menu{
		ParentID:  req.ParentID,
		Name:      req.Name,
		Slug:      req.Slug,
		Path:      req.Path,
		Icon:      req.Icon,
		SortOrder: req.SortOrder,
		IsVisible: req.IsVisible,
		IsActive:  req.IsActive,
		CreatedBy: req.CreatedBy,
	}

	err = u.repository.Create(
		ctx,
		menu,
	)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
func (u *menuUsecase) Update(
	ctx context.Context,
	req request.UpdateMenuRequest,
) (*entity.Menu, error) {

	err := u.validateUpdate(
		ctx,
		req,
	)
	if err != nil {
		return nil, err
	}

	menu, err := u.findMenuByID(
		ctx,
		req.ID,
	)
	if err != nil {
		return nil, err
	}

	menu.ParentID = req.ParentID
	menu.Name = req.Name
	menu.Slug = req.Slug
	menu.Path = req.Path
	menu.Icon = req.Icon
	menu.SortOrder = req.SortOrder
	menu.IsVisible = req.IsVisible
	menu.IsActive = req.IsActive
	menu.UpdatedBy = &req.UpdatedBy

	err = u.repository.Update(
		ctx,
		menu,
	)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (u *menuUsecase) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	_, err := u.findMenuByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	err = u.repository.Delete(
		ctx,
		id,
		deletedBy,
	)
	if err != nil {
		return err
	}

	return nil
}
