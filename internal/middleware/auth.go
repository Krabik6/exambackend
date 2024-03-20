package middleware

import (
	"exambackend/pkg/jwtn"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			return
		}

		// Извлечение токена, удаляя префикс "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID, role, err := jwtn.VerifyToken(token) // Передаем в функцию уже без "Bearer "
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			return
		}

		// Добавляем userID и role в контекст запроса
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

// AdminRoleMiddleware проверяет, что пользователь имеет роль администратора
func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлечение роли пользователя из контекста (предполагается, что AuthMiddleware уже добавила эту информацию)
		userRole, exists := c.Get("role")
		if !exists || userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
			c.Abort()
			return
		}
		c.Next()
	}
}
