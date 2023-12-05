package usecase_client

import entity "main/entity/client"

type UseCaseClient struct {
	repo IRepositoryClient
}

func NewService(repo IRepositoryClient) *UseCaseClient {
	return &UseCaseClient{repo: repo}
}

func (s *UseCaseClient) GetAll(store_id string) ([]*entity.EntityClient, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseClient) GetByID(store_id string, id string) (*entity.EntityClient, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseClient) Create(client entity.EntityClient) (*entity.EntityClient, error) {
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(client)
}

func (s *UseCaseClient) Update(client entity.EntityClient) (*entity.EntityClient, error) {
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(client)
}

func (s *UseCaseClient) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}

func (s *UseCaseClient) GetByCPF(store_id string, id string) (*entity.EntityClient, error) {
	return s.repo.GetByCPF(store_id, id)
}

func (s *UseCaseClient) Search(store_id string, search string) ([]*entity.EntityClient, error) {
	return s.repo.Search(store_id, search)
}
