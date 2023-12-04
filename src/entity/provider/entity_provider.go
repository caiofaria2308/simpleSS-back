package entity

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

type EntityProvider struct {
	ID                string `json:"id" gorm:"primaryKey"`
	SocialReason      string `json:"social_reason" validate:"required" gorm:"not null"`
	BusinessName      string `json:"business_name" validate:"required" gorm:"not null"`
	AddressState      string `json:"address_state"`
	AddressCity       string `json:"address_city"`
	AddressZipCode    string `json:"address_zip_code"`
	AddressStreet     string `json:"address_street"`
	AddressNumber     string `json:"address_number"`
	AddressComplement string `json:"address_complement"`
}

func CreateProvider(providerParams EntityProvider) (*EntityProvider, error) {
	provider := &EntityProvider{
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
