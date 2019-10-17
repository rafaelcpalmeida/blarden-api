package door

import (
	"blarden-api/src/api"
	"blarden-api/src/services"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Open(c *gin.Context) {
	log.Info("Open door requested")

	requestResponse, err := services.RequestOpenDoor()

	if err != nil {
		log.Error(fmt.Sprintf("Open door request errored. Error: %s", err))
		api.ErrorResponse(c, "error", "Door was unable to be opened",
			http.StatusBadRequest)
		return
	}

	log.Info("Open door request served")

	api.ReplyJSON(c, requestResponse, http.StatusOK)
}
