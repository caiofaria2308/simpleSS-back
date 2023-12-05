package usecase_sale_product

import entity "main/entity/sale"

type IRepositorySaleProduct interface {
	GetAll(store_id string, sale_id string) ([]*entity.EntitySaleProduct, error)
	GetByID(store_id string, sale_id string, id string) (*entity.EntitySaleProduct, error)
	Create(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error)
	Update(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error)
	Delete(store_id string, sale_id string, id string) error
}

type IUseCaseSaleProduct interface {
	GetAll(store_id string, sale_id string) ([]*entity.EntitySaleProduct, error)
	GetByID(store_id string, sale_id string, id string) (*entity.EntitySaleProduct, error)
	Create(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error)
	Update(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error)
	Delete(store_id string, sale_id string, id string) error
}
