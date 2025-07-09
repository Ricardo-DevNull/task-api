package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("JWT_SECRET_KEY")

		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}

		tokenHash := strings.TrimPrefix(authToken, "Bearer ")
		if tokenHash == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}

		claims := jwt.MapClaims{}
		decodeToken, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretKey), nil
		})

		if err != nil || !decodeToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := decodeToken.Claims.(jwt.MapClaims)
		if !ok || !decodeToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userID, userIDContains := claims["userID"].(float64)
		role, roleContains := claims["role"].(string)
		if !userIDContains || !roleContains {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token data"})
			return
		}

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}
