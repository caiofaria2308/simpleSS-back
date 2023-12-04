package entity

import (
	"fmt"
	entity_store "main/entity/store"
	"main/utils"
	"time"
)

type EntityProductPromotion struct {
	ID         string                   `json:"id" gorm:"primaryKey"`
	Store      entity_store.EntityStore `json:"store_id" gorm:"foreignKey:StoreID; not null" validate:"required"`
	Type       string                   `json:"type" gorm:"not null" validate:"required"`
	Percentage float64                  `json:"percentage" gorm:"not null" validate:"required"`
	StartDate  time.Time                `json:"start_date" gorm:"not null" validate:"required"`
	EndDate    time.Time                `json:"end_date" gorm:"not null" validate:"required"`
	Code       string                   `json:"code" gorm:"not null"`
}

func GenerateCode(EntityProductPromotion *EntityProductPromotion) {
	date_to_string := fmt.Sprintf("%v", EntityProductPromotion.StartDate)
	EntityProductPromotion.Code = date_to_string + EntityProductPromotion.Type
	return
}

func CreateProductPromotion(productPromotionParams EntityProductPromotion) (*EntityProductPromotion, error) {
	p := &EntityProductPromotion{
		ID:         utils.GenerateID(),
		Store:      productPromotionParams.Store,
		Type:       productPromotionParams.Type,
		Percentage: productPromotionParams.Percentage,
		StartDate:  productPromotionParams.StartDate,
		EndDate:    productPromotionParams.EndDate,
		Code:       productPromotionParams.Code,
	}
	if productPromotionParams.Code == "" {
		GenerateCode(p)
	}
	return p, nil
}

func (p *EntityProductPromotion) Validate() error {
	return validate.Struct(p)
}
