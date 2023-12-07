package entity

import (
	entity_store "main/entity/store"
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityProvider struct {
	ID                string                   `json:"id" gorm:"primaryKey"`
	StoreID           string                   `json:"store_id" gorm:"index:idx_provider_unique" validate:"required"`
	Store             entity_store.EntityStore `gorm:"not null;"`
	SocialReason      string                   `json:"social_reason" validate:"required" gorm:"not null"`
	CNPJ              string                   `json:"cnpj" validate:"required" gorm:"not null; index:idx_provider_unique"`
	BusinessName      string                   `json:"business_name" validate:"required" gorm:"not null"`
	AddressState      string                   `json:"address_state"`
	AddressCity       string                   `json:"address_city"`
	AddressZipCode    string                   `json:"address_zip_code"`
	AddressStreet     string                   `json:"address_street"`
	AddressNumber     string                   `json:"address_number"`
	AddressComplement string                   `json:"address_complement"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

func CreateProvider(providerParams *EntityProvider) error {
	providerParams.ID = utils.GenerateID()
	providerParams.CreatedAt = time.Now()
	providerParams.UpdatedAt = time.Now()
	return nil
}

func UpdateProvider(providerParams *EntityProvider) error {
	providerParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityProvider) Validate() error {
	return validate.Struct(p)
}
