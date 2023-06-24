package transport

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
	if !bindRequestBody(c, &t) {
		return
	}

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
		return
	}

	c.JSON(http.StatusOK, jwt.ToResponse())
}

func (h *Handler) signUp(c *gin.Context) {
	type req struct {
		Username string `json:"username" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		core.JWT
	}
	var rBody req
	if !bindRequestBody(c, &rBody) {
		return
	}
	id, registered, err := h.services.ParseToken(rBody.JWT.Token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, core.CodeTokenInvalid, err.Error())
		return
	}
	if registered {
		newErrorResponse(c, http.StatusBadRequest, core.CodeAccessDenied, "user already registered")
		return
	}
	newUser, err := h.services.User.RegisterUser(id, core.User{Username: rBody.Username, Nickname: rBody.Nickname})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, err.Error())
		return
	}

	jwt, err := h.services.GenerateJWT(newUser.Id, true)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		logrus.Errorf("while generating JWT" + err.Error())
		return
	}

	type response struct {
		core.UserPayload
		core.JWT
	}

	c.JSON(http.StatusOK, response{UserPayload: newUser.ToPayload(), JWT: jwt})
}
