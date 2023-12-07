package api

import (
	"log"

	"main/api/handlers"
	handlers_auth "main/api/handlers/auth"
	handlers_client "main/api/handlers/client"
	handlers_store "main/api/handlers/store"
	database "main/infrastructure/database/drivers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupDatabase() *gorm.DB {
	database := &database.SqliteDatabase{}
	conn := database.ConnectDB()
	return conn
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handlers.MountSamplesHandlers(r)
	handlers_auth.MountHandlers(r, db)
	handlers_store.MountStoreChainHandlers(r, db)
	handlers_store.MountStoreHandlers(r, db)
	handlers_client.MountClientHandlers(r, db)

	return r
}

func SetupRouters(db *gorm.DB) *gin.Engine {
	return setupRouter(db)
}

func StartWebServer(db *gorm.DB) {

	r := SetupRouters(db)

	// Bind to a port and pass our router in
	log.Fatal(r.Run())
}
