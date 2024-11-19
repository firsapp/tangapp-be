package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"tangapp-be/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Validate if "Bearer " within "authHeader" exists
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix("authHeader", "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ") //	Removes "Bearer " prefix from jwt token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Validate if JWT signing method is 256, 384, or 512
				return nil, fmt.Errorf("wrong JWT signing method")
			}
			return []byte(config.JWTSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["sub"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}
