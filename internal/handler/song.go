package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

func (h *Handler) getSongById(c *gin.Context) {
	id := h.parseUUIDFromParam(c)
	s, err := h.services.Song.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h *Handler) SearchSongs(c *gin.Context) {

}
