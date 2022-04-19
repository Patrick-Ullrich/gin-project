package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// access the status we are sending
		status := c.Writer.Status()
		latency := time.Since(t)
		log.Printf("[%s] %s => %d - %s", c.Request.Method, c.Request.URL.Path, status, latency)
	}
}
