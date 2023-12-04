package repository

import (
	entity "main/entity/client"

	"gorm.io/gorm"
)

type RepositoryClient struct {
	DB *gorm.DB
}

func (r *RepositoryClient) Create(client entity.EntityClient) (*entity.EntityClient, error) {
	err := r.DB.Create(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *RepositoryClient) GetAll(store_id string) ([]*entity.EntityClient, error) {
	var clients []*entity.EntityClient
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *RepositoryClient) GetByID(id string) (*entity.EntityClient, error) {
	var client entity.EntityClient
	err := r.DB.Where("id =?", id).First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *RepositoryClient) Update(client entity.EntityClient) (*entity.EntityClient, error) {
	err := r.DB.Save(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *RepositoryClient) Delete(id string) (*entity.EntityClient, error) {
	var client entity.EntityClient
	err := r.DB.Where("id =?", id).Delete(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *RepositoryClient) GetByCPF(store_id string, cpf string) (clients []*entity.EntityClient, err error) {
	clients, err = r.GetAll(store_id)
	err = r.DB.Where("cpf =?", cpf).First(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *RepositoryClient) Search(store_id string, search string) (clients []*entity.EntityClient, err error) {
	clients, err = r.GetAll(store_id)
	err = r.DB.Where("name LIKE ? OR cpf LIKE ?", "%"+search+"%", "%"+search+"%").Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}
