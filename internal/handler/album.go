package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

func (h *Handler) getAlbumById(c *gin.Context) {
	id := h.parseUUIDFromParam(c)
	a, err := h.services.Album.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	c.JSON(http.StatusOK, a.ToDTO())
}
func (h *Handler) getAlbumWitSongsById(c *gin.Context) {
	id := h.parseUUIDFromParam(c)
	a, err := h.services.Album.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	songs, err := h.services.Album.GetSongsFromAlbum(id)
	c.JSON(http.StatusOK, a.ToFullDTO(songs))
}

func (h *Handler) SearchAlbums(context *gin.Context) {

}
