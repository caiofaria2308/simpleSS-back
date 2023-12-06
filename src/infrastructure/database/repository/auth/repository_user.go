package repository

import (
	entity "main/entity/auth"

	"gorm.io/gorm"
)

type RepositoryUser struct {
	DB *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{DB: db}
}

func (r *RepositoryUser) Create(user entity.EntityUser) (*entity.EntityUser, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryUser) GetAll() (*[]entity.EntityUser, error) {
	var users []entity.EntityUser
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *RepositoryUser) GetByID(id string) (*entity.EntityUser, error) {
	var user entity.EntityUser
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryUser) GetByEmail(email string) (*entity.EntityUser, error) {
	var user entity.EntityUser
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryUser) Update(user entity.EntityUser) (*entity.EntityUser, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryUser) Delete(id string) error {
	var user entity.EntityUser
	err := r.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryUser) GetPermissionsByUserID(id string) (*[]entity.EntityPermission, error) {
	var permissions []entity.EntityPermission
	rgp := *&RepositoryGroup{DB: r.DB}
	rpp := *&RepositoryPermission{DB: r.DB}
	groups, err := rgp.GetUserGroups(id)
	if err != nil {
		return nil, err
	}
	for _, group := range *groups {
		temp_permissions, err := rpp.GetByGroupID(group.ID)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, *temp_permissions...)
	}
	return &permissions, nil
}
