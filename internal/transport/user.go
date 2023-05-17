package transport

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tbalthazar/onesignal-go"
	"net/http"
	"server/internal/core"
	"strconv"
)

// GET

func (h *Handler) getUserActivityFeed(c *gin.Context) {
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

	subscribers, err := h.services.Review.GetReviewsOfUserSubscriptions(clientId, limit, offset)
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
	sAmount, err := h.services.Review.CountSongReviewsOf(userId)
	if err != nil {
		sAmount = 0
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	aAmount, err := h.services.Review.CountAlbumReviewsOf(userId)
	if err != nil {
		aAmount = 0
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	bio, err := h.services.User.GetBio(userId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":                 user,
		"is_client_subscribed": isSubscribed,
		"bio":                  bio,
		"song_amount":          sAmount,
		"album_amount":         aAmount,
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
		if errors.Is(err, core.ErrAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeAlreadyExists, "already subscribed")
			return
		}
		if errors.Is(err, core.ErrIncorrectBody) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "can't subscribe to yourself")
			return
		}
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, core.ErrInternal.Error())
		return
	}
	playerID, err := h.services.User.GetPlayerID(userId)
	user, errUser := h.services.User.GetById(clientId)
	if err == nil && errUser == nil {
		CreateNotification(playerID, user)
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
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, "incorrect subscription id")
			return
		}
		if errors.Is(err, core.ErrAlreadyExists) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeAlreadyExists, "already unsubscribed")
			return
		}
		if errors.Is(err, core.ErrIncorrectBody) {
			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "can't unsubscribe to yourself")
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
func (h *Handler) CreateClientBio(c *gin.Context) {
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	type userBio struct {
		Bio string `json:"bio" binding:"required"`
	}
	var bio userBio
	if !bindRequestBody(c, &bio) {
		return
	}

	resBio, err := h.services.User.CreateBio(clientId, bio.Bio)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeAlreadyExists, "bio already exists")
		return
	}
	bio.Bio = resBio
	c.JSON(http.StatusOK, bio)
}

func (h *Handler) postPlayerId(c *gin.Context) {
	playerID := h.parseUUIDFromParam(c)
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}

	err = h.services.User.AddPlayerID(clientId, playerID)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, core.CodeNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}

	c.JSON(http.StatusOK, clientId)
}
func CreateNotification(playerID string, user core.UserDTO) {
	client := onesignal.NewClient(nil)
	// TODO API KEY
	client.AppKey = "YOUR API KEY"
	CreateNotificationHelper(client, playerID, user)
}
func CreateNotificationHelper(client *onesignal.Client, playerID string, dto core.UserDTO) *onesignal.NotificationCreateResponse {
	notificationReq := &onesignal.NotificationRequest{
		AppID:            "8af60ff7-a3f4-4c99-8658-3fbe8538cdb9",
		Headings:         map[string]string{"en": "New Subscription !"},
		Contents:         map[string]string{"en": dto.Username + "has subscribed to you"},
		IncludePlayerIDs: []string{playerID},
	}
	createRes, _, err := client.Notifications.Create(notificationReq)
	if err != nil {
		fmt.Println(err)
	} else {
		return createRes
	}
	return createRes
}
