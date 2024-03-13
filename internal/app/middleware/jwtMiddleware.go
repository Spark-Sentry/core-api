package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRole, roleExists := claims["role"].(string)
			if !roleExists {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Role claim missing in token"})
				return
			}

			userEmail, emailExists := claims["email"].(string)
			if !emailExists {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email claim missing in token"})
				return
			}

			accountId, accountExists := claims["accountId"].(float64)
			if !accountExists {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "accountId claim missing in token"})
				return
			}

			c.Set("userRole", userRole)
			c.Set("userEmail", userEmail)
			c.Set("accountId", accountId)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
