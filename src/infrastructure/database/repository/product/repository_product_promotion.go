package repository

import (
	entity "main/entity/product"

	"gorm.io/gorm"
)

type RepositoryProductPromotion struct {
	DB *gorm.DB
}

func (r *RepositoryProductPromotion) Create(productPromotion entity.EntityProductPromotion) (*entity.EntityProductPromotion, error) {
	err := r.DB.Create(&productPromotion).Error
	if err != nil {
		return nil, err
	}
	return &productPromotion, nil
}

func (r *RepositoryProductPromotion) GetAll(store_id string) ([]*entity.EntityProductPromotion, error) {
	var productPromotions []*entity.EntityProductPromotion
	err := r.DB.Where("store.id =?", store_id).Joins("Product").Joins("JOIN store ON store.id = product.store_id").Find(&productPromotions).Error
	if err != nil {
		return nil, err
	}
	return productPromotions, nil
}

func (r *RepositoryProductPromotion) GetByID(store_id string, id string) (*entity.EntityProductPromotion, error) {
	var productPromotion entity.EntityProductPromotion
	products, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&products).First(&productPromotion).Error
	if err != nil {
		return nil, err
	}
	return &productPromotion, nil
}

func (r *RepositoryProductPromotion) Update(productPromotion entity.EntityProductPromotion) (*entity.EntityProductPromotion, error) {
	err := r.DB.Save(&productPromotion).Error
	if err != nil {
		return nil, err
	}
	return &productPromotion, nil
}

func (r *RepositoryProductPromotion) Delete(store_id string, id string) error {
	var productPromotion entity.EntityProductPromotion
	products, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&products).Delete(&productPromotion).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProductPromotion) FilterByPromotion(store_id string, promotion_id string) (productPromotions []*entity.EntityProductPromotion, err error) {
	productPromotions, err = r.GetAll(store_id)
	err = r.DB.Where("promotion_id =?", promotion_id).Find(&productPromotions).Error
	if err != nil {
		return nil, err
	}
	return productPromotions, nil
}

func (r *RepositoryProductPromotion) FilterByProduct(store_id string, product_id string) (productPromotions []*entity.EntityProductPromotion, err error) {
	productPromotions, err = r.GetAll(store_id)
	err = r.DB.Where("product_id =?", product_id).Find(&productPromotions).Error
	if err != nil {
		return nil, err
	}
	return productPromotions, nil
}
