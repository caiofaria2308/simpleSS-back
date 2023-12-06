package repository

import (
	entity "main/entity/provider"

	"gorm.io/gorm"
)

type RepositoryProvider struct {
	DB *gorm.DB
}

func NewRepositoryProvider(db *gorm.DB) *RepositoryProvider {
	return &RepositoryProvider{DB: db}
}

func (r *RepositoryProvider) Create(provider entity.EntityProvider) (*entity.EntityProvider, error) {
	err := r.DB.Create(&provider).Error
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *RepositoryProvider) GetAll(store_id string) ([]*entity.EntityProvider, error) {
	var providers []*entity.EntityProvider
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *RepositoryProvider) GetByID(store_id string, id string) (*entity.EntityProvider, error) {
	var provider entity.EntityProvider
	providers, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&providers).First(&provider).Error
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *RepositoryProvider) Update(provider entity.EntityProvider) (*entity.EntityProvider, error) {
	err := r.DB.Save(&provider).Error
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *RepositoryProvider) Delete(store_id string, id string) error {
	var provider entity.EntityProvider
	providers, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&providers).Delete(&provider).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProvider) SearchByCNPJ(store_id string, cnpj string) (providers []*entity.EntityProvider, err error) {
	providers, err = r.GetAll(store_id)
	err = r.DB.Where("cnpj = ?", cnpj).Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *RepositoryProvider) SearchByName(store_id string, name string) (providers []*entity.EntityProvider, err error) {
	providers, err = r.GetAll(store_id)
	err = r.DB.Where("social_reason LIKE ? OR business_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}
