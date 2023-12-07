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
	ID         string                           `json:"id" gorm:"primaryKey"`
	StoreID    string                           `json:"store_id" gorm:"index:idx_purchaseorder_unique"`
	Store      entity_store.EntityStore         `gorm:"not null;"`
	Number     string                           `json:"number" gorm:"not null; index:idx_purchaseorder_unique"`
	Date       time.Time                        `json:"date" gorm:"not null" validate:"required"`
	ProviderID string                           `json:"provider_id" validate:"required"`
	Provider   entity_provider.EntityProvider   `gorm:"not null;"`
	EmployeeID string                           `json:"employee_id" validate:"required"`
	Employee   entity_store.EntityStoreEmployee `gorm:"not null;"`
	CreatedAt  time.Time                        `json:"created_at"`
	UpdatedAt  time.Time                        `json:"updated_at"`
}

func CreatePurchaseOrder(purchaseOrderParams *EntityPurchaseOrder) error {
	purchaseOrderParams.ID = utils.GenerateID()
	purchaseOrderParams.Date = time.Now()
	purchaseOrderParams.Number = utils.GenerateNumber(time.Now())
	purchaseOrderParams.CreatedAt = time.Now()
	purchaseOrderParams.UpdatedAt = time.Now()
	return nil
}

func UpdatePurchaseOrder(purchaseOrderParams *EntityPurchaseOrder) error {
	purchaseOrderParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityPurchaseOrder) Validate() error {
	return validate.Struct(p)
}
