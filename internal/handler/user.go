package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

// GET

func (h *Handler) SearchUsers(c *gin.Context) {

}

// POST
func (h *Handler) subscribe(c *gin.Context) {
	userId := h.parseUUIDFromParam(c)
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
		return
	}
	updatedUserSubscription, err := h.services.User.Subscribe(clientId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrIncorrectBody.Error())
		return
	}
	c.JSON(http.StatusOK, updatedUserSubscription)
}

func (h *Handler) postStory(c *gin.Context) {

}
