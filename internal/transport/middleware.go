package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/internal/core"
)

const (
	authHeader     = "Authorization"
	userContextKey = "clientId"
)

func (h *Handler) authenticateClient(c *gin.Context) {
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
	c.Set(userContextKey, id.String())
	return
}

func (h *Handler) getClientId(c *gin.Context) (uuid.UUID, error) {
	idStr := c.GetString(userContextKey)
	return uuid.Parse(idStr)
}
