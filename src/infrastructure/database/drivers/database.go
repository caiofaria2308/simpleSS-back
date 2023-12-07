package database

import (
	"fmt"
	entity_auth "main/entity/auth"
	entity_client "main/entity/client"
	entity_product "main/entity/product"
	entity_provider "main/entity/provider"
	entity_purchase "main/entity/purchase"
	entity_sale "main/entity/sale"
	entity_store "main/entity/store"
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
	db.AutoMigrate(&entity_auth.EntityUser{})
	db.AutoMigrate(&entity_auth.EntityGroup{})
	db.AutoMigrate(&entity_auth.EntityPermission{})
	db.AutoMigrate(&entity_store.EntityStoreChain{})
	db.AutoMigrate(&entity_store.EntityStore{})
	db.AutoMigrate(&entity_store.EntityStoreEmployee{})
	db.AutoMigrate(&entity_client.EntityClient{})
	db.AutoMigrate(&entity_product.EntityProduct{})
	db.AutoMigrate(&entity_product.EntityProductGroup{})
	db.AutoMigrate(&entity_product.EntityProductPromotion{})
	db.AutoMigrate(&entity_product.EntityProductPromotionProduct{})
	db.AutoMigrate(&entity_product.EntityProductStock{})
	db.AutoMigrate(&entity_provider.EntityProvider{})
	db.AutoMigrate(&entity_purchase.EntityPurchaseOrder{})
	db.AutoMigrate(&entity_purchase.EntityPurchaseOrderProduct{})
	db.AutoMigrate(&entity_sale.EntitySale{})
	db.AutoMigrate(&entity_sale.EntitySaleProduct{})

	fmt.Println("Migrations completed")
}

func (a *AbstractDatabase) CreatePermissions(db *gorm.DB) {
	fmt.Println("\nCreating permissions...")
	permissions := GeneratePermissions()
	repo := repository.RepositoryPermission{DB: db}
	var existPermission *entity_auth.EntityPermission
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

func GeneratePermissions() (permissions []entity_auth.EntityPermission) {
	permissions = append(permissions, entity_auth.GeneratePermissions()...)
	permissions = append(permissions, entity_client.GeneratePermissions()...)
	permissions = append(permissions, entity_product.GeneratePermissions()...)
	permissions = append(permissions, entity_provider.GeneratePermissions()...)
	permissions = append(permissions, entity_purchase.GeneratePermissions()...)
	permissions = append(permissions, entity_sale.GeneratePermissions()...)
	permissions = append(permissions, entity_store.GeneratePermissions()...)
	return permissions
}

func (a *AbstractDatabase) CreateAdmin(db *gorm.DB) {
	var err error
	var group *entity_auth.EntityGroup
	fmt.Println("\nCreating admin group...")
	repoGroup := repository.RepositoryGroup{DB: db}
	group, err = repoGroup.GetByName("Super Group")
	if err != nil {
		fmt.Println(err)
	}
	if group == nil {
		gp := entity_auth.EntityGroup{
			Name: "Super Group",
		}
		err = entity_auth.CreateGroup(&gp)
		if err != nil {
			fmt.Println(err)
		}
		repoGroup.Create(gp)
	}
	fmt.Println("Admin group created")

	fmt.Println("Creating admin user...")
	repoUser := repository.RepositoryUser{DB: db}
	user, err := repoUser.GetByEmail(os.Getenv("DEFAULT_ADMIN_EMAIL"))
	if err != nil {
		fmt.Println(err)
	}
	if user == nil {
		admin := entity_auth.EntityUser{
			Email:    os.Getenv("DEFAULT_ADMIN_EMAIL"),
			Name:     "Admin",
			Password: os.Getenv("DEFAULT_ADMIN_PASSWORD"),
		}
		err = entity_auth.CreateUser(&admin)
		repoUser.Create(admin)
	}
	fmt.Println("Admin user created")
	fmt.Println("Creating admin permissions...")
	repoPermission := repository.RepositoryPermission{DB: db}
	var permissions *[]entity_auth.EntityPermission
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
