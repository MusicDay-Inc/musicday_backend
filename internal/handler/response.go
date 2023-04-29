package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorBody struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf("%d: %s", statusCode, message)
	c.AbortWithStatusJSON(statusCode, ErrorBody{message})
}
