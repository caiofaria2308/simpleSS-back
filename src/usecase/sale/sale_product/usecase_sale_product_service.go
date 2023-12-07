package usecase_sale_product

import entity "main/entity/sale"

type UseCaseSaleProduct struct {
	repo IRepositorySaleProduct
}

func NewService(repo IRepositorySaleProduct) *UseCaseSaleProduct {
	return &UseCaseSaleProduct{repo: repo}
}

func (s *UseCaseSaleProduct) GetAll(store_id string, sale_id string) ([]*entity.EntitySaleProduct, error) {
	return s.repo.GetAll(store_id, sale_id)
}

func (s *UseCaseSaleProduct) GetByID(store_id string, sale_id string, id string) (*entity.EntitySaleProduct, error) {
	return s.repo.GetByID(store_id, sale_id, id)
}

func (s *UseCaseSaleProduct) Create(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error) {
	err := saleProduct.Validate()
	err = entity.CreateSaleProduct(&saleProduct)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(saleProduct)
}

func (s *UseCaseSaleProduct) Update(saleProduct entity.EntitySaleProduct) (*entity.EntitySaleProduct, error) {
	err := saleProduct.Validate()
	err = entity.UpdateSaleProduct(&saleProduct)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(saleProduct)
}

func (s *UseCaseSaleProduct) Delete(store_id string, sale_id string, id string) error {
	return s.repo.Delete(store_id, sale_id, id)
}
