package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	roleEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
)

func (r *assignmentRepository) GetMe(
	ctx context.Context,
	userID uint64,
) (*entity.Me, error) {

	me := &entity.Me{}

	// ==========================
	// User
	// ==========================

	userQuery := `
		SELECT
			id,
			name,
			username,
			email,
			status
		FROM users
		WHERE id = ?
		AND deleted_at IS NULL
	`

	err := r.db.QueryRowContext(ctx, userQuery, userID).Scan(
		&me.User.ID,
		&me.User.FullName,
		&me.User.Username,
		&me.User.Email,
		&me.User.Status,
	)
	if err != nil {
		return nil, err
	}

	// ==========================
	// Roles
	// ==========================

	roleQuery := `
		SELECT
			r.id,
			r.name,
			r.slug,
			r.description
		FROM user_roles ur
		INNER JOIN roles r
			ON r.id = ur.role_id
		WHERE ur.user_id = ?
		AND r.deleted_at IS NULL
		ORDER BY r.name
	`

	roleRows, err := r.db.QueryContext(ctx, roleQuery, userID)
	if err != nil {
		return nil, err
	}
	defer roleRows.Close()

	for roleRows.Next() {

		var role roleEntity.Role

		if err := roleRows.Scan(
			&role.ID,
			&role.Name,
			&role.Slug,
			&role.Description,
		); err != nil {
			return nil, err
		}

		me.Roles = append(me.Roles, role)
	}

	if err := roleRows.Err(); err != nil {
		return nil, err
	}

	// ==========================
	// Permissions
	// ==========================

	permissionQuery := `
		SELECT DISTINCT
			p.id,
			p.name,
			p.slug,
			p.group_name,
			p.description
		FROM permissions p
		INNER JOIN role_permissions rp
			ON rp.permission_id = p.id
		INNER JOIN user_roles ur
			ON ur.role_id = rp.role_id
		WHERE ur.user_id = ?
		AND p.deleted_at IS NULL

		UNION

		SELECT DISTINCT
			p.id,
			p.name,
			p.slug,
			p.group_name,
			p.description
		FROM permissions p
		INNER JOIN user_permissions up
			ON up.permission_id = p.id
		WHERE up.user_id = ?
		AND p.deleted_at IS NULL

		ORDER BY slug
	`

	permissionRows, err := r.db.QueryContext(ctx, permissionQuery, userID, userID)
	if err != nil {
		return nil, err
	}
	defer permissionRows.Close()

	for permissionRows.Next() {

		var permission permissionEntity.Permission

		if err := permissionRows.Scan(
			&permission.ID,
			&permission.Name,
			&permission.Slug,
			&permission.GroupName,
			&permission.Description,
		); err != nil {
			return nil, err
		}

		me.Permissions = append(me.Permissions, permission)
	}

	if err := permissionRows.Err(); err != nil {
		return nil, err
	}

	return me, nil
}
