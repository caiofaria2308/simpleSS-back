package usecase_product_promotion_product

import entity "main/entity/product"

type IRepositoryProductPromotionProduct interface {
	GetAll(store_id string, promotion_id string) (*[]entity.EntityProductPromotionProduct, error)
	GetByID(store_id string, promotion_id string, id string) (*entity.EntityProductPromotionProduct, error)
	Create(entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error)
	Update(entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error)
	Delete(store_id string, promotion_id string, id string) error
}

type IUseCaseProdutPromotionProduct interface {
	GetAll(store_id string, promotion_id string) (*[]entity.EntityProductPromotionProduct, error)
	GetByID(store_id string, promotion_id string, id string) (*entity.EntityProductPromotionProduct, error)
	Create(entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error)
	Update(entity.EntityProductPromotionProduct) (*entity.EntityProductPromotionProduct, error)
	Delete(store_id string, promotion_id string, id string) error
}
