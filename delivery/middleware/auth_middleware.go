package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yaqubmw/web-sales-app-golang/utils/security"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		tokenHeader := strings.Split(h.AuthorizationHeader, " ")
		if len(tokenHeader) != 2 || tokenHeader[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		token := tokenHeader[1]
		claims, err := security.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("name", claims["name"])
		c.Next()
	}
}
