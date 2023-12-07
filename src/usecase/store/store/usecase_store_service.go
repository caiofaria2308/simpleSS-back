package usecase_store

import entity "main/entity/store"

type UseCaseStore struct {
	repo IRepositoryStore
}

func NewService(repo IRepositoryStore) *UseCaseStore {
	return &UseCaseStore{repo: repo}
}

func (s *UseCaseStore) GetAll(chain_id string) (*[]entity.EntityStore, error) {
	return s.repo.GetAll(chain_id)
}

func (s *UseCaseStore) GetByID(chain_id string, id string) (*entity.EntityStore, error) {
	return s.repo.GetByID(chain_id, id)
}

func (s *UseCaseStore) Create(store entity.EntityStore) (*entity.EntityStore, error) {
	err := store.Validate()
	err = entity.CreateStore(&store)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(store)
}

func (s *UseCaseStore) Update(store entity.EntityStore) (*entity.EntityStore, error) {
	err := store.Validate()
	err = entity.UpdateStore(&store)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(store)
}

func (s *UseCaseStore) Delete(chain_id string, id string) error {
	return s.repo.Delete(chain_id, id)
}

func (s *UseCaseStore) FilterByCNPJ(chain_id string, cnpj string) (store *[]entity.EntityStore, err error) {
	return s.repo.FilterByCNPJ(chain_id, cnpj)
}

func (s *UseCaseStore) Search(chain_id string, search string) (store *[]entity.EntityStore, err error) {
	return s.repo.Search(chain_id, search)
}
