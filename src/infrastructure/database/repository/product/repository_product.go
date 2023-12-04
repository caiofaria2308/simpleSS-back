package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProduct struct {
	DB *gorm.DB
}

func (r *RepositoryProduct) Create(product entity.EntityProduct) (*entity.EntityProduct, error) {
	err := r.DB.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *RepositoryProduct) GetAll(store_id string) ([]*entity.EntityProduct, error) {
	var products []*entity.EntityProduct
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryProduct) GetByID(id string) (*entity.EntityProduct, error) {
	var product entity.EntityProduct
	err := r.DB.Where("id =?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *RepositoryProduct) Update(product entity.EntityProduct) (*entity.EntityProduct, error) {
	err := r.DB.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *RepositoryProduct) Delete(id string) error {
	var product entity.EntityProduct
	err := r.DB.Where("id =?", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProduct) SearchByName(store_id string, name string) (products []*entity.EntityProduct, err error) {
	products, err = r.GetAll(store_id)
	err = r.DB.Where("name =?", name).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryProduct) SearchByBarcode(store_id string, barcode string) (products []*entity.EntityProduct, err error) {
	products, err = r.GetAll(store_id)
	err = r.DB.Where("barcode =?", barcode).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryProduct) SearchByGroup(store_id string, group_id string) (products []*entity.EntityProduct, err error) {
	products, err = r.GetAll(store_id)
	err = r.DB.Where("group.id = ?", group_id).Joins("Group").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
