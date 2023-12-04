package entity

import (
	entity_product "main/entity/product"
	"main/utils"
)

type EntitySaleProduct struct {
	ID        string                                `json:"id" gorm:"primaryKey"`
	Sale      EntitySale                            `json:"sale_id" gorm:"foreignKey:SaleID; not null" validate:"required"`
	Product   entity_product.EntityProduct          `json:"product_id" gorm:"foreignKey:ProductID; not null" validate:"required"`
	Promotion entity_product.EntityProductPromotion `json:"promotion_id" gorm:"foreignKey:PromotionID"`
	Quantity  float64                               `json:"quantity" gorm:"not null" validate:"required"`
	FullPrice float64                               `json:"full_price" gorm:"not null"`
	Discount  float64                               `json:"discount" gorm:"not null; default: 0.0"`
	Price     float64                               `json:"price" gorm:"not null"`
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
