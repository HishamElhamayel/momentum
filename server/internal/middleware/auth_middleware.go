package middleware

import (
	"net/http"
	"strings"

	"example.com/momentum/pkg"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims,err := pkg.VerifyToken(tokenString)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
			c.Abort()
			return
		}
		
		c.Set("userID", claims.UserID)
		c.Next()
	}
}