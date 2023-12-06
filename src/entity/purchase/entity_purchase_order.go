package entity

import (
	"errors"
	entity_provider "main/entity/provider"
	entity_store "main/entity/store"
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityPurchaseOrder struct {
	ID         string                           `json:"id" gorm:"primaryKey"`
	StoreID    string                           `json:"store_id" gorm:"index:idx_purchaseorder_unique"`
	Store      entity_store.EntityStore         `gorm:"not null;"`
	Number     string                           `json:"number" gorm:"not null; index:idx_purchaseorder_unique"`
	Date       time.Time                        `json:"date" gorm:"not null" validate:"required"`
	ProviderID string                           `json:"provider_id" validate:"required"`
	Provider   entity_provider.EntityProvider   `gorm:"not null;"`
	EmployeeID string                           `json:"employee_id" validate:"required"`
	Employee   entity_store.EntityStoreEmployee `gorm:"not null;"`
}

func CreatePurchaseOrder(purchaseOrderParams EntityPurchaseOrder) (*EntityPurchaseOrder, error) {
	if entity_store.VerifyStoreEmployeeIsActive(&purchaseOrderParams.Employee) == false {
		return nil, errors.New("Employee is not active")
	}
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
