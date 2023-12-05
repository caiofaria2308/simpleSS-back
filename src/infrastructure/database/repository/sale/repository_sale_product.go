package repository

import (
	entity "main/entity/sale"

	"gorm.io/gorm"
)

type RepositorySaleProduct struct {
	DB *gorm.DB
}

func (r *RepositorySaleProduct) Create(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error) {
	err := r.DB.Create(&saleProduct).Error
	if err != nil {
		return nil, err
	}
	return &saleProduct, nil
}

func (r *RepositorySaleProduct) GetAll(store_id string, sale_id string) ([]*entity.EntitySaleProduct, error) {
	var saleProducts []*entity.EntitySaleProduct
	err := r.DB.Where("sale_id =? and store.id = ?", sale_id, store_id).Joins("Sale").Joins("JOIN store ON store.id = sale.store_id").Find(&saleProducts).Error
	if err != nil {
		return nil, err
	}
	return saleProducts, nil
}

func (r *RepositorySaleProduct) GetByID(store_id string, sale_id string, id string) (*entity.EntitySaleProduct, error) {
	var saleProduct entity.EntitySaleProduct
	saleProducts, err := r.GetAll(store_id, sale_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&saleProducts).First(&saleProduct).Error
	if err != nil {
		return nil, err
	}
	return &saleProduct, nil
}

func (r *RepositorySaleProduct) Update(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error) {
	err := r.DB.Save(&saleProduct).Error
	if err != nil {
		return nil, err
	}
	return &saleProduct, nil
}

func (r *RepositorySaleProduct) Delete(store_id string, sale_id string, id string) (*entity.EntitySaleProduct, error) {
	var saleProduct entity.EntitySaleProduct
	saleProducts, err := r.GetAll(store_id, sale_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&saleProducts).Delete(&saleProduct).Error
	if err != nil {
		return nil, err
	}
	return &saleProduct, nil
}
