package entity

import (
	entity "main/entity/auth"
	"main/utils"
	"time"
)

type EntityStoreEmployee struct {
	ID        string             `json:"id" gorm:"primaryKey"`
	StoreID   string             `json:"store_id" gorm:"index:idx_store_employee_unique" validate:"required"`
	Store     EntityStore        `gorm:"not null;" validate:"required"`
	UserID    string             `json:"user_id" gorm:"index:idx_store_employee_unique" validate:"required"`
	User      entity.EntityUser  `gorm:"not null;" validate:"required"`
	GroupID   string             `json:"group_id" validate:"required"`
	Group     entity.EntityGroup `gorm:"not null;" validate:"required"`
	IsActive  bool               `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func CreateStoreEmployee(storeEmployeeParams *EntityStoreEmployee) error {
	storeEmployeeParams.ID = utils.GenerateID()
	storeEmployeeParams.CreatedAt = time.Now()
	storeEmployeeParams.UpdatedAt = time.Now()
	return nil
}

func UpdateStoreEmployee(storeEmployeeParams *EntityStoreEmployee) error {
	storeEmployeeParams.UpdatedAt = time.Now()
	return nil
}

func (s *EntityStoreEmployee) Validate() error {
	return validate.Struct(s)
}

func VerifyStoreEmployeeIsActive(storeEmployee *EntityStoreEmployee) bool {
	return storeEmployee.IsActive
}
