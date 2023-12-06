package entity

import (
	entity_product "main/entity/product"
	"main/utils"
)

type EntitySaleProduct struct {
	ID          string                       `json:"id" gorm:"primaryKey"`
	SaleID      string                       `json:"sale_id" validate:"required"`
	Sale        EntitySale                   `gorm:"not null"`
	ProductID   string                       `json:"product_id" validate:"required"`
	Product     entity_product.EntityProduct `gorm:"not null"`
	PromotionID string                       `json:"promotion_id"`
	Promotion   entity_product.EntityProductPromotion
	Quantity    float64 `json:"quantity" gorm:"not null" validate:"required"`
	FullPrice   float64 `json:"full_price" gorm:"not null"`
	Discount    float64 `json:"discount" gorm:"not null; default: 0.0"`
	Price       float64 `json:"price" gorm:"not null"`
}

func CalculatePrice(saleProduct *EntitySaleProduct) {
	saleProduct.FullPrice = saleProduct.Product.Price * saleProduct.Quantity
	if saleProduct.Promotion.ID != "" {
		saleProduct.Price = saleProduct.FullPrice - (saleProduct.FullPrice * saleProduct.Promotion.Percentage)
	} else if saleProduct.Discount > 0 {
		saleProduct.Price = saleProduct.FullPrice - (saleProduct.FullPrice * saleProduct.Discount)
	} else {
		saleProduct.Price = saleProduct.FullPrice
	}
	return
}

func CreateSaleProduct(saleProductParams EntitySaleProduct) (*EntitySaleProduct, error) {
	p := &EntitySaleProduct{
		ID:        utils.GenerateID(),
		Sale:      saleProductParams.Sale,
		Product:   saleProductParams.Product,
		Promotion: saleProductParams.Promotion,
		Quantity:  saleProductParams.Quantity,
	}
	CalculatePrice(p)
	return p, nil
}

func (p *EntitySaleProduct) Validate() error {
	return validate.Struct(p)
}
