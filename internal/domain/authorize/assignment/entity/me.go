package entity

import (
	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	roleEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	userEntity "github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

type Me struct {
	User        userEntity.User
	Roles       []roleEntity.Role
	Permissions []permissionEntity.Permission
}
