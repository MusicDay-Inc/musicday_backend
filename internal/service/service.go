package service

import (
	"server/internal/core"
	"server/internal/repository"
)

type Token interface {
	GetJWT(gmail string) (core.JWT, error)
	ParseToken(token string) (int, error)
}

type User interface {
	RegisterUser(userId int, user core.User) (core.User, error)
}

type Service struct {
	Token
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Token: NewTokenService(repos.User),
		User:  NewUserService(repos.User),
	}
}
