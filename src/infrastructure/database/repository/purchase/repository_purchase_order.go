package repository

import (
	entity "main/entity/purchase"

	"gorm.io/gorm"
)

type RepositoryPurchaseOrder struct {
	DB *gorm.DB
}

func NewRepositoryPurchaseOrder(db *gorm.DB) *RepositoryPurchaseOrder {
	return &RepositoryPurchaseOrder{DB: db}
}

func (r *RepositoryPurchaseOrder) Create(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error) {
	err := r.DB.Create(&purchaseOrder).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *RepositoryPurchaseOrder) GetAll(store_id string) ([]*entity.EntityPurchaseOrder, error) {
	var purchaseOrders []*entity.EntityPurchaseOrder
	err := r.DB.Where("store.id =?", store_id).Joins("Store").Find(&purchaseOrders).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *RepositoryPurchaseOrder) GetByID(store_id string, id string) (*entity.EntityPurchaseOrder, error) {
	var purchaseOrder entity.EntityPurchaseOrder
	purchases, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id = ?", id).Find(&purchases).First(&purchaseOrder).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *RepositoryPurchaseOrder) FilterbyNumber(store_id string, number string) (purchases []*entity.EntityPurchaseOrder, err error) {
	purchases, err = r.GetAll(store_id)
	number = "%" + number + "%"
	err = r.DB.Where("number like ?", number).Find(&purchases).Error
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (r *RepositoryPurchaseOrder) FilterByDate(store_id string, initialDate string, endDate string) (purchases []*entity.EntityPurchaseOrder, err error) {
	purchases, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("date between %s and %s", initialDate, endDate).Find(&purchases).Error
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (r *RepositoryPurchaseOrder) FilterByProvider(store_id string, provider string) (purchases []*entity.EntityPurchaseOrder, err error) {
	purchases, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("provider_id =?", provider).Find(&purchases).Error
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (r *RepositoryPurchaseOrder) FilterByEmployee(store_id string, employee string) (purchases []*entity.EntityPurchaseOrder, err error) {
	purchases, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("employee_id =?", employee).Find(&purchases).Error
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (r *RepositoryPurchaseOrder) Update(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error) {
	err := r.DB.Save(&purchaseOrder).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *RepositoryPurchaseOrder) Delete(store_id string, id string) error {
	var purchaseOrder entity.EntityPurchaseOrder
	purchases, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id = ?", id).Find(&purchases).Delete(&purchaseOrder).Error
	if err != nil {
		return err
	}
	return nil
}
