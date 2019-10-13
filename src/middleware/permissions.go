package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func CheckUserPermissions() gin.HandlerFunc {
	allowedPhoneNumber := os.Getenv("ALLOWED_PHONE_NUMBER")

	if allowedPhoneNumber == "" {
		log.Fatal("Please, set ALLOWED_PHONE_NUMBER environment variable")
	}

	return func(c *gin.Context) {
		phoneNumber, exists := c.Request.URL.Query()["phone"]

		if !exists {
			log.Error("No phone number provided")
			c.AbortWithStatus(400)
			return
		}

		if phoneNumber[0] != allowedPhoneNumber {
			log.Error("Invalid phone number")
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
