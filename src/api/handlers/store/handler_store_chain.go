package handlers

import (
	"main/api/handlers"
	handlers_auth "main/api/handlers/auth"
	"main/api/middleware"
	entity_store "main/entity/store"
	repository "main/infrastructure/database/repository/auth"
	repository_store_chain "main/infrastructure/database/repository/store"
	usecase_user "main/usecase/auth/user"
	usecase_store_chain "main/usecase/store/store_chain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreChainHandlers struct {
	UseCaseStore usecase_store_chain.IUseCaseStoreChain
	DB           *gorm.DB
}

func NewStoreChainHandlers(useCaseStore usecase_store_chain.IUseCaseStoreChain) *StoreChainHandlers {
	return &StoreChainHandlers{UseCaseStore: useCaseStore}
}

func (h *StoreChainHandlers) GetAll(c *gin.Context) {
	if !handlers.VerifyPermission(c, "entity_store_chain-list", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	stores, err := h.UseCaseStore.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}

func (h *StoreChainHandlers) GetByID(c *gin.Context) {
	id := c.Param("id")
	if !handlers.VerifyPermission(c, "entity_store_chain-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	store, err := h.UseCaseStore.GetByID(id)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, store)
}

func (h *StoreChainHandlers) SearchByName(c *gin.Context) {
	name := c.Param("name")
	if !handlers.VerifyPermission(c, "entity_store_chain-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	stores, err := h.UseCaseStore.SearchByName(name)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, stores)
}

func (h *StoreChainHandlers) Create(c *gin.Context) {
	var store entity_store.EntityStoreChain
	if !handlers.VerifyPermission(c, "entity_store_chain-create", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store_chain, err := h.UseCaseStore.Create(store)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, store_chain)
}

func (h *StoreChainHandlers) Update(c *gin.Context) {
	var store entity_store.EntityStoreChain
	if !handlers.VerifyPermission(c, "entity_store_chain-update", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store_chain, err := h.UseCaseStore.Update(store)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, store_chain)
}

func (h *StoreChainHandlers) Delete(c *gin.Context) {
	if !handlers.VerifyPermission(c, "entity_store_chain-delete", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	store_id := c.Param("id")
	err := h.UseCaseStore.Delete(store_id)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

func MountStoreChainHandlers(gin *gin.Engine, conn *gorm.DB) {
	handler := NewStoreChainHandlers(
		usecase_store_chain.NewService(
			repository_store_chain.NewRepositoryStoreChain(conn),
		),
	)
	handler.DB = conn
	userHandlers := handlers_auth.NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)
	group := gin.Group("/api/store-chain/")
	group.Use(middleware.AuthenticatedMiddleware(userHandlers.UsecaseUser))
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.GET("/search/:name", handler.SearchByName)
	group.POST("/", handler.Create)
	group.PUT("/", handler.Update)
	group.DELETE("/:id", handler.Delete)

}
