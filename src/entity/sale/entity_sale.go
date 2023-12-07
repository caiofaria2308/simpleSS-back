package entity

import (
	entity_client "main/entity/client"
	entity_store "main/entity/store"
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntitySale struct {
	ID         string                   `json:"id" gorm:"primaryKey"`
	Date       time.Time                `json:"date" gorm:"not null"`
	Number     string                   `json:"number" gorm:"index:idx_sale_unique; not null"`
	StoreID    string                   `json:"store_id" gorm:"index:idx_sale_unique;" validate:"required"`
	Store      entity_store.EntityStore `gorm:"not null;"`
	ClientID   string                   `json:"client_id"`
	Client     entity_client.EntityClient
	EmployeeID string                           `json:"employee_id" validate:"required"`
	Employee   entity_store.EntityStoreEmployee `gorm:"not null;"`
	CreatedAt  time.Time                        `json:"created_at"`
	UpdatedAt  time.Time                        `json:"updated_at"`
}

func CreateSale(saleParams *EntitySale) error {
	saleParams.ID = utils.GenerateID()
	saleParams.Number = utils.GenerateNumber(time.Now())
	saleParams.CreatedAt = time.Now()
	saleParams.UpdatedAt = time.Now()
	return nil
}

func UpdateSale(saleParams *EntitySale) error {
	saleParams.UpdatedAt = time.Now()
	return nil
}

func (s *EntitySale) Validate() error {
	return validate.Struct(s)
}
