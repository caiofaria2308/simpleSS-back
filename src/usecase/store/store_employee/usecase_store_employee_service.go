package usecase_store_employee

import entity "main/entity/store"

type UseCaseStoreEmployee struct {
	repo IRepositoryStoreEmployee
}

func NewService(repo IRepositoryStoreEmployee) *UseCaseStoreEmployee {
	return &UseCaseStoreEmployee{repo: repo}
}

func (s *UseCaseStoreEmployee) GetAll(store_id string) (*[]entity.EntityStoreEmployee, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseStoreEmployee) GetByID(store_id string, id string) (*entity.EntityStoreEmployee, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseStoreEmployee) Create(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error) {
	err := storeemployee.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(storeemployee)
}

func (s *UseCaseStoreEmployee) Update(storeemployee entity.EntityStoreEmployee) (*entity.EntityStoreEmployee, error) {
	err := storeemployee.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(storeemployee)
}

func (s *UseCaseStoreEmployee) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}

func (s *UseCaseStoreEmployee) SearchByUserName(store_id string, name string) (storeEmployees *[]entity.EntityStoreEmployee, err error) {
	return s.repo.SearchByUserName(store_id, name)
}

func (s *UseCaseStoreEmployee) SearchByEmail(store_id string, email string) (storeEmployees *[]entity.EntityStoreEmployee, err error) {
	return s.repo.SearchByEmail(store_id, email)
}
