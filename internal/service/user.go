package service

import (
	"errors"
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
	"strings"
)

const UsernameSymbols = "qwertyuiopasdfghjklzxcvbnm_-.0123456789"
const NicknameSymbols = " qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_-.0123456789"

type UserService struct {
	repo repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{repo: repository}
}
func (s *UserService) RegisterUser(id uuid.UUID, user core.User) (core.User, error) {
	if ok, err := s.validateUserFields(user); err != nil {
		return core.User{}, err
	} else if !ok {
		return core.User{}, core.ErrIncorrectBody
	}

	user.Id = id
	u, err := s.repo.Register(user)
	return u.ToDomain(), err
}

func (s *UserService) validateUserFields(user core.User) (bool, error) {
	_, err := s.repo.GetByUsername(user.Username)
	if err == nil {
		return false, errors.New("username already exists")
	}
	if !validateUsername(user.Username) {
		return false, errors.New("invalid username")
	}
	if !validateNickname(user.Nickname) {
		return false, errors.New("invalid nickname")
	}
	return true, nil
}

func validateUsername(username string) bool {
	if len(username) < 5 || len(username) > 30 {
		return false
	}
	for _, s := range username {
		if !strings.ContainsRune(UsernameSymbols, s) {
			return false
		}
	}
	return true
}
func validateNickname(nickname string) bool {
	if len(nickname) < 5 || len(nickname) > 30 {
		return false
	}
	for _, s := range nickname {
		if !strings.ContainsRune(NicknameSymbols, s) {
			return false
		}
	}
	return true
}
