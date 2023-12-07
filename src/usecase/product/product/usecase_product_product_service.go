package usecase_product_product

import (
	entity "main/entity/product"
)

type UseCaseProductProduct struct {
	repo IRepositoryProductProduct
}

func NewService(repo IRepositoryProductProduct) *UseCaseProductProduct {
	return &UseCaseProductProduct{repo: repo}
}

func (s *UseCaseProductProduct) GetAll(store_id string) ([]*entity.EntityProduct, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseProductProduct) GetByID(store_id string, id string) (*entity.EntityProduct, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseProductProduct) SearchByName(store_id string, name string) (*entity.EntityProduct, error) {
	return s.repo.SearchByName(store_id, name)
}

func (s *UseCaseProductProduct) SearchByBarcode(store_id string, barcode string) (*entity.EntityProduct, error) {
	return s.repo.SearchByBarcode(store_id, barcode)
}

func (s *UseCaseProductProduct) SearchByGroup(store_id string, group_id string) (*entity.EntityProduct, error) {
	return s.repo.SearchByGroup(store_id, group_id)
}

func (s *UseCaseProductProduct) Create(product entity.EntityProduct) (*entity.EntityProduct, error) {
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	entity.CreateProduct(&product)
	return s.repo.Create(product)
}

func (s *UseCaseProductProduct) Update(product entity.EntityProduct) (*entity.EntityProduct, error) {
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	entity.UpdateProduct(&product)
	return s.repo.Update(product)
}

func (s *UseCaseProductProduct) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}
