package entity

import (
	entity_store "main/entity/store"
	"main/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityProductGroup struct {
	ID      string                   `json:"id" gorm:"primaryKey"`
	StoreID string                   `json:"store_id" gorm:"foreignKey:Store; not null; index:idx_product_group_unique"`
	Store   entity_store.EntityStore `gorm:"not null;"`
	Name    string                   `json:"name" gorm:"not null; index:idx_product_group_unique" validate:"required"`
}

func CreateProductGroup(productGroupParams EntityProductGroup) (*EntityProductGroup, error) {
	p := &EntityProductGroup{
		ID:    utils.GenerateID(),
		Store: productGroupParams.Store,
		Name:  productGroupParams.Name,
	}
	return p, nil
}

func (p *EntityProductGroup) Validate() error {
	return validate.Struct(p)
}
