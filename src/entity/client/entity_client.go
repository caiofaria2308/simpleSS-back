package entity

import (
	entity_store "main/entity/store"
	"main/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityClient struct {
	ID    string                   `json:"id" gorm:"primaryKey"`
	CPF   string                   `json:"cpf" gorm:"unique;not null" validate:"required"`
	Name  string                   `json:"name" validate:"required" gorm:"not null"`
	Store entity_store.EntityStore `json:"store_id" gorm:"foreignKey:StoreID; not null" validate:"required"`
}

func CreateClient(clientParam EntityClient) (*EntityClient, error) {
	u := &EntityClient{
		ID:   utils.GenerateID(),
		CPF:  clientParam.CPF,
		Name: clientParam.Name,
	}
	return u, nil
}

func (u *EntityClient) Validate() error {
	return validate.Struct(u)
}
