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

func CreateGroup(groupParams EntityGroup) (*EntityGroup, error) {
	g := &EntityGroup{
		ID:          utils.GenerateID(),
		Name:        groupParams.Name,
		Permissions: groupParams.Permissions,
		Users:       groupParams.Users,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return g, nil
}

func (g *EntityGroup) Validate() error {
	return validate.Struct(g)
}
