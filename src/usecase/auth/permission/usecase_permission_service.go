package usecase_permission

import entity "main/entity/auth"

type UseCasePermission struct {
	repo IRepositoryPermission
}

func NewService(repository IRepositoryPermission) *UseCasePermission {
	return &UseCasePermission{repo: repository}
}

func (s *UseCasePermission) Create(permission entity.EntityPermission) (*entity.EntityPermission, error) {
	err := permission.Validate()
	if err != nil {
		return nil, err
	}
	entity.CreatePermission(&permission)

	return s.repo.CreatePermission(permission)
}

func (s *UseCasePermission) GetAll() (*[]entity.EntityPermission, error) {
	return s.repo.GetAll()
}

func (s *UseCasePermission) GetBySlug(slug string) (*entity.EntityPermission, error) {
	return s.repo.GetBySlug(slug)
}
