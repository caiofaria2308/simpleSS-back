package usecase_provider

import entity "main/entity/provider"

type IRepositoryProvider interface {
	GetAll(store_id string) ([]*entity.EntityProvider, error)
	GetByID(store_id string, id string) (*entity.EntityProvider, error)
	Create(provider entity.EntityProvider) (*entity.EntityProvider, error)
	Update(provider entity.EntityProvider) (*entity.EntityProvider, error)
	Delete(store_id string, id string) error
	SearchByCNPJ(store_id string, cnpj string) (providers []*entity.EntityProvider, err error)
	SearchByName(store_id string, name string) (providers []*entity.EntityProvider, err error)
}

type IUseCaseProvider interface {
	GetAll(store_id string) ([]*entity.EntityProvider, error)
	GetByID(store_id string, id string) (*entity.EntityProvider, error)
	Create(provider entity.EntityProvider) (*entity.EntityProvider, error)
	Update(provider entity.EntityProvider) (*entity.EntityProvider, error)
	Delete(store_id string, id string) error
	SearchByCNPJ(store_id string, cnpj string) (providers []*entity.EntityProvider, err error)
	SearchByName(store_id string, name string) (providers []*entity.EntityProvider, err error)
}
