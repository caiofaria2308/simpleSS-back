package entity

import (
	"main/utils"
)

type EntityPermission struct {
	ID    string `json:"id" gorm:"primary_key"`
	Table string `json:"table" validate:"required" gorm:"not null"`
	Name  string `json:"name" validate:"required" gorm:" not null"`
	Slug  string `json:"slug" validate:"required"`
}

func CreatePermission(permissionParams *EntityPermission) error {
	permissionParams.ID = utils.GenerateID()
	permissionParams.Slug = utils.GenerateSlug(permissionParams.Table + " " + permissionParams.Name)
	return nil
}

func UpdatePermission(permissionParams *EntityPermission) error {
	return nil
}

func (u *EntityPermission) Validate() error {
	return validate.Struct(u)
}
