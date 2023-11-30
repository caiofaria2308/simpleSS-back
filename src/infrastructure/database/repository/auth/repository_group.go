package repository

import (
	entity "main/entity/auth"

	"gorm.io/gorm"
)

type RepositoryGroup struct {
	DB *gorm.DB
}

func (r *RepositoryGroup) Create(group entity.EntityGroup) (*entity.EntityGroup, error) {
	err := r.DB.Create(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *RepositoryGroup) GetAll() (*[]entity.EntityGroup, error) {
	var groups []entity.EntityGroup
	err := r.DB.Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return &groups, nil
}

func (r *RepositoryGroup) GetByID(id string) (*entity.EntityGroup, error) {
	var group entity.EntityGroup
	err := r.DB.Where("id = ?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *RepositoryGroup) GetByName(name string) (*entity.EntityGroup, error) {
	var group entity.EntityGroup
	err := r.DB.Where("name =?", name).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *RepositoryGroup) Update(group entity.EntityGroup) (*entity.EntityGroup, error) {
	err := r.DB.Save(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *RepositoryGroup) Delete(id string) error {
	var group entity.EntityGroup
	err := r.DB.Where("id = ?", id).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}
