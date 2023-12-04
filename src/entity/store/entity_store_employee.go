package entity

import (
	entity "main/entity/auth"
	"main/utils"
)

type EntityStoreEmployee struct {
	ID       string             `json:"id" gorm:"primaryKey"`
	Store    EntityStore        `json:"store_id" gorm:"foreignKey:StoreID; not null" validate:"required"`
	User     entity.EntityUser  `json:"user_id" gorm:"foreignKey:UserID; not null" validate:"required"`
	Group    entity.EntityGroup `json:"group_id" gorm:"foreignKey:GroupID"`
	IsActive bool               `json:"is_active" gorm:"default:true"`
}

func CreateStoreEmployee(storeEmployeeParams EntityStoreEmployee) (*EntityStoreEmployee, error) {
	s := &EntityStoreEmployee{
		ID:       utils.GenerateID(),
		Store:    storeEmployeeParams.Store,
		User:     storeEmployeeParams.User,
		Group:    storeEmployeeParams.Group,
		IsActive: storeEmployeeParams.IsActive,
	}
	return s, nil
}

func (s *EntityStoreEmployee) Validate() error {
	return validate.Struct(s)
}

func VerifyStoreEmployeeIsActive(storeEmployee *EntityStoreEmployee) bool {
	return storeEmployee.IsActive
}
