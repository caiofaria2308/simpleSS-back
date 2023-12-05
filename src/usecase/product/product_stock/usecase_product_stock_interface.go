package usecase_product_stock

import entity "main/entity/product"

type IRepositoryProductStock interface {
	GetAll(store_id string) (*[]entity.EntityProductStock, error)
	GetByID(store_id string, id string) (*entity.EntityProductStock, error)
	FilterByProductID(store_id string, product_id string) (*[]entity.EntityProductStock, error)
	Create(productStock entity.EntityProductStock) (*entity.EntityProductStock, error)
	Update(productStock entity.EntityProductStock) (*entity.EntityProductStock, error)
	Delete(store_id string, id string) error
}

type IUseCaseProductStock interface {
	GetAll(store_id string) (*[]entity.EntityProductStock, error)
	GetByID(store_id string, id string) (*entity.EntityProductStock, error)
	FilterByProductID(store_id string, product_id string) (*[]entity.EntityProductStock, error)
	Create(productStock entity.EntityProductStock) (*entity.EntityProductStock, error)
	Update(productStock entity.EntityProductStock) (*entity.EntityProductStock, error)
	Delete(store_id string, id string) error
}
