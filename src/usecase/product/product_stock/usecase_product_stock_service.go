package usecase_product_stock

import entity "main/entity/product"

type UseCaseProductStock struct {
	repo IRepositoryProductStock
}

func NewService(repo IRepositoryProductStock) *UseCaseProductStock {
	return &UseCaseProductStock{repo}
}

func (s *UseCaseProductStock) GetAll(store_id string) (*[]entity.EntityProductStock, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseProductStock) GetByID(store_id string, id string) (*entity.EntityProductStock, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseProductStock) FilterByProductID(store_id string, product_id string) (*[]entity.EntityProductStock, error) {
	return s.repo.FilterByProductID(store_id, product_id)
}

func (s *UseCaseProductStock) Create(productStock entity.EntityProductStock) (*entity.EntityProductStock, error) {
	err := productStock.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(productStock)
}

func (s *UseCaseProductStock) Update(productStock entity.EntityProductStock) (*entity.EntityProductStock, error) {
	err := productStock.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(productStock)
}

func (s *UseCaseProductStock) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}
