package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProductPromotionProduct struct {
	DB *gorm.DB
}

func NewRepositoryProductPromotionProduct(db *gorm.DB) *RepositoryProductPromotionProduct {
	return &RepositoryProductPromotionProduct{DB: db}
}

func (r *RepositoryProductPromotionProduct) Create(productPromotionProduct entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error) {
	err := r.DB.Create(&productPromotionProduct).Error
	if err != nil {
		return nil, err
	}
	return &productPromotionProduct, nil
}

func (r *RepositoryProductPromotionProduct) GetAll(store_id string, promotion_id string) ([]*entity.EntityProductPromotionProduct, error) {
	var productPromotionProducts []*entity.EntityProductPromotionProduct
	err := r.DB.Where("store.id =?", store_id).Where("promotion.id =?", promotion_id).Joins("Product").Joins("JOIN store ON store.id = product.store_id").Joins("JOIN promotion ON promotion.id = product_promotion.promotion_id").Find(&productPromotionProducts).Error
	if err != nil {
		return nil, err
	}
	return productPromotionProducts, nil
}

func (r *RepositoryProductPromotionProduct) GetByID(store_id string, promotion_id string, id string) (*entity.EntityProductPromotionProduct, error) {
	var productPromotionProduct entity.EntityProductPromotionProduct
	products, err := r.GetAll(store_id, promotion_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&products).First(&productPromotionProduct).Error
	if err != nil {
		return nil, err
	}
	return &productPromotionProduct, nil
}

func (r *RepositoryProductPromotionProduct) Update(productPromotionProduct entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error) {
	err := r.DB.Save(&productPromotionProduct).Error
	if err != nil {
		return nil, err
	}
	return &productPromotionProduct, nil
}

func (r *RepositoryProductPromotionProduct) Delete(store_id string, promotion_id string, id string) error {
	var productPromotionProduct entity.EntityProductPromotionProduct
	products, err := r.GetAll(store_id, promotion_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&products).Delete(&productPromotionProduct).Error
	if err != nil {
		return err
	}
	return nil
}
