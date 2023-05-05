package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
)

func (h *Handler) reviewRelease(c *gin.Context) {
	releaseId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get userId from ctx")
	}
	var reviewInput core.ReviewDTO
	bindRequestBody(c, &reviewInput)
	if len(reviewInput.Text) > 2000 {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "review text is too long")
		return
	}
	r, err := h.services.Review.PostReview(reviewInput.FormReview(releaseId, userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrIncorrectBody.Error())
		return
	}
	c.JSON(http.StatusOK, r)

}
