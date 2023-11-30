package entity

import (
	"main/utils"
	"time"
)

type EntityStore struct {
	ID                string           `json:"id" gorm:"primary_key"`
	CNPJ              string           `json:"cnpj" validate:"required"`
	Chain             EntityStoreChain `json:"chain" gorm:"foreignKey:ChainID"`
	SocialReason      string           `json:"social_reason" validate:"required"`
	BusinessName      string           `json:"business_name" validate:"required"`
	AddressState      string           `json:"address_state"`
	AddressCity       string           `json:"address_city"`
	AddressZipCode    string           `json:"address_zip_code"`
	Address           string           `json:"address"`
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
		Address:           storeParams.Address,
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
