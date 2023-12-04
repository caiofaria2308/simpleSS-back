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

func (r *RepositoryStoreEmployee) GetAll(store_id string) (*[]entity.EntityStoreEmployee, error) {
	var storeEmployees []entity.EntityStoreEmployee
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&storeEmployees).Error
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

func (r *RepositoryStoreEmployee) SearchByUserName(store_id string, name string) (storeEmployees *[]entity.EntityStoreEmployee, err error) {
	storeEmployees, err = r.GetAll(store_id)
	err = r.DB.Where("user.name LIKE ?", "%"+name+"%").Joins("User").Find(&storeEmployees).Error
	if err != nil {
		return nil, err
	}
	return storeEmployees, nil
}

func (r *RepositoryStoreEmployee) SearchByEmail(store_id string, email string) (storeEmployees *[]entity.EntityStoreEmployee, err error) {
	storeEmployees, err = r.GetAll(store_id)
	err = r.DB.Where("user.email LIKE ?", "%"+email+"%").Joins("User").Find(&storeEmployees).Error
	if err != nil {
		return nil, err
	}
	return storeEmployees, nil
}
