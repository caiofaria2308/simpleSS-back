package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProductStock struct {
	DB *gorm.DB
}

func (r *RepositoryProductStock) Create(productStock entity.EntityProductStock) (*entity.EntityProductStock, error) {
	err := r.DB.Create(&productStock).Error
	if err != nil {
		return nil, err
	}
	return &productStock, nil
}

func (r *RepositoryProductStock) GetAll(store_id string) ([]*entity.EntityProductStock, error) {
	var productStocks []*entity.EntityProductStock
	err := r.DB.Where("store.id =?", store_id).Joins("Product").Joins("JOIN store ON store.id = product.store_id").Find(&productStocks).Error
	if err != nil {
		return nil, err
	}
	return productStocks, nil
}

func (r *RepositoryProductStock) GetByID(id string) (*entity.EntityProductStock, error) {
	var productStock entity.EntityProductStock
	err := r.DB.Where("id =?", id).First(&productStock).Error
	if err != nil {
		return nil, err
	}
	return &productStock, nil
}

func (r *RepositoryProductStock) Update(productStock entity.EntityProductStock) (*entity.EntityProductStock, error) {
	err := r.DB.Save(&productStock).Error
	if err != nil {
		return nil, err
	}
	return &productStock, nil
}

func (r *RepositoryProductStock) Delete(id string) error {
	var productStock entity.EntityProductStock
	err := r.DB.Where("id =?", id).Delete(&productStock).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProductStock) FilterByProduct(store_id string, product_id string) (productStocks []*entity.EntityProductStock, err error) {
	productStocks, err = r.GetAll(store_id)
	err = r.DB.Where("product_id =?", product_id).Find(&productStocks).Error
	if err != nil {
		return nil, err
	}
	return productStocks, nil
}
