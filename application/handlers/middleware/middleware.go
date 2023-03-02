package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func IdentifierMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")

		if paramId != "" {
			_, err := strconv.ParseUint(c.Param("id"), 10, 64)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "only numbers are accepted as an identifier parameter"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
