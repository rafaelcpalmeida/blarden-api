package users

import (
	"blarden-api/src/api"
	"blarden-api/src/db/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
)

func All(c *gin.Context) {
	users, err := models.AllUsers()

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
	user.ID, _ = uuid.NewV4()
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

	api.ReplyJSON(c, createdUser, http.StatusCreated)
}

func Update(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		api.ErrorResponse(c, "invalid_id", "The user id is invalid", http.StatusBadRequest)
	}

	user := models.User{}
	err = c.ShouldBindJSON(&user)

	user.ID = id

	if err != nil {
		api.ErrorResponse(c, "invalid_payload", "The user payload given is invalid",
			http.StatusBadRequest)
		return
	}

	updatedUser, err := models.UpdateUser(user)
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