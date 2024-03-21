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
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-ijt")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
