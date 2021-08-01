package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MessageErr struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewMessageErrorResponse(c *gin.Context, statuscode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statuscode, MessageErr{message})
}
