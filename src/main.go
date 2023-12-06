package main

// "app/api"
// "app/infrastructure/postgres"
// "app/infrastructure/repository"
// usecase_user "app/usecase/user"
import (
	"main/api"
	database "main/infrastructure/database/drivers"
)

func main() {
	database := &database.SqliteDatabase{}
	db := database.ConnectDB()
	database.RunMigrations(db)
	database.CreatePermissions(db)
	database.CreateAdmin(db)
	api.StartWebServer(db)
}
