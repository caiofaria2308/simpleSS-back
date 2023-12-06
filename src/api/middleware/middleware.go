package middleware

import (
	usecase_user "main/usecase/auth/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticatedMiddleware(usercase usecase_user.IUsecaseUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get bearer token from header
		bearerToken := c.Request.Header.Get("Authorization")

		if len(strings.Split(bearerToken, " ")) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}

		token := strings.Split(bearerToken, " ")[1]

		user, err := usercase.GetUserByToken(token)

		// check if token is valid
		if err == nil {

			// set user to context
			c.Set("user", *user)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	}
}
