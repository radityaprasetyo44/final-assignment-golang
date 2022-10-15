package middleware

import (
	"final-assignment/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyBearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "invalid token",
			})
			return
		}

		c.Set("user_data", verifyToken)
		c.Next()
	}
}
