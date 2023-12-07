package usecase_product_promotion_product

import entity "main/entity/product"

type UseCaseProductPromotionProduct struct {
	repo IRepositoryProductPromotionProduct
}

func NewService(repo IRepositoryProductPromotionProduct) *UseCaseProductPromotionProduct {
	return &UseCaseProductPromotionProduct{repo}
}

func (s *UseCaseProductPromotionProduct) GetAll(store_id string, promotion_id string) (*[]entity.EntityProductPromotionProduct, error) {
	return s.repo.GetAll(store_id, promotion_id)
}

func (s *UseCaseProductPromotionProduct) GetByID(store_id string, promotion_id string, id string) (*entity.EntityProductPromotionProduct, error) {
	return s.repo.GetByID(store_id, promotion_id, id)
}

func (s *UseCaseProductPromotionProduct) Create(productPromotionProduct entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error) {
	err := productPromotionProduct.Validate()
	if err != nil {
		return nil, err
	}
	entity.CreateProductPromotionProduct(&productPromotionProduct)
	return s.repo.Create(productPromotionProduct)
}

func (s *UseCaseProductPromotionProduct) Update(productPromotionProduct entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error) {
	err := productPromotionProduct.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(productPromotionProduct)
}

func (s *UseCaseProductPromotionProduct) Delete(store_id string, promotion_id string, id string) error {
	return s.repo.Delete(store_id, promotion_id, id)
}
