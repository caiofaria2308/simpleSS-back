package entity

import (
	"main/utils"
	"time"
)

type EntityStore struct {
	ID                string           `json:"id" gorm:"primary_key"`
	CNPJ              string           `json:"cnpj" validate:"required" gorm:"not null; unique; index:idx_store_unique"`
	ChainID           string           `json:"chain_id" validate:"required" gorm:" index:idx_store_unique"`
	Chain             EntityStoreChain ` gorm:"not null;"`
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

func CreateStore(storeParams *EntityStore) error {
	storeParams.ID = utils.GenerateID()
	storeParams.CreatedAt = time.Now()
	storeParams.UpdatedAt = time.Now()
	return nil
}

func UpdateStore(storeParams *EntityStore) error {
	storeParams.UpdatedAt = time.Now()
	return nil
}

func (s *EntityStore) Validate() error {
	return validate.Struct(s)
}
