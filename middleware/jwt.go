package middleware

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AuthenticateJWT(role string, config configuration.Config) gin.HandlerFunc {
	//jwtSecret := config.Get("JWT_SECRET_KEY")
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
