package api

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse Json response error structure
type errorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// DefaultHandler - Default handler
func DefaultHandler(ginContext *gin.Context) {
	ginContext.String(200, "ALIVE")
}

func ReplyJSON(ginContext *gin.Context, payload interface{}, statusCode int) {
	ginContext.JSON(statusCode, payload)
}

func ErrorResponse(ginContext *gin.Context, errorCode string, message string, statusCode int) {
	apiResponse := errorResponse{ErrorCode: errorCode, ErrorMessage: message}
	ginContext.AbortWithStatusJSON(statusCode, apiResponse)
}

