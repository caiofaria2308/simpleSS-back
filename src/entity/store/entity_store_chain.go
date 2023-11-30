package entity

import (
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type EntityStoreChain struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateStoreChain(storeChainParams EntityStoreChain) (*EntityStoreChain, error) {
	s := &EntityStoreChain{
		ID:        utils.GenerateID(),
		Name:      storeChainParams.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s, nil
}

func (s *EntityStoreChain) Validate() error {
	return validate.Struct(s)
}
