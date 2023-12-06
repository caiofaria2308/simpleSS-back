package repository

import (
	entity "main/entity/client"

	"gorm.io/gorm"
)

type RepositoryClient struct {
	DB *gorm.DB
}

func NewRepositoryClient(db *gorm.DB) *RepositoryClient {
	return &RepositoryClient{DB: db}
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

func (r *RepositoryClient) GetByID(store_id string, id string) (*entity.EntityClient, error) {
	var client entity.EntityClient
	clients, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&clients).First(&client).Error
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

func (r *RepositoryClient) Delete(store_id string, id string) error {
	var client entity.EntityClient
	clients, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&clients).Delete(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryClient) GetByCPF(store_id string, cpf string) (client *entity.EntityClient, err error) {
	var clients []*entity.EntityClient
	clients, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("cpf =?", cpf).First(&clients).Error
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *RepositoryClient) Search(store_id string, search string) (clients []*entity.EntityClient, err error) {
	clients, err = r.GetAll(store_id)
	err = r.DB.Where("name LIKE ? OR cpf LIKE ?", "%"+search+"%", "%"+search+"%").Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}
