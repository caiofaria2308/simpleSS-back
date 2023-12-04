package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProductGroup struct {
	DB *gorm.DB
}

func (r *RepositoryProductGroup) Create(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error) {
	err := r.DB.Create(&productGroup).Error
	if err != nil {
		return nil, err
	}
	return &productGroup, nil
}

func (r *RepositoryProductGroup) GetAll(store_id string) ([]*entity.EntityProductGroup, error) {
	var productGroups []*entity.EntityProductGroup
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&productGroups).Error
	if err != nil {
		return nil, err
	}
	return productGroups, nil
}

func (r *RepositoryProductGroup) GetByID(id string) (*entity.EntityProductGroup, error) {
	var productGroup entity.EntityProductGroup
	err := r.DB.Where("id =?", id).First(&productGroup).Error
	if err != nil {
		return nil, err
	}
	return &productGroup, nil
}

func (r *RepositoryProductGroup) Update(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error) {
	err := r.DB.Save(&productGroup).Error
	if err != nil {
		return nil, err
	}
	return &productGroup, nil
}

func (r *RepositoryProductGroup) Delete(id string) error {
	var productGroup entity.EntityProductGroup
	err := r.DB.Where("id =?", id).Delete(&productGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProductGroup) SearchByName(store_id string, name string) (productGroups []*entity.EntityProductGroup, err error) {
	productGroups, err = r.GetAll(store_id)
	err = r.DB.Where("name =?", name).Find(&productGroups).Error
	if err != nil {
		return nil, err
	}
	return productGroups, nil
}
