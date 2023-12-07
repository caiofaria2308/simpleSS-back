package usecase_product_group

import entity "main/entity/product"

type UseCaseProductGroup struct {
	repo IRepositoryProductGroup
}

func NewService(repo IRepositoryProductGroup) *UseCaseProductGroup {
	return &UseCaseProductGroup{repo: repo}
}

func (u *UseCaseProductGroup) GetAll(store_id string) ([]*entity.EntityProductGroup, error) {
	return u.repo.GetAll(store_id)
}

func (u *UseCaseProductGroup) GetByID(store_id string, id string) (*entity.EntityProductGroup, error) {
	return u.repo.GetByID(store_id, id)
}

func (u *UseCaseProductGroup) SearchByName(store_id string, name string) (*entity.EntityProductGroup, error) {
	return u.repo.SearchByName(store_id, name)
}

func (u *UseCaseProductGroup) Create(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error) {
	err := productGroup.Validate()
	if err != nil {
		return nil, err
	}
	entity.CreateProductGroup(&productGroup)
	return u.repo.Create(productGroup)
}

func (u *UseCaseProductGroup) Update(productGroup entity.EntityProductGroup) (*entity.EntityProductGroup, error) {
	err := productGroup.Validate()
	if err != nil {
		return nil, err
	}
	entity.UpdateProductGroup(&productGroup)
	return u.repo.Update(productGroup)
}

func (u *UseCaseProductGroup) Delete(store_id string, id string) error {
	return u.repo.Delete(store_id, id)
}
