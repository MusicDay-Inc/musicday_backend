package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/core"
	"server/internal/service"
)

type idToken struct {
	IdToken string `json:"id_token" binding:"required"`
}

func (h *Handler) start(c *gin.Context) {
	var t idToken

	// TODO вынести отдельно ГОТОВО
	if err := c.BindJSON(&t); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO вынести отдельно НЕТ?
	gmail, err := service.GetGmail(t.IdToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Info("invalid Google token")
		// TODO вынесли в DTO DELETE
		c.JSON(http.StatusOK, map[string]interface{}{
			"jwt_token": "",
		})
		return
	}

	jwt, err := h.services.GetJWT(gmail)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Errorf("while generating JWT" + err.Error())
	}

	// TODO заменить на выдачу JWT токена
	// TODO вынесли в DTO
	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt_token": jwt,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	type req struct {
		core.User
		JWT
	}

	var requestBody req
	bindRequestBody(c, &requestBody)
	id, err := h.services.ParseToken(requestBody.JWT.Token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	newUser, err := h.services.User.RegisterUser(id, requestBody.User)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, newUser)
}
