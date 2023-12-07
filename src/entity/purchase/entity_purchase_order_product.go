package entity

import (
	entity_product "main/entity/product"
	"main/utils"
	"time"
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
	CreatedAt          time.Time                    `json:"created_at"`
	UpdatedAt          time.Time                    `json:"updated_at"`
}

func CreatePurchaseOrderProduct(purchaseOrderProductParams *EntityPurchaseOrderProduct) error {
	purchaseOrderProductParams.ID = utils.GenerateID()
	purchaseOrderProductParams.CreatedAt = time.Now()
	purchaseOrderProductParams.UpdatedAt = time.Now()
	return nil
}

func UpdatePurchaseOrderProduct(purchaseOrderProductParams *EntityPurchaseOrderProduct) error {
	purchaseOrderProductParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityPurchaseOrderProduct) Validate() error {
	return validate.Struct(p)
}
