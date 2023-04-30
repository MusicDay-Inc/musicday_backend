package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/core"
)

func bindRequestBody(c *gin.Context, obj any) {
	if err := c.BindJSON(&obj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrIncorrectBody.Error())
		return
	}
}

func newErrorResponse(c *gin.Context, statusCode int, code int, message string) {
	logrus.Infof("%d: %s", statusCode, message)
	c.AbortWithStatusJSON(statusCode, core.ErrorBody{Message: message, Code: code})
}
