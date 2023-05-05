package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

func (h *Handler) getSongById(c *gin.Context) {
	songId := h.parseUUIDFromParam(c)

	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
	}

	s, err := h.services.Song.GetById(songId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}

	review, err := h.services.GetReviewToRelease(s.Id, userId)
	type response struct {
		core.SongDTO   `json:"song,omitempty"`
		core.ReviewDTO `json:"review,omitempty"`
	}
	c.JSON(http.StatusOK, response{
		SongDTO:   s.ToDTO(),
		ReviewDTO: review.ToEmptyDTO(),
	})
}

func (h *Handler) SearchSongs(c *gin.Context) {

}
