package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

const menuColumns = `
	id,
	parent_id,
	name,
	slug,
	path,
	icon,
	sort_order,
	is_visible,
	is_active,
	created_at,
	updated_at,
	deleted_at,
	created_by,
	updated_by,
	deleted_by
`

type menuRepository struct {
	db *sql.DB
}

func NewMenuRepository(
	db *sql.DB,
) MenuRepository {

	return &menuRepository{
		db: db,
	}
}

// scanMenu scans a menu from sql.Row or sql.Rows.
func scanMenu(scanner interface {
	Scan(dest ...any) error
}) (*entity.Menu, error) {

	menu := &entity.Menu{}

	err := scanner.Scan(
		&menu.ID,
		&menu.ParentID,
		&menu.Name,
		&menu.Slug,
		&menu.Path,
		&menu.Icon,
		&menu.SortOrder,
		&menu.IsVisible,
		&menu.IsActive,
		&menu.CreatedAt,
		&menu.UpdatedAt,
		&menu.DeletedAt,
		&menu.CreatedBy,
		&menu.UpdatedBy,
		&menu.DeletedBy,
	)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

// findOne executes a query that returns a single menu.
func (r *menuRepository) findOne(
	ctx context.Context,
	query string,
	args ...any,
) (*entity.Menu, error) {

	row := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	)

	menu, err := scanMenu(row)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve menu",
			err,
		)
	}

	return menu, nil
}

// exists executes an EXISTS query.
func (r *menuRepository) exists(
	ctx context.Context,
	query string,
	args ...any,
) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(&exists)

	if err != nil {
		return false, apperror.Internal(
			"Failed to check menu existence",
			err,
		)
	}

	return exists, nil
}

// buildPlaceholders builds placeholders for IN queries.
func buildPlaceholders(
	count int,
) string {

	placeholders := make(
		[]string,
		count,
	)

	for i := range placeholders {
		placeholders[i] = "?"
	}

	return strings.Join(
		placeholders,
		",",
	)
}
func (r *menuRepository) FindAll(
	ctx context.Context,
	req request.GetMenusRequest,
) ([]entity.Menu, error) {

	query := `
		SELECT
			` + menuColumns + `
		FROM menus
		WHERE deleted_at IS NULL
	`

	args := make([]any, 0)

	if req.ID != nil {
		query += " AND id = ?"
		args = append(args, *req.ID)
	}

	if req.ParentID != nil {
		query += " AND parent_id = ?"
		args = append(args, *req.ParentID)
	}

	if req.Name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+req.Name+"%")
	}

	if req.Slug != "" {
		query += " AND slug LIKE ?"
		args = append(args, "%"+req.Slug+"%")
	}

	if req.Path != "" {
		query += " AND path LIKE ?"
		args = append(args, "%"+req.Path+"%")
	}

	if req.IsVisible != nil {
		query += " AND is_visible = ?"
		args = append(args, *req.IsVisible)
	}

	if req.IsActive != nil {
		query += " AND is_active = ?"
		args = append(args, *req.IsActive)
	}

	if req.Search != "" {
		query += `
			AND (
				name LIKE ?
				OR slug LIKE ?
				OR path LIKE ?
			)
		`

		keyword := "%" + req.Search + "%"

		args = append(
			args,
			keyword,
			keyword,
			keyword,
		)
	}

	query += " ORDER BY sort_order ASC, id ASC"

	rows, err := r.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve menus",
			err,
		)
	}

	defer rows.Close()

	menus := make([]entity.Menu, 0)

	for rows.Next() {

		menu, err := scanMenu(rows)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan menu",
				err,
			)
		}

		menus = append(
			menus,
			*menu,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate menus",
			err,
		)
	}

	return menus, nil
}

func (r *menuRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*entity.Menu, error) {

	query := `
		SELECT
			` + menuColumns + `
		FROM menus
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(
		ctx,
		query,
		id,
	)
}

func (r *menuRepository) FindBySlug(
	ctx context.Context,
	slug string,
) (*entity.Menu, error) {

	query := `
		SELECT
			` + menuColumns + `
		FROM menus
		WHERE slug = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(
		ctx,
		query,
		slug,
	)
}

func (r *menuRepository) FindByIDs(
	ctx context.Context,
	ids []uint64,
) ([]entity.Menu, error) {

	if len(ids) == 0 {
		return []entity.Menu{}, nil
	}

	args := make([]any, len(ids))

	for i, id := range ids {
		args[i] = id
	}

	query := `
		SELECT
			` + menuColumns + `
		FROM menus
		WHERE id IN (` + buildPlaceholders(len(ids)) + `)
		  AND deleted_at IS NULL
		ORDER BY sort_order ASC, id ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve menus",
			err,
		)
	}

	defer rows.Close()

	menus := make([]entity.Menu, 0, len(ids))

	for rows.Next() {

		menu, err := scanMenu(rows)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan menu",
				err,
			)
		}

		menus = append(
			menus,
			*menu,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate menus",
			err,
		)
	}

	return menus, nil
}

func (r *menuRepository) IsSlugExists(
	ctx context.Context,
	slug string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM menus
			WHERE slug = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		slug,
		excludeID,
		excludeID,
	)
}

func (r *menuRepository) IsNameExists(
	ctx context.Context,
	name string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM menus
			WHERE name = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		name,
		excludeID,
		excludeID,
	)
}

func (r *menuRepository) IsMenuExists(
	ctx context.Context,
	id uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM menus
			WHERE id = ?
			  AND deleted_at IS NULL
		)
	`

	return r.exists(
		ctx,
		query,
		id,
	)
}
func (r *menuRepository) Create(
	ctx context.Context,
	menu *entity.Menu,
) error {

	query := `
		INSERT INTO menus (
			parent_id,
			name,
			slug,
			path,
			icon,
			sort_order,
			is_visible,
			is_active,
			created_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		menu.ParentID,
		menu.Name,
		menu.Slug,
		menu.Path,
		menu.Icon,
		menu.SortOrder,
		menu.IsVisible,
		menu.IsActive,
		menu.CreatedBy,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to create menu",
			err,
		)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return apperror.Internal(
			"Failed to retrieve inserted menu ID",
			err,
		)
	}

	menu.ID = uint64(id)

	return nil
}

func (r *menuRepository) Update(
	ctx context.Context,
	menu *entity.Menu,
) error {

	query := `
		UPDATE menus
		SET
			parent_id = ?,
			name = ?,
			slug = ?,
			path = ?,
			icon = ?,
			sort_order = ?,
			is_visible = ?,
			is_active = ?,
			updated_by = ?
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		menu.ParentID,
		menu.Name,
		menu.Slug,
		menu.Path,
		menu.Icon,
		menu.SortOrder,
		menu.IsVisible,
		menu.IsActive,
		menu.UpdatedBy,
		menu.ID,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to update menu",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to retrieve affected rows",
			err,
		)
	}

	return nil
}

func (r *menuRepository) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	query := `
		UPDATE menus
		SET
			deleted_at = NOW(),
			deleted_by = ?
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		deletedBy,
		id,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to delete menu",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to retrieve affected rows",
			err,
		)
	}

	return nil
}
