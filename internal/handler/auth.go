package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/service"
)

type Token struct {
	IdToken string `json:"id_token" binding:"required"`
}

func (h *Handler) start(c *gin.Context) {
	var t Token

	// TODO вынести отдельно
	if err := c.BindJSON(&t); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO вынести отдельно
	gmail, err := service.GetGmail(t.IdToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Info("invalid Google token")
		// TODO вынесли в DTO
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
