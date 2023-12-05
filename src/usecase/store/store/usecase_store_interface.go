package usecase_store

import (
	entity "main/entity/store"
)

type IRepositoryStore interface {
	GetAll(chain_id string) (*[]entity.EntityStore, error)
	GetByID(chain_id string, id string) (*entity.EntityStore, error)
	Create(store entity.EntityStore) (*entity.EntityStore, error)
	Update(store entity.EntityStore) (*entity.EntityStore, error)
	Delete(chain_id string, id string) error
	FilterByCNPJ(chain_id string, cnpj string) (store *[]entity.EntityStore, err error)
	Search(chain_id string, search string) (store *[]entity.EntityStore, err error)
}

type IUseCaseStore interface {
	GetAll(chain_id string) (*[]entity.EntityStore, error)
	GetByID(chain_id string, id string) (*entity.EntityStore, error)
	Create(store entity.EntityStore) (*entity.EntityStore, error)
	Update(store entity.EntityStore) (*entity.EntityStore, error)
	Delete(chain_id string, id string) error
	FilterByCNPJ(chain_id string, cnpj string) (store *[]entity.EntityStore, err error)
	Search(chain_id string, search string) (store *[]entity.EntityStore, err error)
}
