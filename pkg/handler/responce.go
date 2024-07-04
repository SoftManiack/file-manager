package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type HttpResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)

	c.AbortWithStatusJSON(statusCode, errorResponse{message, "error", statusCode})
}
