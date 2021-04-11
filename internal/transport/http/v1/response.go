package v1

import (
	"github.com/IDarar/hub/pkg/logger"
	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

type IdResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
