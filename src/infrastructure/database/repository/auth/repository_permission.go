package repository

import (
	"fmt"
	entity "main/entity/auth"

	"gorm.io/gorm"
)

type RepositoryPermission struct {
	DB *gorm.DB
}

func (r *RepositoryPermission) Create(permission entity.EntityPermission) (*entity.EntityPermission, error) {
	err := r.DB.Create(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *RepositoryPermission) GetAll() (*[]entity.EntityPermission, error) {
	var permissions []entity.EntityPermission
	err := r.DB.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return &permissions, nil
}

func (r *RepositoryPermission) GetBySlug(slug string) (*entity.EntityPermission, error) {
	var permission entity.EntityPermission
	err := r.DB.Where("slug = ?", slug).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *RepositoryPermission) HasPermission(slug string, user entity.EntityUser) bool {
	permission, err := r.GetBySlug(slug)
	err = r.DB.Where("users.id = ?", user.ID).Joins("Users").Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	if permission.ID == "" {
		return false
	}
	return true
}
