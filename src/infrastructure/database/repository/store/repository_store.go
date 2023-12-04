package repository

import (
	entity "main/entity/store"

	"gorm.io/gorm"
)

type RepositoryStore struct {
	DB *gorm.DB
}

func (r *RepositoryStore) Create(store entity.EntityStore) (*entity.EntityStore, error) {
	err := r.DB.Create(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *RepositoryStore) GetAll() (*[]entity.EntityStore, error) {
	var stores []entity.EntityStore
	err := r.DB.Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return &stores, nil
}

func (r *RepositoryStore) GetByID(id string) (*entity.EntityStore, error) {
	var store entity.EntityStore
	err := r.DB.Where("id =?", id).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *RepositoryStore) FilterByChain(chain_id string) (store *[]entity.EntityStore, err error) {
	err = r.DB.Where("chain_id =?", chain_id).Find(&store).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (r *RepositoryStore) FilterByCNPJ(chain_id string, cnpj string) (store *[]entity.EntityStore, err error) {
	store, err = r.FilterByChain(chain_id)
	err = r.DB.Where("cnpj =?", cnpj).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (r *RepositoryStore) Search(chain_id string, search string) (store *[]entity.EntityStore, err error) {
	store, err = r.FilterByChain(chain_id)
	err = r.DB.Where("social_reason LIKE ? OR business_name LIKE ?", "%"+search+"%", "%"+search+"%").Find(&store).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (r *RepositoryStore) Update(store entity.EntityStore) (*entity.EntityStore, error) {
	err := r.DB.Save(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *RepositoryStore) Delete(store entity.EntityStore) error {
	err := r.DB.Delete(&store).Error
	if err != nil {
		return err
	}
	return nil
}
