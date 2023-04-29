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
	userId, err := service.IdToken(t.IdToken)
	if err != nil {
		newErrorResponse(c, http.StatusNonAuthoritativeInfo, err.Error())
		logrus.Info("invalid Google token")
	}

	// TODO заменить на выдачу JWT токена
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})

	// TODO сделать добавление в БД

}
