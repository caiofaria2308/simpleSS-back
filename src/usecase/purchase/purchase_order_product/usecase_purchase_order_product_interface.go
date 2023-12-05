package usecase_purchase_order_product

import entity "main/entity/purchase"

type IRepositoryPurchaseOrderProduct interface {
	GetAll(store_id string) ([]*entity.EntityPurchaseOrderProduct, error)
	GetByID(store_id string, purchaseOrderID string, id string) (*entity.EntityPurchaseOrderProduct, error)
	Create(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error)
	Update(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error)
	Delete(store_id string, purchaseOrderID string, id string) error
}

type IUseCasePurchaseOrderProduct interface {
	GetAll(store_id string) ([]*entity.EntityPurchaseOrderProduct, error)
	GetByID(store_id string, purchaseOrderID string, id string) (*entity.EntityPurchaseOrderProduct, error)
	Create(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error)
	Update(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error)
	Delete(store_id string, purchaseOrderID string, id string) error
}
