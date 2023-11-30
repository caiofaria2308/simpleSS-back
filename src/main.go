package main

// "app/api"
// "app/infrastructure/postgres"
// "app/infrastructure/repository"
// usecase_user "app/usecase/user"
import (
	"main/infrastructure/database/drivers"
)

func main() {
	database := &drivers.SqliteDatabase{}
	db := database.ConnectDB()
	database.RunMigrations(db)
	database.CreatePermissions(db)
	database.CreateAdmin(db)
}
