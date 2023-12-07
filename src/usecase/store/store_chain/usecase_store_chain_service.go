package usecase_store_chain

import (
	entity "main/entity/store"
)

type UseCaseStoreChain struct {
	repo IRepositoryStoreChain
}

func NewService(repo IRepositoryStoreChain) *UseCaseStoreChain {
	return &UseCaseStoreChain{repo: repo}
}

func (uc *UseCaseStoreChain) GetAll() (*[]entity.EntityStoreChain, error) {
	return uc.repo.GetAll()
}

func (uc *UseCaseStoreChain) GetByID(id string) (*entity.EntityStoreChain, error) {
	return uc.repo.GetByID(id)
}

func (uc *UseCaseStoreChain) Create(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error) {
	err := storeChain.Validate()
	err = entity.CreateStoreChain(&storeChain)
	if err != nil {
		return nil, err
	}
	return uc.repo.Create(storeChain)
}

func (uc *UseCaseStoreChain) Update(storeChain entity.EntityStoreChain) (*entity.EntityStoreChain, error) {
	err := storeChain.Validate()
	err = entity.UpdateStoreChain(&storeChain)
	if err != nil {
		return nil, err
	}
	return uc.repo.Update(storeChain)
}

func (uc *UseCaseStoreChain) Delete(id string) error {
	return uc.repo.Delete(id)
}

func (uc *UseCaseStoreChain) SearchByName(name string) (*[]entity.EntityStoreChain, error) {
	return uc.repo.SearchByName(name)
}
