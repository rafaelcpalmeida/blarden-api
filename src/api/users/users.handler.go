package users

import (
	"blarden-api/src/api"
	"blarden-api/src/db/models"
	"blarden-api/src/services"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/sha3"
	"net/http"
	"time"
)

func All(c *gin.Context) {
	queryParameters := make(map[string]interface{})

	if token := c.Query("token"); token != "" {
		queryParameters["user_unique_key"] = c.Query("token")
	}

	users, err := models.AllUsers(queryParameters)

	if err != nil {
		api.ErrorResponse(c, "list_error", fmt.Sprintf("Unable to show users. Error: %s", err),
			http.StatusBadRequest)
	}

	api.ReplyJSON(c, users, http.StatusOK)
}

func WithId(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		api.ErrorResponse(c, "invalid_id", "The user id is invalid", http.StatusBadRequest)
	}

	user, err := models.SpecificUser(id)
	if err != nil {
		api.ErrorResponse(c, "list_error", fmt.Sprintf("Unable to show user. Error: %s", err),
			http.StatusBadRequest)
		return
	}

	api.ReplyJSON(c, user, http.StatusOK)
}

func Create(c *gin.Context) {
	user := models.User{}
	user.Id, _ = uuid.NewV4()
	userUniqueSecretToken := sha3.Sum512([]byte(fmt.Sprintf("%d", time.Now().Unix())))
	user.UserUniqueKey = hex.EncodeToString(userUniqueSecretToken[:])
	err := c.ShouldBindJSON(&user)
	if err != nil {
		api.ErrorResponse(c, "invalid_payload", "The user payload given is invalid",
			http.StatusBadRequest)
		return
	}

	createdUser, err := models.NewUser(user)

	if err != nil {
		api.ErrorResponse(c, "create_error", fmt.Sprintf("Unable to create the new user. Error: %s", err),
			http.StatusBadRequest)
		return
	}

	aesKey, err := services.GetAESToken()

	if err != nil {
		api.ErrorResponse(c, "Error", "Can't get encryption key. Check logs", http.StatusInternalServerError)
	}

	encryptedUserKey, err := services.Encrypt([]byte(createdUser.UserUniqueKey), &aesKey)


	if err != nil {
		api.ErrorResponse(c, "Error", "Can't encrypt user payload. Check logs", http.StatusInternalServerError)
	}

	createdUser.UserUniqueKey = hex.EncodeToString(encryptedUserKey)

	api.ReplyJSON(c, createdUser, http.StatusCreated)
}

func Update(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		api.ErrorResponse(c, "invalid_id", "The user id is invalid", http.StatusBadRequest)
	}

	user := models.User{}
	err = c.ShouldBindJSON(&user)

	user.Id = id

	if err != nil {
		api.ErrorResponse(c, "invalid_payload", "The user payload given is invalid",
			http.StatusBadRequest)
		return
	}

	updatedUser, err := models.UpdateUser(id, user)
	if err != nil {
		api.ErrorResponse(c, "create_error", fmt.Sprintf("Unable to update user. Error: %s", err),
			http.StatusBadRequest)
		return
	}

	api.ReplyJSON(c, updatedUser, http.StatusCreated)
}

func Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		api.ErrorResponse(c, "invalid_id", "The user id is invalid", http.StatusBadRequest)
	}

	_, err = models.DeleteUser(id)
	if err != nil {
		api.ErrorResponse(c, "delete_error", fmt.Sprintf("Unable to delete user. Error: %s", err),
			http.StatusBadRequest)
		return
	}

	api.ReplyJSON(c, nil, http.StatusNoContent)
}