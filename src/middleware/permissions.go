package middleware

import (
	"blarden-api/src/db/models"
	"blarden-api/src/services"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func AdminPermissions() gin.HandlerFunc {
	queryParameters := make(map[string]interface{})
	aesToken, err := services.GetAESToken()

	return func(c *gin.Context) {
		encryptedAuthorizationToken := c.Request.Header.Get("Authorization")

		if err != nil {
			log.Error("Can't get encryption key.")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if encryptedAuthorizationToken == "" {
			log.Error("No user token provided")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		decryptedAuthorizationToken, _ := hex.DecodeString(encryptedAuthorizationToken)
		authorizationToken, err := services.Decrypt(decryptedAuthorizationToken, &aesToken)

		if err != nil {
			log.Error("Error decrypting token")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		queryParameters["user_unique_key"] = authorizationToken

		users, err := models.AllUsers(queryParameters)

		if err != nil {
			log.Error("Error querying users")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if len(users) != 1 {
			log.Error("No user was found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user := users[0]; user.UserType != 0 {
			log.Error("User is not administrator")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func CheckUserPermissions() gin.HandlerFunc {
	allowedPhoneNumber := os.Getenv("ALLOWED_PHONE_NUMBER")

	if allowedPhoneNumber == "" {
		log.Fatal("Please, set ALLOWED_PHONE_NUMBER environment variable")
	}

	return func(c *gin.Context) {
		phoneNumber, exists := c.Request.URL.Query()["phone"]

		if !exists {
			log.Error("No phone number provided")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if phoneNumber[0] != allowedPhoneNumber {
			log.Error("Invalid phone number")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
