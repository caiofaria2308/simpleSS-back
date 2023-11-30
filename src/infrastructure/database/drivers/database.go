package drivers

import (
	"fmt"
	entity "main/entity/auth"
	repository "main/infrastructure/database/repository/auth"
	"os"

	"gorm.io/gorm"
)

type DatabaseInterface interface {
	ConnectDB() (db *gorm.DB)
	RunMigrations(db *gorm.DB)
}

type AbstractDatabase struct {
	DatabaseInterface
}

func (a *AbstractDatabase) ConnectDB() gorm.DB {
	panic("ConnectDB not implemented")
}

func (a *AbstractDatabase) RunMigrations(db *gorm.DB) {
	fmt.Println("\nRunning migrations...")
	db.AutoMigrate(&entity.EntityUser{})
	db.AutoMigrate(&entity.EntityGroup{})
	db.AutoMigrate(&entity.EntityPermission{})
	fmt.Println("Migrations completed")
}

func (a *AbstractDatabase) CreatePermissions(db *gorm.DB) {
	fmt.Println("\nCreating permissions...")
	permissions := entity.GeneratePermissions()
	repo := repository.RepositoryPermission{DB: db}
	var existPermission *entity.EntityPermission
	var err error
	for _, permission := range permissions {
		existPermission, err = repo.GetBySlug(permission.Slug)
		if err != nil {
			fmt.Println(err)
		}
		if existPermission == nil {
			repo.Create(permission)
		}
	}
	fmt.Println("Permissions created")
}

func (a *AbstractDatabase) CreateAdmin(db *gorm.DB) {
	var err error
	var group *entity.EntityGroup
	fmt.Println("\nCreating admin group...")
	repoGroup := repository.RepositoryGroup{DB: db}
	group, err = repoGroup.GetByName("Super Group")
	if err != nil {
		fmt.Println(err)
	}
	if group == nil {
		group, err = entity.CreateGroup(entity.EntityGroup{
			Name: "Super Group",
		})
		if err != nil {
			fmt.Println(err)
		}
		repoGroup.Create(*group)
	}
	fmt.Println("Admin group created")

	fmt.Println("Creating admin user...")
	repoUser := repository.RepositoryUser{DB: db}
	user, err := repoUser.GetByEmail(os.Getenv("DEFAULT_ADMIN_EMAIL"))
	if err != nil {
		fmt.Println(err)
	}
	if user == nil {
		user, err = entity.CreateUser(entity.EntityUser{
			Email:    os.Getenv("DEFAULT_ADMIN_EMAIL"),
			Name:     "Admin",
			Password: os.Getenv("DEFAULT_ADMIN_PASSWORD"),
		})
		repoUser.Create(*user)
	}
	fmt.Println("Admin user created")
	fmt.Println("Creating admin permissions...")
	repoPermission := repository.RepositoryPermission{DB: db}
	var permissions *[]entity.EntityPermission
	permissions, err = repoPermission.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, permission := range *permissions {
		group.Permissions = append(group.Permissions, permission)
	}
	group.Users = append(group.Users, *user)
	repoGroup.Update(*group)
	fmt.Println("Admin permissions created")

}
