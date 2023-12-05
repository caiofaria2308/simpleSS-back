package usecase_sale

import entity "main/entity/sale"

type IRepositorySale interface {
	GetAll(store_id string) ([]*entity.EntitySale, error)
	GetByID(store_id string, id string) (*entity.EntitySale, error)
	FilterByNumber(store_id string, number string) (sales []*entity.EntitySale, err error)
	FilterByDate(store_id string, initialDate string, endDate string) (sales []*entity.EntitySale, err error)
	FilterByEmployee(store_id string, employee_id string) (sales []*entity.EntitySale, err error)
	FilterByClient(store_id string, client_id string) (sales []*entity.EntitySale, err error)
	Create(sale entity.EntitySale) (*entity.EntitySale, error)
	Update(sale entity.EntitySale) (*entity.EntitySale, error)
	Delete(store_id string, id string) error
}

type IUseCaseSale interface {
	GetAll(store_id string) ([]*entity.EntitySale, error)
	GetByID(store_id string, id string) (*entity.EntitySale, error)
	FilterByNumber(store_id string, number string) (sales []*entity.EntitySale, err error)
	FilterByDate(store_id string, initialDate string, endDate string) (sales []*entity.EntitySale, err error)
	FilterByEmployee(store_id string, employee_id string) (sales []*entity.EntitySale, err error)
	FilterByClient(store_id string, client_id string) (sales []*entity.EntitySale, err error)
	Create(sale entity.EntitySale) (*entity.EntitySale, error)
	Update(sale entity.EntitySale) (*entity.EntitySale, error)
	Delete(store_id string, id string) error
}
