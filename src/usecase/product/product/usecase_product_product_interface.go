package usecase_product_product

import (
	entity "main/entity/product"
)

type IRepositoryProductProduct interface {
	GetAll(store_id string) ([]*entity.EntityProduct, error)
	GetByID(store_id string, id string) (*entity.EntityProduct, error)
	SearchByName(store_id string, name string) (*entity.EntityProduct, error)
	SearchByBarcode(store_id string, barcode string) (*entity.EntityProduct, error)
	SearchByGroup(store_id string, group_id string) (*entity.EntityProduct, error)
	Create(product entity.EntityProduct) (*entity.EntityProduct, error)
	Update(product entity.EntityProduct) (*entity.EntityProduct, error)
	Delete(store_id string, id string) error
}

type IUseCaseProductProduct interface {
	GetAll(store_id string) ([]*entity.EntityProduct, error)
	GetByID(store_id string, id string) (*entity.EntityProduct, error)
	SearchByName(store_id string, name string) (*entity.EntityProduct, error)
	SearchByBarcode(store_id string, barcode string) (*entity.EntityProduct, error)
	SearchByGroup(store_id string, group_id string) (*entity.EntityProduct, error)
	Create(product entity.EntityProduct) (*entity.EntityProduct, error)
	Update(product entity.EntityProduct) (*entity.EntityProduct, error)
	Delete(store_id string, id string) error
}
