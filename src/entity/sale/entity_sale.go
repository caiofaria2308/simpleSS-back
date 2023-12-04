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
	ID     string                     `json:"id" gorm:"primaryKey"`
	Date   time.Time                  `json:"date" gorm:"not null"`
	Number string                     `json:"number" gorm:"unique; not null"`
	Store  entity_store.EntityStore   `json:"store_id" gorm:"foreignKey:StoreID; not null" validate:"required"`
	Client entity_client.EntityClient `json:"client_id" gorm:"foreignKey:ClientID"`
}

func CreateSale(saleParams EntitySale) (*EntitySale, error) {
	s := &EntitySale{
		ID:     utils.GenerateID(),
		Number: utils.GenerateNumber(saleParams.Date),
		Date:   saleParams.Date,
		Store:  saleParams.Store,
		Client: saleParams.Client,
	}
	return s, nil
}

func (s *EntitySale) Validate() error {
	return validate.Struct(s)
}
