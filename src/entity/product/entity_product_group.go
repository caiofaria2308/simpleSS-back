package entity

import (
	entity_store "main/entity/store"
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityProductGroup struct {
	ID        string                   `json:"id" gorm:"primaryKey"`
	StoreID   string                   `json:"store_id" gorm:"foreignKey:Store; not null; index:idx_product_group_unique"`
	Store     entity_store.EntityStore `gorm:"not null;"`
	Name      string                   `json:"name" gorm:"not null; index:idx_product_group_unique" validate:"required"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

func CreateProductGroup(productGroupParams *EntityProductGroup) error {
	productGroupParams.ID = utils.GenerateID()
	productGroupParams.CreatedAt = time.Now()
	productGroupParams.UpdatedAt = time.Now()
	return nil
}

func UpdateProductGroup(productGroupParams *EntityProductGroup) error {
	productGroupParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityProductGroup) Validate() error {
	return validate.Struct(p)
}
