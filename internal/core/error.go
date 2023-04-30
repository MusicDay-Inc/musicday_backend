package core

import (
	"errors"
)

type ErrorBody struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var (
	ErrTokenInvalid  = errors.New("token invalid")
	ErrIncorrectBody = errors.New("incorrect json body")
	ErrInternal      = errors.New("server internal error")
)

const (
	CodeTokenInvalid  = 1
	CodeIncorrectBody = 2
	CodeInternalError = 3
)
