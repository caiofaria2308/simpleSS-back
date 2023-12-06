package entity

import (
	"errors"
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
}

func CreateSale(saleParams EntitySale) (*EntitySale, error) {
	if entity_store.VerifyStoreEmployeeIsActive(&saleParams.Employee) == false {
		return nil, errors.New("Employee is not active")
	}
	s := &EntitySale{
		ID:       utils.GenerateID(),
		Number:   utils.GenerateNumber(saleParams.Date),
		Date:     saleParams.Date,
		Store:    saleParams.Store,
		Client:   saleParams.Client,
		Employee: saleParams.Employee,
	}
	return s, nil
}

func (s *EntitySale) Validate() error {
	return validate.Struct(s)
}
