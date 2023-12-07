package entity

import "main/utils"

type EntityProductStock struct {
	ID        string        `json:"id" gorm:"primaryKey"`
	ProductID string        `json:"product_id" validate:"required"`
	Product   EntityProduct `gorm:"not null; unique"`
	Stock     float64       `json:"stock" gorm:"not null" validate:"required"`
}

func CreateProductStock(productStockParams *EntityProductStock) error {
	productStockParams.ID = utils.GenerateID()
	return nil
}

func (p *EntityProductStock) Validate() error {
	return validate.Struct(p)
}
