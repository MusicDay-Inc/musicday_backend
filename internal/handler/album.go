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
	userId := h.parseUUIDFromParam(c)
	a, err := h.services.Album.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	songs, err := h.services.Album.GetSongsFromAlbum(userId)
	review, err := h.services.GetReviewToRelease(a.Id, userId)
	type response struct {
		core.AlbumDTO  `json:"album,omitempty"`
		core.ReviewDTO `json:"review,omitempty"`
	}
	c.JSON(http.StatusOK, response{
		AlbumDTO:  a.ToFullDTO(songs),
		ReviewDTO: review.ToEmptyDTO(),
	})
}

func (h *Handler) SearchAlbums(context *gin.Context) {

}
