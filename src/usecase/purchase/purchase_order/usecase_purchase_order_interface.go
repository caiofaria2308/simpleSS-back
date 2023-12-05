package usecase_purchaseorder

import entity "main/entity/purchase"

type IRepositoryPurchaseOrder interface {
	GetAll(store_id string) ([]*entity.EntityPurchaseOrder, error)
	GetByID(store_id string, id string) (*entity.EntityPurchaseOrder, error)
	FilterbyNumber(store_id string, number string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByDate(store_id string, initialDate string, endDate string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByProvider(store_id string, provider string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByEmployee(store_id string, employee string) (purchases []*entity.EntityPurchaseOrder, err error)
	Create(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error)
	Update(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error)
	Delete(store_id string, id string) error
}

type IUseCasePurchaseOrder interface {
	GetAll(store_id string) ([]*entity.EntityPurchaseOrder, error)
	GetByID(store_id string, id string) (*entity.EntityPurchaseOrder, error)
	FilterbyNumber(store_id string, number string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByDate(store_id string, initialDate string, endDate string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByProvider(store_id string, provider string) (purchases []*entity.EntityPurchaseOrder, err error)
	FilterByEmployee(store_id string, employee string) (purchases []*entity.EntityPurchaseOrder, err error)
	Create(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error)
	Update(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error)
	Delete(store_id string, id string) error
}
