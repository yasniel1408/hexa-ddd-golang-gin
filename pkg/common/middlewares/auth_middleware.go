package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Authorization header missing"})
			c.Abort()
			return
		}

		// El token debe tener el formato "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parsear y validar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Guardar la información del usuario en el contexto si es necesario
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		}

		c.Next()
	}
}
