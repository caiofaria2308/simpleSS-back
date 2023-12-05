package usecase_product_promotion

import entity "main/entity/product"

type UseCaseProductPromotion struct {
	repo IRepositoryProductPromotion
}

func NewService(repo IRepositoryProductPromotion) *UseCaseProductPromotion {
	return &UseCaseProductPromotion{repo}
}

func (s *UseCaseProductPromotion) GetAll(store_id string) (*[]entity.EntityProductPromotion, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseProductPromotion) GetByID(store_id string, id string) (*entity.EntityProductPromotion, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseProductPromotion) Create(productPromotion entity.EntityProductPromotion) (*entity.EntityProductPromotion, error) {
	err := productPromotion.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(productPromotion)
}

func (s *UseCaseProductPromotion) Update(productPromotion entity.EntityProductPromotion) (*entity.EntityProductPromotion, error) {
	err := productPromotion.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(productPromotion)
}

func (s *UseCaseProductPromotion) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}
