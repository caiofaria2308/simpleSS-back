package handlers

import (
	"main/api/handlers"
	handlers_auth "main/api/handlers/auth"
	"main/api/middleware"
	entity_client "main/entity/client"
	repository "main/infrastructure/database/repository/auth"
	repository_client "main/infrastructure/database/repository/client"
	usecase_user "main/usecase/auth/user"
	usecase_client "main/usecase/client/client"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientHandlers struct {
	UseCaseClient usecase_client.IUseCaseClient
	DB            *gorm.DB
}

func NewClientHandlers(useCaseClient usecase_client.IUseCaseClient) *ClientHandlers {
	return &ClientHandlers{UseCaseClient: useCaseClient}
}

func (h *ClientHandlers) GetAll(c *gin.Context) {
	store_id, is_receive := c.GetQuery("store_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}
	if !handlers.VerifyPermission(c, "entity_client-list", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	clients, err := h.UseCaseClient.GetAll(store_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandlers) GetByID(c *gin.Context) {
	var id, store_id string
	var is_receive bool

	id = c.Param("id")

	store_id, is_receive = c.GetQuery("store_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_client-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	client, err := h.UseCaseClient.GetByID(store_id, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandlers) GetByCPF(c *gin.Context) {
	var cpf, store_id string
	var is_receive bool
	cpf = c.Param("cpf")

	store_id, is_receive = c.GetQuery("store_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_client-read", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	client, err := h.UseCaseClient.GetByCPF(store_id, cpf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandlers) Search(c *gin.Context) {
	var search, store_id string
	var is_receive bool

	search = c.Param("search")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search is required"})
		return
	}

	store_id, is_receive = c.GetQuery("store_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	if !handlers.VerifyPermission(c, "entity_client-list", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	clients, err := h.UseCaseClient.Search(store_id, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandlers) Create(c *gin.Context) {
	if !handlers.VerifyPermission(c, "entity_client-create", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	var entityClient entity_client.EntityClient
	if err := c.ShouldBindJSON(&entityClient); err != nil {
		handlers.HandleError(c, err)
		return
	}
	client, err := h.UseCaseClient.Create(entityClient)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, client)

}

func (h *ClientHandlers) Update(c *gin.Context) {
	if !handlers.VerifyPermission(c, "entity_client-update", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	var entityClient entity_client.EntityClient
	if err := c.ShouldBindJSON(&entityClient); err != nil {
		handlers.HandleError(c, err)
		return
	}
	client, err := h.UseCaseClient.Update(entityClient)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandlers) Delete(c *gin.Context) {
	if !handlers.VerifyPermission(c, "entity_client-delete", h.DB) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission"})
		return
	}
	client_id := c.Param("id")
	store, is_receive := c.GetQuery("store_id")
	if !is_receive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	err := h.UseCaseClient.Delete(store, client_id)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

func MountClientHandlers(gin *gin.Engine, conn *gorm.DB) {
	handler := NewClientHandlers(
		usecase_client.NewService(
			repository_client.NewRepositoryClient(conn),
		),
	)
	handler.DB = conn
	userHandlers := handlers_auth.NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)
	group := gin.Group("/api/clients")
	group.Use(middleware.AuthenticatedMiddleware(userHandlers.UsecaseUser))
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.GET("/cpf/:cpf", handler.GetByCPF)
	group.GET("/search/:search", handler.Search)
	group.POST("/", handler.Create)
	group.PUT("/", handler.Update)
	group.DELETE("/:id", handler.Delete)
}
