package usecase_product_group

import entity "main/entity/product"

type IRepositoryProductGroup interface {
	GetAll(store_id string) ([]*entity.EntityProductGroup, error)
	GetByID(store_id string, id string) (*entity.EntityProductGroup, error)
	SearchByName(store_id string, name string) (*entity.EntityProductGroup, error)
	Create(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error)
	Update(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error)
	Delete(store_id string, id string) error
}

type IUseCaseProductGroup interface {
	GetAll(store_id string) ([]*entity.EntityProductGroup, error)
	GetByID(store_id string, id string) (*entity.EntityProductGroup, error)
	SearchByName(store_id string, name string) (*entity.EntityProductGroup, error)
	Create(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error)
	Update(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error)
	Delete(store_id string, id string) error
}
