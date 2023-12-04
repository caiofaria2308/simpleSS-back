package entity

import (
	entity_store "main/entity/store"
	"main/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityProvider struct {
	ID                string                   `json:"id" gorm:"primaryKey"`
	Store             entity_store.EntityStore `json:"store_id" gorm:"foreignKey:StoreID; not null; index:idx_unique" validate:"required"`
	SocialReason      string                   `json:"social_reason" validate:"required" gorm:"not null"`
	CNPJ              string                   `json:"cnpj" validate:"required" gorm:"not null; index:idx_unique"`
	BusinessName      string                   `json:"business_name" validate:"required" gorm:"not null"`
	AddressState      string                   `json:"address_state"`
	AddressCity       string                   `json:"address_city"`
	AddressZipCode    string                   `json:"address_zip_code"`
	AddressStreet     string                   `json:"address_street"`
	AddressNumber     string                   `json:"address_number"`
	AddressComplement string                   `json:"address_complement"`
}

func CreateProvider(providerParams EntityProvider) (*EntityProvider, error) {
	provider := &EntityProvider{
		ID:                utils.GenerateID(),
		CNPJ:              providerParams.CNPJ,
		SocialReason:      providerParams.SocialReason,
		BusinessName:      providerParams.BusinessName,
		AddressState:      providerParams.AddressState,
		AddressCity:       providerParams.AddressCity,
		AddressZipCode:    providerParams.AddressZipCode,
		AddressStreet:     providerParams.AddressStreet,
		AddressNumber:     providerParams.AddressNumber,
		AddressComplement: providerParams.AddressComplement,
	}

	return provider, nil
}

func (p *EntityProvider) Validate() error {
	return validate.Struct(p)
}
