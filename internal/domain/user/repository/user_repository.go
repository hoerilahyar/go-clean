package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

const userColumns = `
	id,
	full_name,
	username,
	email,
	phone_number,
	avatar,
	status,
	last_login_at,
	created_at,
	updated_at,
	deleted_at,
	created_by,
	updated_by,
	deleted_by
`

const userColumnsWithPassword = `
	id,
	full_name,
	username,
	email,
	phone_number,
	avatar,
	password,
	status,
	last_login_at,
	created_at,
	updated_at,
	deleted_at,
	created_by,
	updated_by,
	deleted_by
`

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) UserRepository {

	return &userRepository{
		db: db,
	}
}

// scanUser scans a user from sql.Row or sql.Rows.
func scanUser(scanner interface {
	Scan(dest ...any) error
}) (*entity.User, error) {

	user := &entity.User{}

	err := scanner.Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Avatar,
		&user.Status,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.DeletedBy,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// scanUserWithPassword scans a user with password from sql.Row or sql.Rows.
func scanUserWithPassword(scanner interface {
	Scan(dest ...any) error
}) (*entity.User, error) {

	user := &entity.User{}

	err := scanner.Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Avatar,
		&user.Password,
		&user.Status,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.DeletedBy,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// findOne executes a query that returns a single user.
func (r *userRepository) findOne(
	ctx context.Context,
	query string,
	args ...any,
) (*entity.User, error) {

	row := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	)

	user, err := scanUser(row)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, apperror.Internal(
			"Failed to retrieve user",
			err,
		)
	}

	return user, nil
}

// findOneWithPassword executes a query that returns a single user with password.
func (r *userRepository) findOneWithPassword(
	ctx context.Context,
	query string,
	args ...any,
) (*entity.User, error) {

	row := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	)

	user, err := scanUserWithPassword(row)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, apperror.Internal(
			"Failed to retrieve user",
			err,
		)
	}

	return user, nil
}

// exists executes an EXISTS query.
func (r *userRepository) exists(
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

		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}

		return false, apperror.Internal(
			"Failed to check user existence",
			err,
		)
	}

	return exists, nil
}

func (r *userRepository) FindAll(
	ctx context.Context,
	req request.GetUsersRequest,
) ([]*entity.User, int64, error) {

	baseQuery := `
		FROM users
		WHERE deleted_at IS NULL
	`

	args := make([]any, 0)

	if req.UserID != nil {
		baseQuery += " AND id = ?"
		args = append(args, *req.UserID)
	}

	if req.Status != "" {
		baseQuery += " AND status = ?"
		args = append(args, req.Status)
	}

	if req.Search != "" {

		baseQuery += `
			AND (
				full_name LIKE ?
				OR username LIKE ?
				OR email LIKE ?
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

	countQuery := `
		SELECT COUNT(*)
	` + baseQuery

	var total int64

	err := r.db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total)
	if err != nil {
		return nil, 0, apperror.Internal(
			"Failed to count users",
			err,
		)
	}

	query := `
		SELECT
			` + userColumns + `
	` + baseQuery

	sortBy := "id"

	switch req.SortBy {
	case "id",
		"full_name",
		"username",
		"email",
		"created_at":

		sortBy = req.SortBy
	}

	sortOrder := "ASC"

	if strings.EqualFold(req.SortOrder, "desc") {
		sortOrder = "DESC"
	}

	query += " ORDER BY " + sortBy + " " + sortOrder

	queryArgs := append([]any{}, args...)

	if req.UserID == nil {

		offset := (req.Page - 1) * req.Limit

		query += " LIMIT ? OFFSET ?"

		queryArgs = append(
			queryArgs,
			req.Limit,
			offset,
		)
	}

	rows, err := r.db.QueryContext(
		ctx,
		query,
		queryArgs...,
	)
	if err != nil {
		return nil, 0, apperror.Internal(
			"Failed to retrieve users",
			err,
		)
	}

	defer rows.Close()

	users := make([]*entity.User, 0)

	for rows.Next() {

		user, err := scanUser(rows)
		if err != nil {
			return nil, 0, apperror.Internal(
				"Failed to scan user",
				err,
			)
		}

		users = append(
			users,
			user,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, apperror.Internal(
			"Failed to iterate users",
			err,
		)
	}

	return users, total, nil
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*entity.User, error) {

	query := `
		SELECT
			` + userColumns + `
		FROM users
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(
		ctx,
		query,
		id,
	)
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {

	query := `
		SELECT
			` + userColumnsWithPassword + `
		FROM users
		WHERE email = ?
		  AND deleted_at IS NULL
	`

	return r.findOneWithPassword(
		ctx,
		query,
		email,
	)
}

func (r *userRepository) FindByUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {

	query := `
		SELECT
			` + userColumnsWithPassword + `
		FROM users
		WHERE username = ?
		  AND deleted_at IS NULL
	`

	return r.findOneWithPassword(
		ctx,
		query,
		username,
	)
}

func (r *userRepository) IsUserExists(
	ctx context.Context,
	id uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
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

func (r *userRepository) IsEmailExists(
	ctx context.Context,
	email string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		email,
		excludeID,
		excludeID,
	)
}

func (r *userRepository) IsUsernameExists(
	ctx context.Context,
	username string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE username = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		username,
		excludeID,
		excludeID,
	)
}

func (r *userRepository) IsPhoneNumberExists(
	ctx context.Context,
	phoneNumber string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE phone_number = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		phoneNumber,
		excludeID,
		excludeID,
	)
}

func (r *userRepository) Create(
	ctx context.Context,
	user *entity.User,
) error {

	query := `
		INSERT INTO users (
			full_name,
			username,
			email,
			phone_number,
			password,
			avatar,
			status,
			created_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.Avatar,
		user.Status,
		user.CreatedBy,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to create user",
			err,
		)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return apperror.Internal(
			"Failed to retrieve inserted user ID",
			err,
		)
	}

	user.ID = uint64(id)

	return nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *entity.User,
) error {

	query := `
		UPDATE users
		SET
			full_name = ?,
			username = ?,
			email = ?,
			phone_number = ?,
			avatar = ?,
			status = ?,
			updated_by = ?
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.Avatar,
		user.Status,
		user.UpdatedBy,
		user.ID,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to update user",
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

func (r *userRepository) SoftDelete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	query := `
		UPDATE users
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
			"Failed to delete user",
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

func (r *userRepository) HardDelete(
	ctx context.Context,
	id uint64,
) error {

	query := `
		DELETE
		FROM users
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to hard delete user",
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

func (r *userRepository) Restore(
	ctx context.Context,
	id uint64,
) error {

	query := `
		UPDATE users
		SET
			deleted_at = NULL,
			deleted_by = NULL
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to restore user",
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
