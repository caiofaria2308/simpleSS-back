package handlers

import (
	entity "main/entity/auth"
	repository "main/infrastructure/database/repository/auth"
	usecase_user "main/usecase/auth/user"
	"net/http"

	"main/api/handlers"
	middleware "main/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordData struct {
	Email           string `json:"email"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserHandlers struct {
	UsecaseUser usecase_user.IUsecaseUser
}

func NewUserHandler(usecaseUser usecase_user.IUsecaseUser) *UserHandlers {
	return &UserHandlers{UsecaseUser: usecaseUser}
}

func (h UserHandlers) LoginHandler(c *gin.Context) {

	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		handlers.HandleError(c, err)
		return
	}

	user, err := h.UsecaseUser.LoginUser(loginData.Email, loginData.Password)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	token, refreshToken, err := usecase_user.JWTTokenGenerator(*user)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}

func (h UserHandlers) GetMeHandler(c *gin.Context) {
	user, err := h.UsecaseUser.GetUserByToken(c.GetHeader("Authorization"))

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, user)
}

func (h UserHandlers) CreateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handlers.HandleError(c, err)
		return
	}

	user, err := h.UsecaseUser.Create(entityUser)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, gin.H{"message": "User created successfully", "user": user})

}

func (h UserHandlers) UpdateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handlers.HandleError(c, err)
		return
	}

	user, err := h.UsecaseUser.Update(entityUser)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func (h UserHandlers) DeleteUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handlers.HandleError(c, err)
		return
	}

	err := h.UsecaseUser.Delete(entityUser.ID)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h UserHandlers) GetUserHandler(c *gin.Context) {

	email := c.Param("email")
	user, err := h.UsecaseUser.GetByEmail(email)

	if exception := handlers.HandleError(c, err); exception {
		return
	}

	handlers.JsonResponse(c, http.StatusOK, user)
}

func MountHandlers(gin *gin.Engine, conn *gorm.DB) {

	userHandlers := NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)

	gin.POST("/api/login", userHandlers.LoginHandler)

	// user
	group := gin.Group("/api/user")
	group.POST("/create", userHandlers.CreateUserHandler)
	group.Use(middleware.AuthenticatedMiddleware(userHandlers.UsecaseUser))
	group.GET("/me", userHandlers.GetMeHandler)
	group.PUT("/", userHandlers.UpdateUserHandler)
	group.DELETE("/:id", userHandlers.DeleteUserHandler)
	group.GET("/:email", userHandlers.GetUserHandler)
}
