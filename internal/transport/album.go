package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
	"strconv"
)

func (h *Handler) getAlbumById(c *gin.Context) {
	albumId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
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
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
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

func (h *Handler) SearchAlbums(c *gin.Context) {
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
	res, err := h.services.Album.SearchAlbumsWithReview(query, userId, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	c.JSON(http.StatusOK, res)
}
