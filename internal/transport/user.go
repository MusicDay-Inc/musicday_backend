package transport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
	"strconv"
)

// GET
func (h *Handler) getUserSubscribers(c *gin.Context) {
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
	userId := h.parseUUIDFromParam(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	subscribers, err := h.services.User.GetSubscribers(userId, limit, offset)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "couldn't find this user")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subscribers)
}

func (h *Handler) getUserSubscriptions(c *gin.Context) {
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
	userId := h.parseUUIDFromParam(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	subscribers, err := h.services.User.GetSubscriptions(userId, limit, offset)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "couldn't find this user")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subscribers)
}

func (h *Handler) getUserProfile(c *gin.Context) {
	clientId, err := h.getClientId(c)
	userId := h.parseUUIDFromParam(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	ok := h.services.User.Exists(userId)
	if !ok {
		newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "couldn't find this user")
		return
	}
	user, err := h.services.User.GetById(userId)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "couldn't find this user")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	isSubscribed := h.services.User.SubscriptionExists(clientId, userId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":                 user,
		"is_client_subscribed": isSubscribed,
	})
}

func (h *Handler) SearchUsers(c *gin.Context) {
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
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	if len(query) > 30 {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "search string is too long")
		return
	}
	res, err := h.services.User.SearchUsers(query, clientId, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "server search error")
		return
	}
	c.JSON(http.StatusOK, res)
}

// POST

func (h *Handler) changeUsername(c *gin.Context) {
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	var u core.UserDTO
	if !bindRequestBody(c, &u) {
		return
	}
	newUser, err := h.services.User.ChangeUsername(clientId, u.Username)
	if err != nil {
		if errors.Is(err, core.ErrInternal) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, err.Error())
		return
	}
	c.JSON(http.StatusOK, newUser.ToDTO())
}

func (h *Handler) changeNickname(c *gin.Context) {
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	var u core.UserDTO
	if !bindRequestBody(c, &u) {
		return
	}
	newUser, err := h.services.User.ChangeNickname(clientId, u.Nickname)
	if err != nil {
		if errors.Is(err, core.ErrInternal) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, err.Error())
		return
	}
	c.JSON(http.StatusOK, newUser.ToDTO())
}

func (h *Handler) subscribe(c *gin.Context) {
	userId := h.parseUUIDFromParam(c)
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	updatedUserSubscription, err := h.services.User.Subscribe(clientId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "already subscribed")
		return
	}
	//c.JSON(http.StatusOK, updatedUserSubscription)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":                 updatedUserSubscription,
		"is_client_subscribed": true,
	})
}
func (h *Handler) unsubscribe(c *gin.Context) {
	userId := h.parseUUIDFromParam(c)
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	updatedUserSubscription, err := h.services.User.Unsubscribe(clientId, userId)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeNotFound, "already unsubscribed")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":                 updatedUserSubscription,
		"is_client_subscribed": false,
	})
}
