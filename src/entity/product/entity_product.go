package entity

import (
	entity_store "main/entity/store"
	"main/utils"
	"time"
)

type EntityProduct struct {
	ID          string                   `json:"id" gorm:"primaryKey"`
	StoreID     string                   `json:"store_id" gorm:"idx_product_unique" validate:"required"`
	Store       entity_store.EntityStore `json:"store" gorm:"not null;"`
	GroupID     string                   `json:"group_id" validate:"required"`
	Group       EntityProductGroup       `json:"group" gorm:"not null;"`
	BarCode     string                   `json:"bar_code" gorm:"idx_product_unique; not null" validate:"required"`
	Name        string                   `json:"name" gorm:"not null" validate:"required"`
	Price       float64                  `json:"price" gorm:"not null" validate:"required"`
	MaxDiscount float64                  `json:"max_discount" gorm:"not null; default:0.0" validate:"required"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

func CreateProduct(productParams *EntityProduct) error {
	productParams.ID = utils.GenerateID()
	productParams.CreatedAt = time.Now()
	productParams.UpdatedAt = time.Now()
	return nil
}

func UpdateProduct(productParams *EntityProduct) error {
	productParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityProduct) Validate() error {
	return validate.Struct(p)
}
