package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

const (
	authHeader     = "Authorization"
	userContextKey = "clientId"
)

func (h *Handler) authenticateUser(c *gin.Context) {
	header := c.GetHeader(authHeader)
	id, registered, err := h.services.ParseToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, core.CodeTokenInvalid, err.Error())
		return
	}
	if !registered {
		newErrorResponse(c, http.StatusForbidden, core.CodeAccessDenied, core.ErrAccessDenied.Error())
		return
	}
	c.Set(userContextKey, id)
	return
}
