package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
	"strconv"
)

func (h *Handler) getSongById(c *gin.Context) {
	songId := h.parseUUIDFromParam(c)

	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}

	s, err := h.services.Song.GetById(songId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, core.ErrNotFound.Error())
		return
	}

	review, err := h.services.Review.GetReviewToRelease(s.Id, userId)
	//c.JSON(http.StatusOK, core.SongWithReviewDTO{
	//	SongDTO:   s.ToDTO(),
	//	ReviewDTO: review.ToEmptyDTO(),
	//})

	c.JSON(http.StatusOK, core.SongWithReviewDTO{
		SongDTO:   s.ToDTO(),
		ReviewDTO: review.ToEmptyDTO(),
	})
}

func (h *Handler) SearchSongs(c *gin.Context) {
	//c.Get()
	query := c.Query("query")
	limitP := c.Query("limit")
	offsetP := c.Query("offset")
	limit, err := strconv.Atoi(limitP)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "couldn't get limit from parameter")
		return
	}
	offset, err := strconv.Atoi(offsetP)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "couldn't get offset from parameter")
		return
	}
	if limit > 50 {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "limit is too big")
		return
	}
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	if len(query) > 510 {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "search string is too long")
		return
	}
	res, err := h.services.Song.SearchSongsWithReview(query, userId, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	c.JSON(http.StatusOK, res)
}
