package repository

import (
	entity "main/entity/store"

	"gorm.io/gorm"
)

type RepositoryStoreEmployee struct {
	DB *gorm.DB
}

func (r *RepositoryStoreEmployee) Create(storeEmployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error) {
	err := r.DB.Create(&storeEmployee).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployee, nil
}

func (r *RepositoryStoreEmployee) GetAll() (*[]entity.EntityStoreEmployee, error) {
	var storeEmployees []entity.EntityStoreEmployee
	err := r.DB.Find(&storeEmployees).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployees, nil
}

func (r *RepositoryStoreEmployee) GetByID(id string) (*entity.EntityStoreEmployee, error) {
	var storeEmployee entity.EntityStoreEmployee
	err := r.DB.Where("id =?", id).First(&storeEmployee).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployee, nil
}

func (r *RepositoryStoreEmployee) Update(storeEmployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error) {
	err := r.DB.Save(&storeEmployee).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployee, nil
}

func (r *RepositoryStoreEmployee) Delete(id string) error {
	var storeEmployee entity.EntityStoreEmployee
	err := r.DB.Where("id =?", id).Delete(&storeEmployee).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryStoreEmployee) SearchByUserName(name string) (*[]entity.EntityStoreEmployee, error) {
	var storeEmployees []entity.EntityStoreEmployee
	err := r.DB.Where("users.name LIKE ?", "%"+name+"%").Joins("Users").Find(&storeEmployees).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployees, nil
}

func (r *RepositoryStoreEmployee) SearchByEmail(email string) (*[]entity.EntityStoreEmployee, error) {
	var storeEmployees []entity.EntityStoreEmployee
	err := r.DB.Where("users.email LIKE ?", "%"+email+"%").Joins("Users").Find(&storeEmployees).Error
	if err != nil {
		return nil, err
	}
	return &storeEmployees, nil
}
