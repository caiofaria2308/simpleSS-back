package entity

import (
	entity_store "main/entity/store"
	"main/utils"
)

type EntityProduct struct {
	ID          string                   `json:"id" gorm:"primaryKey"`
	Store       entity_store.EntityStore `json:"store_id" gorm:"foreignKey:StoreID; not null" validate:"required"`
	Group       EntityProductGroup       `json:"group_id" gorm:"foreignKey:GroupID; not null" validate:"required"`
	BarCode     string                   `json:"bar_code" gorm:"unique; not null" validate:"required"`
	Name        string                   `json:"name" gorm:"not null" validate:"required"`
	Price       float64                  `json:"price" gorm:"not null" validate:"required"`
	MaxDiscount float64                  `json:"max_discount" gorm:"not null; default:0.0" validate:"required"`
}

func CreateProduct(productParams EntityProduct) (*EntityProduct, error) {
	p := &EntityProduct{
		ID:          utils.GenerateID(),
		Store:       productParams.Store,
		Group:       productParams.Group,
		BarCode:     productParams.BarCode,
		Name:        productParams.Name,
		Price:       productParams.Price,
		MaxDiscount: productParams.MaxDiscount,
	}
	return p, nil
}

func (p *EntityProduct) Validate() error {
	return validate.Struct(p)
}
