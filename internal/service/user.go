package service

import (
	"errors"
	"server/internal/core"
	"server/internal/repository"
	"strings"
)

const goodSymbols = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_-.0123456789"

type UserService struct {
	repo repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{repo: repository}
}
func (s *UserService) RegisterUser(id int, user core.User) (core.User, error) {
	if ok, err := s.validateUserFields(user); err != nil {
		return core.User{}, err
	} else if !ok {
		return core.User{}, core.ErrIncorrectBody
	}

	return s.repo.Change(id, user), nil
}

func (s *UserService) validateUserFields(user core.User) (bool, error) {
	_, err := s.repo.GetByUsername(user.Username)
	if err == nil {
		return false, errors.New("username already exists")
	}
	if !validateName(user.Username) {
		return false, errors.New("invalid username")
	}
	if !validateName(user.Nickname) {
		return false, errors.New("invalid nickname")
	}
	return true, nil
}

func validateName(username string) bool {
	for _, s := range username {
		if !strings.ContainsRune(goodSymbols, s) {
			return false
		}
	}
	return true
}
