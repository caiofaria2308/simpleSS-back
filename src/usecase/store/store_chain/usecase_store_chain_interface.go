package usecase_store_chain

import (
	entity "main/entity/store"
)

type IRepositoryStoreChain interface {
	GetAll() (*[]entity.EntityStoreChain, error)
	GetByID(id string) (*entity.EntityStoreChain, error)
	Create(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error)
	Update(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error)
	Delete(id string) error
	SearchByName(name string) (*[]entity.EntityStoreChain, error)
}

type IUseCaseStoreChain interface {
	GetAll() (*[]entity.EntityStoreChain, error)
	GetByID(id string) (*entity.EntityStoreChain, error)
	Create(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error)
	Update(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error)
	Delete(id string) error
	SearchByName(name string) (*[]entity.EntityStoreChain, error)
}
