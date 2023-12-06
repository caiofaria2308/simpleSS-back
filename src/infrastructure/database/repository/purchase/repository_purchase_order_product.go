package repository

import (
	entity "main/entity/purchase"

	"gorm.io/gorm"
)

type RepositoryPurchaseOrderProduct struct {
	DB *gorm.DB
}

func NewRepositoryPurchaseOrderProduct(db *gorm.DB) *RepositoryPurchaseOrderProduct {
	return &RepositoryPurchaseOrderProduct{DB: db}
}

func (r *RepositoryPurchaseOrderProduct) Create(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error) {
	err := r.DB.Create(&purchaseOrderProduct).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrderProduct, nil
}

func (r *RepositoryPurchaseOrderProduct) GetAll(store_id string, purchaseOrderID string) ([]*entity.EntityPurchaseOrderProduct, error) {
	var purchaseOrderProducts []*entity.EntityPurchaseOrderProduct
	err := r.DB.Where("order_id =? and store.id = ?", purchaseOrderID, store_id).Joins("Purchase").Joins("JOIN store ON store.id = purchase.store_id").Find(&purchaseOrderProducts).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrderProducts, nil
}

func (r *RepositoryPurchaseOrderProduct) GetByID(store_id string, purchaseOrderID string, id string) (*entity.EntityPurchaseOrderProduct, error) {
	var purchaseOrderProduct entity.EntityPurchaseOrderProduct
	purchases, err := r.GetAll(store_id, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", id).Find(&purchases).First(&purchaseOrderProduct).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrderProduct, nil
}

func (r *RepositoryPurchaseOrderProduct) Update(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error) {
	err := r.DB.Save(&purchaseOrderProduct).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrderProduct, nil
}

func (r *RepositoryPurchaseOrderProduct) Delete(store_id string, purchaseOrderID string, id string) error {
	var purchaseOrderProduct entity.EntityPurchaseOrderProduct
	purchases, err := r.GetAll(store_id, purchaseOrderID)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", id).Find(&purchases).Delete(&purchaseOrderProduct).Error
	if err != nil {
		return err
	}
	return nil
}
