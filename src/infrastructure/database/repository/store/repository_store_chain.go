package repository

import (
	entity "main/entity/store"

	"gorm.io/gorm"
)

type RepositoryStoreChain struct {
	DB *gorm.DB
}

func (r *RepositoryStoreChain) Create(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error) {
	err := r.DB.Create(&storeChain).Error
	if err != nil {
		return nil, err
	}
	return &storeChain, nil
}

func (r *RepositoryStoreChain) GetAll() (*[]entity.EntityStoreChain, error) {
	var storeChains []entity.EntityStoreChain
	err := r.DB.Find(&storeChains).Error
	if err != nil {
		return nil, err
	}
	return &storeChains, nil
}

func (r *RepositoryStoreChain) GetByID(id string) (*entity.EntityStoreChain, error) {
	var storeChain entity.EntityStoreChain
	err := r.DB.Where("id = ?", id).First(&storeChain).Error
	if err != nil {
		return nil, err
	}
	return &storeChain, nil
}

func (r *RepositoryStoreChain) Update(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error) {
	err := r.DB.Save(&storeChain).Error
	if err != nil {
		return nil, err
	}
	return &storeChain, nil
}

func (r *RepositoryStoreChain) Delete(id string) error {
	var storeChain entity.EntityStoreChain
	err := r.DB.Where("id =?", id).Delete(&storeChain).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryStoreChain) SearchByName(name string) (*[]entity.EntityStoreChain, error) {
	var storeChains []entity.EntityStoreChain
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&storeChains).Error
	if err != nil {
		return nil, err
	}
	return &storeChains, nil
}
