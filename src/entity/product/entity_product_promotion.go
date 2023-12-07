package entity

import (
	"fmt"
	entity_store "main/entity/store"
	"main/utils"
	"time"
)

type EntityProductPromotion struct {
	ID         string                   `json:"id" gorm:"primaryKey"`
	StoreID    string                   `json:"store_id" gorm:"index:idx_product_promotion_unique" validate:"required"`
	Store      entity_store.EntityStore `gorm:"not null;" validate:"required"`
	Type       string                   `json:"type" gorm:"not null" validate:"required"`
	Percentage float64                  `json:"percentage" gorm:"not null" validate:"required"`
	StartDate  time.Time                `json:"start_date" gorm:"not null" validate:"required"`
	EndDate    time.Time                `json:"end_date" gorm:"not null" validate:"required"`
	Code       string                   `json:"code" gorm:"not null; index:idx_product_promotion_unique"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}

func GenerateCode(EntityProductPromotion *EntityProductPromotion) {
	date_to_string := fmt.Sprintf("%v", EntityProductPromotion.StartDate)
	EntityProductPromotion.Code = date_to_string + EntityProductPromotion.Type
	return
}

func CreateProductPromotion(productPromotionParams *EntityProductPromotion) error {
	productPromotionParams.ID = utils.GenerateID()
	productPromotionParams.CreatedAt = time.Now()
	productPromotionParams.UpdatedAt = time.Now()
	if productPromotionParams.Code == "" {
		GenerateCode(productPromotionParams)
	}
	return nil
}

func UpdateProductPromotion(productPromotionParams *EntityProductPromotion) error {
	productPromotionParams.UpdatedAt = time.Now()
	return nil
}

func (p *EntityProductPromotion) Validate() error {
	return validate.Struct(p)
}
