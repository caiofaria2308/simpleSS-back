package usecase_user

import entity "main/entity/auth"

type IRepositoryUser interface {
	GetUserByToken(token string) (*entity.EntityUser, error)
	GetByEmail(email string) (*entity.EntityUser, error)
	CreateUser(user entity.EntityUser) (*entity.EntityUser, error)
	UpdateUser(user entity.EntityUser) (*entity.EntityUser, error)
	DeleteUser(id string) error
}

type IUsecaseUser interface {
	LoginUser(email string, password string) (*entity.EntityUser, error)
	Create(user entity.EntityUser) (*entity.EntityUser, error)
	Update(user entity.EntityUser) (*entity.EntityUser, error)
	Delete(id string) error
	GetByEmail(email string) (*entity.EntityUser, error)
	GetUserByToken(token string) (*entity.EntityUser, error)
}
