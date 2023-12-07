package handlers

import (
	"main/api/handlers"
	handlers_auth "main/api/handlers/auth"
	"main/api/middleware"
	entity "main/entity/store"
	repository "main/infrastructure/database/repository/auth"
	repository_store "main/infrastructure/database/repository/store"
	usecase_user "main/usecase/auth/user"
	usecase_store "main/usecase/store/store"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreHandlers struct {
	UseCaseStore usecase_store.IUseCaseStore
	DB           *gorm.DB
}

func NewStoreHandlers(useCaseStore usecase_store.IUseCaseStore) *StoreHandlers {
	return &StoreHandlers{UseCaseStore: useCaseStore}
}

func (h *StoreHandlers) GetAll(c *gin.Context) {
	chain_id, is_receive := c.GetQuery("chain_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain_id is required"})
		return
	}
	if !handlers.VerifyPermission(c, "entity_store-list", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	stores, err := h.UseCaseStore.GetAll(chain_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}

func (h *StoreHandlers) GetByID(c *gin.Context) {
	var id, chain_id string
	var is_receive bool

	id = c.Param("id")

	chain_id, is_receive = c.GetQuery("chain_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_store-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	store, err := h.UseCaseStore.GetByID(chain_id, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, store)
}

func (h *StoreHandlers) FilterByCNPJ(c *gin.Context) {
	var cnpj, chain_id string
	var is_receive bool

	cnpj = c.Param("cnpj")

	chain_id, is_receive = c.GetQuery("chain_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_store-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	store, err := h.UseCaseStore.FilterByCNPJ(chain_id, cnpj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, store)
}

func (h *StoreHandlers) Search(c *gin.Context) {
	var search, chain_id string
	var is_receive bool

	search = c.Param("search")

	chain_id, is_receive = c.GetQuery("chain_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_store-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	store, err := h.UseCaseStore.Search(chain_id, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, store)
}

func (h *StoreHandlers) Create(c *gin.Context) {
	var store entity.EntityStore
	if !handlers.VerifyPermission(c, "entity_store-create", h.DB) {
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

func (h *StoreHandlers) Update(c *gin.Context) {
	var store entity.EntityStore
	if !handlers.VerifyPermission(c, "entity_store-update", h.DB) {
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

func (h *StoreHandlers) Delete(c *gin.Context) {
	var id, chain_id string
	var is_receive bool

	id = c.Param("id")

	chain_id, is_receive = c.GetQuery("chain_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_store-delete", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	err := h.UseCaseStore.Delete(chain_id, id)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

func MountStoreHandlers(gin *gin.Engine, conn *gorm.DB) {
	handler := NewStoreHandlers(
		usecase_store.NewService(
			repository_store.NewRepositoryStore(conn),
		),
	)
	handler.DB = conn
	userHandlers := handlers_auth.NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)
	group := gin.Group("/api/store")
	group.Use(middleware.AuthenticatedMiddleware(userHandlers.UsecaseUser))
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.GET("/cnpj/:cnpj", handler.FilterByCNPJ)
	group.GET("/search/:search", handler.Search)
	group.POST("/", handler.Create)
	group.PUT("/", handler.Update)
	group.DELETE("/:id", handler.Delete)
	return
}
