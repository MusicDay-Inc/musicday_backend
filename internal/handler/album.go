package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

func (h *Handler) getAlbumById(c *gin.Context) {
	albumId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
		return
	}

	a, err := h.services.Album.GetById(albumId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	review, err := h.services.GetReviewToRelease(a.Id, userId)
	c.JSON(http.StatusOK, core.AlbumReviewDTO{
		AlbumDTO:  a.ToDTO(),
		ReviewDTO: review.ToEmptyDTO(),
	})
}

func (h *Handler) getAlbumWitSongsById(c *gin.Context) {
	albumId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
		return
	}

	a, err := h.services.Album.GetById(albumId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}
	songs, err := h.services.Album.GetSongsFromAlbum(albumId)
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
