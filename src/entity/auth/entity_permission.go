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

func CreatePermission(permissionParam EntityPermission) (*EntityPermission, error) {
	slug := utils.GenerateSlug(permissionParam.Table + " " + permissionParam.Name)
	u := &EntityPermission{
		ID:    utils.GenerateID(),
		Table: permissionParam.Table,
		Name:  permissionParam.Name,
		Slug:  slug,
	}
	return u, nil
}

func (u *EntityPermission) Validate() error {
	return validate.Struct(u)
}
