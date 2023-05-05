package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/core"
)

func bindRequestBody(c *gin.Context, obj any) bool {
	if err := c.BindJSON(&obj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrIncorrectBody.Error())
		return false
	}
	return true
}

func newErrorResponse(c *gin.Context, statusCode int, code int, message string) {
	logrus.Infof("%d: %s", statusCode, message)
	c.AbortWithStatusJSON(statusCode, core.ErrorBody{Message: message, Code: code})
}

// parseUUIDFromParam returns Error response if it couldn't parse token
func (h *Handler) parseUUIDFromParam(c *gin.Context) uuid.UUID {
	id := c.Param("id")
	itemUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, core.CodeIncorrectBody, "could not parse uuid from id parameter")
		return uuid.Nil
	}
	return itemUUID
}
