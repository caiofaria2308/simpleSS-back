package usecase_product_promotion

import entity "main/entity/product"

type IRepositoryProductPromotion interface {
	GetAll(store_id string) (*[]entity.EntityProductPromotion, error)
	GetByID(store_id string, id string) (*entity.EntityProductPromotion, error)
	FilterByPromotion(store_id string, promotion_id string) (*[]entity.EntityProductPromotion, error)
	FilterByProduct(store_id string, product_id string) (*[]entity.EntityProductPromotion, error)
	Create(entity.EntityProductPromotion) (*entity.EntityProductPromotion, error)
	Update(entity.EntityProductPromotion) (*entity.EntityProductPromotion, error)
	Delete(store_id string, id string) error
}

type IUseCaseProdutPromotion interface {
	GetAll(store_id string) (*[]entity.EntityProductPromotion, error)
	GetByID(store_id string, id string) (*entity.EntityProductPromotion, error)
	FilterByPromotion(store_id string, promotion_id string) (*[]entity.EntityProductPromotion, error)
	FilterByProduct(store_id string, product_id string) (*[]entity.EntityProductPromotion, error)
	Create(entity.EntityProductPromotion) (*entity.EntityProductPromotion, error)
	Update(entity.EntityProductPromotion) (*entity.EntityProductPromotion, error)
	Delete(store_id string, id string) error
}
