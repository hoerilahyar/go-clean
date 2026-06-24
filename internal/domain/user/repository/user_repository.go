package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context, req request.GetUsersRequest) ([]*entity.User, int64, error) {

	var (
		query strings.Builder
		args  []interface{}
	)

	query.WriteString(`
		SELECT
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
		FROM users
		WHERE deleted_at IS NULL
	`)

	// Filter User ID
	if req.UserID != nil {
		query.WriteString(" AND id = ?")
		args = append(args, *req.UserID)
	}

	// Filter Status
	if req.Status != "" {
		query.WriteString(" AND status = ?")
		args = append(args, req.Status)
	}

	// Search
	if req.Search != "" {

		keywords := strings.Fields(strings.TrimSpace(req.Search))

		query.WriteString(" AND (")

		for i, keyword := range keywords {

			if i > 0 {
				query.WriteString(" OR ")
			}

			query.WriteString(`
			(
				full_name LIKE ?
				OR username LIKE ?
				OR email LIKE ?
			)
		`)

			like := "%" + keyword + "%"

			args = append(args,
				like,
				like,
				like,
			)
		}

		query.WriteString(")")
	}

	// ------------------------
	// COUNT
	// ------------------------

	countQuery := "SELECT COUNT(*) FROM (" + query.String() + ") x"

	var total int64

	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// ------------------------
	// SORT
	// ------------------------

	sortBy := "id"

	switch req.SortBy {
	case "id", "full_name", "username", "email", "created_at":
		sortBy = req.SortBy
	}

	sortOrder := "ASC"

	if strings.EqualFold(req.SortOrder, "desc") {
		sortOrder = "DESC"
	}

	query.WriteString(fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder))

	// ------------------------
	// PAGINATION
	// ------------------------

	if req.UserID == nil {

		offset := (req.Page - 1) * req.Limit

		query.WriteString(" LIMIT ? OFFSET ?")

		args = append(args,
			req.Limit,
			offset,
		)
	}

	rows, err := r.db.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {

		user := &entity.User{}

		err := rows.Scan(
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
			return nil, 0, err
		}

		users = append(users, user)
	}

	fmt.Printf("%+v\n", req)
	fmt.Println(query.String())
	fmt.Println(args)
	return users, total, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	query := `SELECT id, full_name, username, email, phone_number, avatar, status, last_login_at, created_at, updated_at FROM users WHERE id = ? AND deleted_at IS NULL`
	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.FullName, &u.Username, &u.Email, &u.PhoneNumber, &u.Avatar, &u.Status, &u.LastLoginAt, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, full_name, username, email, password, status FROM users WHERE email = ? AND deleted_at IS NULL`
	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.FullName, &u.Username, &u.Email, &u.Password, &u.Status)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT id, full_name, username, email, password, status FROM users WHERE username = ? AND deleted_at IS NULL`
	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(&u.ID, &u.FullName, &u.Username, &u.Email, &u.Password, &u.Status)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) IsEmailExists(ctx context.Context, email string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE email = ?
			  AND deleted_at IS NULL
			  AND id <> ?
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, email, excludeID).Scan(&exists)
	if err != nil {
		return false, err
	}

	fmt.Printf("IsEmailExists: email=%s, excludeID=%d, exists=%v\n query=%s", email, excludeID, exists, query)

	return exists, nil
}

func (r *userRepository) IsUsernameExists(ctx context.Context, username string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE username = ?
			  AND deleted_at IS NULL
			  AND id <> ?
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, username, excludeID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) IsPhoneNumberExists(ctx context.Context, phoneNumber string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE phone_number = ?
			  AND deleted_at IS NULL
			  AND id <> ?
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, phoneNumber, excludeID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) error {
	query := `INSERT INTO users (full_name, username, email, phone_number, password, status, created_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, u.FullName, u.Username, u.Email, u.PhoneNumber, u.Password, u.Status, u.CreatedBy)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	u.ID = uint64(id)
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {

	query := `
		UPDATE users
		SET
			full_name = ?,
			username = ?,
			email = ?,
			phone_number = ?,
			status = ?,
			updated_at = NOW()
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.Status,
		user.ID,
	)

	return err
}

func (r *userRepository) SoftDelete(ctx context.Context, id uint64, deletedBy uint64) error {

	query := `
		UPDATE users
		SET
			deleted_at = NOW(),
			deleted_by = ?
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, deletedBy, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepository) HardDelete(ctx context.Context, id uint64) error {

	query := `
		DELETE
		FROM users
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepository) Restore(ctx context.Context, id uint64) error {

	query := `
		UPDATE users
		SET
			deleted_at = NULL,
			deleted_by = NULL
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
