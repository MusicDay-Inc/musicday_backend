package transport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
	"strconv"
)

func (h *Handler) reviewRelease(c *gin.Context) {
	releaseId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	var reviewInput core.ReviewDTO
	if !bindRequestBody(c, &reviewInput) {
		return
	}
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

func (h *Handler) ReviewsOfSubscribers(c *gin.Context) {
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

	songId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	subReviews, err := h.services.Review.GetSubscriptionReviews(songId, userId, limit, offset)
	var sum int32
	for _, r := range subReviews {
		sum += r.Score
	}
	//c.JSON(http.StatusOK, subReviews)
	c.JSON(http.StatusOK, map[string]interface{}{
		"reviews":    subReviews,
		"mean_score": float32(sum) / float32(len(subReviews)),
	})
}
func (h *Handler) deleteReviewById(c *gin.Context) {
	reviewId := h.parseUUIDFromParam(c)
	userId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	err = h.services.Review.DeleteReviewFromUser(userId, reviewId)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "couldn't find this review")
			return
		}
		if errors.Is(err, core.ErrInternal) {
			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrIncorrectBody.Error())
		return
	}

}
