package entity

import "main/utils"

type EntityProductPromotionProduct struct {
	ID        string                 `json:"id" gorm:"primaryKey"`
	Product   EntityProduct          `json:"product_id" gorm:"foreignKey:ProductID; not null; index:idx_unique" validate:"required"`
	Promotion EntityProductPromotion `json:"promotion_id" gorm:"foreignKey:PromotionID; not null; index:idx_unique" validate:"required"`
}

func CreateProductPromotionProduct(productPromotionProductParams EntityProductPromotionProduct) (*EntityProductPromotionProduct, error) {
	p := &EntityProductPromotionProduct{
		ID:        utils.GenerateID(),
		Product:   productPromotionProductParams.Product,
		Promotion: productPromotionProductParams.Promotion,
	}
	return p, nil
}

func (p *EntityProductPromotionProduct) Validate() error {
	return validate.Struct(p)
}
