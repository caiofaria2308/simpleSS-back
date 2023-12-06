package usecase_client

import entity "main/entity/client"

type IRepositoryClient interface {
	GetAll(store_id string) ([]*entity.EntityClient, error)
	GetByID(store_id string, id string) (*entity.EntityClient, error)
	GetByCPF(store_id string, id string) (*entity.EntityClient, error)
	Search(store_id string, search string) ([]*entity.EntityClient, error)
	Create(client entity.EntityClient) (*entity.EntityClient, error)
	Update(client entity.EntityClient) (*entity.EntityClient, error)
	Delete(store_id string, id string) error
}

type IUseCaseClient interface {
	GetAll(store_id string) ([]*entity.EntityClient, error)
	GetByID(store_id string, id string) (*entity.EntityClient, error)
	GetByCPF(store_id string, id string) (*entity.EntityClient, error)
	Search(store_id string, search string) ([]*entity.EntityClient, error)
	Create(client entity.EntityClient) (*entity.EntityClient, error)
	Update(client entity.EntityClient) (*entity.EntityClient, error)
	Delete(store_id string, id string) error
}
