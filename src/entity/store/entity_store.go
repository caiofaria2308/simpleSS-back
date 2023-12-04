package entity

import (
	"main/utils"
	"time"
)

type EntityStore struct {
	ID                string           `json:"id" gorm:"primary_key"`
	CNPJ              string           `json:"cnpj" validate:"required" gorm:"not null; unique"`
	Chain             EntityStoreChain `json:"chain_id" gorm:"foreignKey:ChainID; not null" validate:"required"`
	SocialReason      string           `json:"social_reason" validate:"required" gorm:"not null"`
	BusinessName      string           `json:"business_name" validate:"required" gorm:"not null"`
	AddressState      string           `json:"address_state"`
	AddressCity       string           `json:"address_city"`
	AddressZipCode    string           `json:"address_zip_code"`
	AddressStreet     string           `json:"address_street"`
	AddressNumber     string           `json:"address_number"`
	AddressComplement string           `json:"address_complement"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

func CreateStore(storeParams EntityStore) (*EntityStore, error) {
	s := &EntityStore{
		ID:                utils.GenerateID(),
		CNPJ:              storeParams.CNPJ,
		Chain:             storeParams.Chain,
		SocialReason:      storeParams.SocialReason,
		BusinessName:      storeParams.BusinessName,
		AddressState:      storeParams.AddressState,
		AddressCity:       storeParams.AddressCity,
		AddressZipCode:    storeParams.AddressZipCode,
		AddressStreet:     storeParams.AddressStreet,
		AddressNumber:     storeParams.AddressNumber,
		AddressComplement: storeParams.AddressComplement,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	return s, nil
}

func (s *EntityStore) Validate() error {
	return validate.Struct(s)
}
