package entity

import (
	"main/utils"
	"time"
)

type EntityGroup struct {
	ID          string             `gorm:"primaryKey" json:"id"`
	Name        string             `gorm:"unique; not null" json:"name"`
	Permissions []EntityPermission `gorm:"many2many:group_permissions;" json:"permissions"`
	Users       []EntityUser       `gorm:"many2many:user_groups;" json:"users"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func CreateGroup(groupParams *EntityGroup) error {
	groupParams.ID = utils.GenerateID()
	groupParams.CreatedAt = time.Now()
	groupParams.UpdatedAt = time.Now()
	return nil
}

func UpdateGroup(groupParams *EntityGroup) error {
	groupParams.UpdatedAt = time.Now()
	return nil
}

func (g *EntityGroup) Validate() error {
	return validate.Struct(g)
}
