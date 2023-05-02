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
	bindRequestBody(c, &t)
	//if err := c.BindJSON(&t); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, CodeIncorrectBody, err.Error())
	//	return
	//}

	gmail, err := service.GetGmail(t.IdToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeTokenInvalid, err.Error())
		logrus.Info("invalid Google token")
		return
	}

	jwt, err := h.services.GetJWT(gmail)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		logrus.Errorf("while generating JWT" + err.Error())
	}

	c.JSON(http.StatusOK, jwt.ToResponse())
}

func (h *Handler) signUp(c *gin.Context) {
	type req struct {
		core.User
		core.JWT
	}

	var requestBody req
	bindRequestBody(c, &requestBody)
	id, err := h.services.ParseToken(requestBody.JWT.Token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, core.CodeTokenInvalid, err.Error())
		return
	}
	newUser, err := h.services.User.RegisterUser(id, requestBody.User)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, err.Error())
		return
	}
	c.JSON(http.StatusOK, newUser)
}
