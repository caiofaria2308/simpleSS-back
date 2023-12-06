package entity

import "main/utils"

type EntityProductPromotionProduct struct {
	ID          string                 `json:"id" gorm:"primaryKey"`
	ProductID   string                 `json:"product_id" gorm:"index:idx_product_promotion_product_unique" validate:"required"`
	Product     EntityProduct          `gorm:"not null" validate:"required"`
	PromotionID string                 `json:"promotion_id" gorm:"index:idx_product_promotion_product_unique" validate:"required"`
	Promotion   EntityProductPromotion `gorm:"not null" validate:"required"`
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
