package entity

import (
	entity_product "main/entity/product"
	"main/utils"
)

type EntityPurchaseOrderProduct struct {
	ID                 string                       `json:"id" gorm:"primaryKey"`
	OrderID            string                       `json:"order_id" validate:"required"`
	Order              EntityPurchaseOrder          `gorm:"not null"`
	ProductID          string                       `json:"product_id" validate:"required"`
	Product            entity_product.EntityProduct `gorm:"not null"`
	FullPrice          float64                      `json:"price" gorm:"not null" validate:"required"`
	DiscountPercentage float64                      `json:"discount_percentage" gorm:"not null; default: 0.0"`
	Quantity           float64                      `json:"quantity" gorm:"not null" validate:"required"`
}

func CreatePurchaseOrderProduct(purchaseOrderProductParams EntityPurchaseOrderProduct) (*EntityPurchaseOrderProduct, error) {
	p := &EntityPurchaseOrderProduct{
		ID:                 utils.GenerateID(),
		Order:              purchaseOrderProductParams.Order,
		Product:            purchaseOrderProductParams.Product,
		FullPrice:          purchaseOrderProductParams.FullPrice,
		DiscountPercentage: purchaseOrderProductParams.DiscountPercentage,
		Quantity:           purchaseOrderProductParams.Quantity,
	}
	return p, nil
}

func (p *EntityPurchaseOrderProduct) Validate() error {
	return validate.Struct(p)
}
