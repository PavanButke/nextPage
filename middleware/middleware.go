package middleware

import (
	token "code/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")

		if ClientToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}

}
