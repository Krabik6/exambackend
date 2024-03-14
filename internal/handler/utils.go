package handler

import "github.com/gin-gonic/gin"

func getUserIDFromContext(c *gin.Context) (int64, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	// Убедитесь, что userID действительно имеет тип int64
	id, ok := userID.(int64)
	return id, ok
}
