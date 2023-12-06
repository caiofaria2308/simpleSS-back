package handlers

import (
	"fmt"
	entity_auth "main/entity/auth"
	repo_user "main/infrastructure/database/repository/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func JsonResponse(c *gin.Context, httpStatus int, data any) {
	c.JSON(httpStatus, data)
}

func RoutersHandler(c *gin.Context, r *gin.Engine) {
	type Router struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	var routers []Router = make([]Router, 0)

	for _, route := range r.Routes() {
		routers = append(routers, Router{
			Method: route.Method,
			Path:   route.Path,
		})
	}

	if gin.Mode() == gin.DebugMode {
		c.JSON(200, routers)
	}
}

func VerifyPermission(c *gin.Context, permission string, db *gorm.DB) bool {
	user := c.MustGet("user").(entity_auth.EntityUser)
	repo := repo_user.RepositoryUser{DB: db}
	permissions, err := repo.GetPermissionsByUserID(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(permissions)
	for _, p := range *permissions {
		fmt.Println(p.Slug, permission)
		if p.Slug == permission {
			return true
		}
	}
	return false
}
