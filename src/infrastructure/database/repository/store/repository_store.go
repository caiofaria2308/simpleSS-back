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

func (r *RepositoryStore) GetAll(chain_id string) (*[]entity.EntityStore, error) {
	var stores []entity.EntityStore
	err := r.DB.Where("chain.id =?", chain_id).Joins("Chain").Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return &stores, nil
}

func (r *RepositoryStore) GetByID(chain_id string, id string) (*entity.EntityStore, error) {
	var store entity.EntityStore
	stores, err := r.GetAll(chain_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&stores).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *RepositoryStore) FilterByCNPJ(chain_id string, cnpj string) (store *[]entity.EntityStore, err error) {
	store, err = r.GetAll(chain_id)
	err = r.DB.Where("cnpj =?", cnpj).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (r *RepositoryStore) Search(chain_id string, search string) (store *[]entity.EntityStore, err error) {
	store, err = r.GetAll(chain_id)
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

func (r *RepositoryStore) Delete(chain_id string, id string) error {
	stores, err := r.GetAll(chain_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&stores).Error
	if err != nil {
		return err
	}
	return nil
}
