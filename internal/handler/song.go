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
		return
	}

	s, err := h.services.Song.GetById(songId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}

	review, err := h.services.GetReviewToRelease(s.Id, userId)
	c.JSON(http.StatusOK, core.SongReviewDTO{
		SongDTO:   s.ToDTO(),
		ReviewDTO: review.ToEmptyDTO(),
	})
}

func (h *Handler) SearchSongs(c *gin.Context) {
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
		return
	}
	var searchInput core.SearchDTO
	if !bindRequestBody(c, &searchInput) {
		return
	}
	if len(searchInput.Request) > 510 {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "search string is too long")
		return
	}
	res, err := h.services.Song.SearchSongsWithReview(searchInput.Request, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	c.JSON(http.StatusOK, res)
}
