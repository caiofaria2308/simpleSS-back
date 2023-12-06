package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProduct struct {
	DB *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *RepositoryProduct {
	return &RepositoryProduct{DB: db}
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

func (r *RepositoryProduct) GetByID(store_id string, id string) (*entity.EntityProduct, error) {
	var product entity.EntityProduct
	products, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&products).First(&product).Error
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

func (r *RepositoryProduct) Delete(store_id string, id string) error {
	var product entity.EntityProduct
	products, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&products).Delete(&product).Error
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
