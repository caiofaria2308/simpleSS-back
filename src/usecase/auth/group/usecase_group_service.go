package usecase_group

import entity "main/entity/auth"

type UseCaseGroup struct {
	repo IRepositoryGroup
}

func NewService(repo IRepositoryGroup) *UseCaseGroup {
	return &UseCaseGroup{repo: repo}
}

func (s *UseCaseGroup) GetAll() (*[]entity.EntityGroup, error) {
	return s.repo.GetAll()
}

func (s *UseCaseGroup) GetByID(id string) (*entity.EntityGroup, error) {
	return s.repo.GetByID(id)
}

func (s *UseCaseGroup) Create(group entity.EntityGroup) (*entity.EntityGroup, error) {
	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(group)
}

func (s *UseCaseGroup) Update(group entity.EntityGroup) (*entity.EntityGroup, error) {
	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(group)
}

func (s *UseCaseGroup) Delete(id string) error {
	return s.repo.Delete(id)
}
