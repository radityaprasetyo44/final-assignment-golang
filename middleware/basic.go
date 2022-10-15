package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	isValid := hasAuth && user == "user" && password == "user"
	if !isValid {
		c.Abort()
		result := gin.H{
			"result": "unauthorized access",
		}
		c.JSON(http.StatusUnauthorized, result)
	}
}
