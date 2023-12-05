package usecase_provider

import entity "main/entity/provider"

type UseCaseProvider struct {
	repo IRepositoryProvider
}

func NewService(repo IRepositoryProvider) *UseCaseProvider {
	return &UseCaseProvider{repo: repo}
}

func (s *UseCaseProvider) GetAll(store_id string) ([]*entity.EntityProvider, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseProvider) GetByID(store_id string, id string) (*entity.EntityProvider, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseProvider) Create(provider entity.EntityProvider) (*entity.EntityProvider, error) {
	err := provider.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(provider)
}

func (s *UseCaseProvider) Update(provider entity.EntityProvider) (*entity.EntityProvider, error) {
	err := provider.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(provider)
}

func (s *UseCaseProvider) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}

func (s *UseCaseProvider) SearchByCNPJ(store_id string, cnpj string) (providers []*entity.EntityProvider, err error) {
	return s.repo.SearchByCNPJ(store_id, cnpj)
}

func (s *UseCaseProvider) SearchByName(store_id string, name string) (providers []*entity.EntityProvider, err error) {
	return s.repo.SearchByName(store_id, name)
}
