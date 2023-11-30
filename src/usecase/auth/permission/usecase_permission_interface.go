package usecase_permission

import entity "main/entity/auth"

type IRepositoryPermission interface {
	GetBySlug(slug string) (*entity.EntityPermission, error)
	GetAll() (*[]entity.EntityPermission, error)
	CreatePermission(permission entity.EntityPermission) (*entity.EntityPermission, error)
}

type IUseCasePermission interface {
	GetBySlug(slug string) (*entity.EntityPermission, error)
	GetAll() (*[]entity.EntityPermission, error)
	Create(permission entity.EntityPermission) (*entity.EntityPermission, error)
}
