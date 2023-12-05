package usecase_store_employee

import entity "main/entity/store"

type IRepositoryStoreEmployee interface {
	Create(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error)
	GetAll(store_id string) (*[]entity.EntityStoreEmployee, error)
	GetByID(store_id string, id string) (*entity.EntityStoreEmployee, error)
	Update(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error)
	Delete(store_id string, id string) error
	SearchByUserName(store_id string, name string) (storeEmployees *[]entity.EntityStoreEmployee, err error)
	SearchByEmail(store_id string, email string) (storeEmployees *[]entity.EntityStoreEmployee, err error)
}

type IUseCaseStoreEmployee interface {
	GetAll(store_id string) (*[]entity.EntityStoreEmployee, error)
	GetByID(store_id string, id string) (*entity.EntityStoreEmployee, error)
	Create(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error)
	Update(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error)
	Delete(store_id string, id string) error
	SearchByUserName(store_id string, name string) (storeEmployees *[]entity.EntityStoreEmployee, err error)
	SearchByEmail(store_id string, email string) (storeEmployees *[]entity.EntityStoreEmployee, err error)
}
