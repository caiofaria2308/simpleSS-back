package entity

import "main/utils"

type EntityProductStock struct {
	ID      string        `json:"id" gorm:"primaryKey"`
	Product EntityProduct `json:"product_id" gorm:"foreignKey:ProductID; not null" validate:"required"`
	Stock   float64       `json:"stock" gorm:"not null" validate:"required"`
}

func CreateProductStock(productStock EntityProductStock) (*EntityProductStock, error) {
	p := &EntityProductStock{
		ID:      utils.GenerateID(),
		Product: productStock.Product,
		Stock:   productStock.Stock,
	}
	return p, nil
}

func (p *EntityProductStock) Validate() error {
	return validate.Struct(p)
}
