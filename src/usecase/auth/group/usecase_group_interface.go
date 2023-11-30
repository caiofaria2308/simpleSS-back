package usecase_group

import (
	entity "main/entity/auth"
)

type IRepositoryGroup interface {
	GetAll() (*[]entity.EntityGroup, error)
	GetByID(id string) (*entity.EntityGroup, error)
	Create(group entity.EntityGroup) (*entity.EntityGroup, error)
	Update(group entity.EntityGroup) (*entity.EntityGroup, error)
	Delete(id string) error
}

type IUseCaseGroup interface {
	GetAll() (*[]entity.EntityGroup, error)
	GetByID(id string) (*entity.EntityGroup, error)
	Create(group entity.EntityGroup) (*entity.EntityGroup, error)
	Update(group entity.EntityGroup) (*entity.EntityGroup, error)
	Delete(id string) error
}
