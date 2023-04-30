package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWT struct {
	Token string `json:"jwt_token" binding:"required"`
}

func bindRequestBody(c *gin.Context, obj any) {
	if err := c.BindJSON(&obj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
