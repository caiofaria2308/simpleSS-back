package entity

import (
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityStoreChain struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null" validate:"required"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateStoreChain(storeChainParams *EntityStoreChain) error {
	storeChainParams.ID = utils.GenerateID()
	storeChainParams.CreatedAt = time.Now()
	storeChainParams.UpdatedAt = time.Now()
	return nil
}

func UpdateStoreChain(storeChainParams *EntityStoreChain) error {
	storeChainParams.UpdatedAt = time.Now()
	return nil
}

func (s *EntityStoreChain) Validate() error {
	return validate.Struct(s)
}
