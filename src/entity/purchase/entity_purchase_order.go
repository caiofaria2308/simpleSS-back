package entity

import (
	entity_provider "main/entity/provider"
	entity_store "main/entity/store"
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityPurchaseOrder struct {
	ID       string                           `json:"id" gorm:"primaryKey"`
	Store    entity_store.EntityStore         `json:"store_id" gorm:"foreignKey:StoreID; not null"`
	Number   string                           `json:"number" gorm:"not null"`
	Date     time.Time                        `json:"date" gorm:"not null" validate:"required"`
	Provider entity_provider.EntityProvider   `json:"provider_id" gorm:"foreignKey: ProviderID" validate:"required"`
	Employee entity_store.EntityStoreEmployee `json:"employee_id" gorm:"foreignKey:EmployeeID; not null" validate:"required"`
}

func CreatePurchaseOrder(purchaseOrderParams EntityPurchaseOrder) (*EntityPurchaseOrder, error) {
	p := &EntityPurchaseOrder{
		ID:       utils.GenerateID(),
		Number:   utils.GenerateNumber(purchaseOrderParams.Date),
		Store:    purchaseOrderParams.Store,
		Date:     purchaseOrderParams.Date,
		Provider: purchaseOrderParams.Provider,
	}
	return p, nil
}

func (p *EntityPurchaseOrder) Validate() error {
	return validate.Struct(p)
}
