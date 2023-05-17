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
	r repository.User
}

func (s *UserService) InitPlayerID() error {
	return s.r.InitPlayerID()
}

func (s *UserService) GetPlayerID(userId uuid.UUID) (string, error) {
	res, err := s.r.GetPlayerID(userId)
	if err != nil {
		return "", core.ErrNotFound
	}
	return res, nil
}

func (s *UserService) AddPlayerID(clientId uuid.UUID, playerID uuid.UUID) error {
	if !s.r.ExistsWithId(clientId) {
		return core.ErrNotFound
	}
	err := s.r.InstallAppID(clientId, playerID)
	return err
}

func (s *UserService) UploadAvatar(clientId uuid.UUID) (core.UserDTO, error) {
	user, err := s.r.InstallPicture(clientId)
	if err != nil {
		userPrev, _ := s.r.GetById(clientId)
		return userPrev.ToDTO(), err
	}
	res := user.ToDomain()
	return res.ToDTO(), nil
}

func (s *UserService) CreateBio(clientId uuid.UUID, bio string) (string, error) {
	resBio, err := s.r.CreateBio(clientId, bio)
	if err != nil {
		return "", core.ErrAlreadyExists
	}
	return resBio, nil
}

func (s *UserService) GetBio(userId uuid.UUID) (string, error) {
	bio, err := s.r.GetBio(userId)
	if err != nil {
		return "", core.ErrNotFound
	}
	return bio, nil
}

func (s *UserService) GetSubscriptions(userId uuid.UUID, limit int, offset int) (res []core.UserDTO, err error) {
	ok := s.r.ExistsWithId(userId)
	if !ok {
		return []core.UserDTO{}, core.ErrNotFound
	}
	users, err := s.r.GetSubscriptionsOf(userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.UserDTO, len(users))
	for i, user := range users {
		uDomain := user.ToDomain()
		res[i] = uDomain.ToDTO()
	}
	return res, nil
}

func (s *UserService) GetSubscribers(userId uuid.UUID, limit int, offset int) (res []core.UserDTO, err error) {
	ok := s.r.ExistsWithId(userId)
	if !ok {
		return []core.UserDTO{}, core.ErrNotFound
	}
	users, err := s.r.GetSubscribers(userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.UserDTO, len(users))
	for i, user := range users {
		uDomain := user.ToDomain()
		res[i] = uDomain.ToDTO()
	}
	return res, nil
}

func (s *UserService) SubscriptionExists(clientId uuid.UUID, userId uuid.UUID) bool {
	return s.r.IsSubscriptionExists(clientId, userId)
}

func (s *UserService) GetById(id uuid.UUID) (core.UserDTO, error) {
	user, err := s.r.GetById(id)
	return user.ToDTO(), err
}

func (s *UserService) Exists(id uuid.UUID) bool {
	return s.r.ExistsWithId(id)
}

func (s *UserService) SearchUsers(query string, clientId uuid.UUID, limit int, offset int) (res []core.UserDTO, err error) {
	users, err := s.r.SearchUsers(query, clientId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.UserDTO, len(users))
	for i, user := range users {
		uDomain := user.ToDomain()
		res[i] = uDomain.ToDTO()
	}
	return res, nil
}

func (s *UserService) ChangeNickname(clientId uuid.UUID, nickname string) (core.User, error) {
	if !s.validateNickname(nickname) {
		return core.User{}, errors.New("username is invalid")
	}
	newClient, err := s.r.ChangeNickname(clientId, nickname)
	if err != nil {
		return core.User{}, core.ErrInternal
	}
	return newClient.ToDomain(), nil
}

func (s *UserService) ChangeUsername(clientId uuid.UUID, username string) (core.User, error) {
	tmpUser, err := s.r.GetByUsername(username)
	if err == nil && tmpUser.IsRegistered {
		return core.User{}, errors.New("user with this nickname already exists")
	}
	if !s.validateUsername(username) {
		return core.User{}, errors.New("username is invalid")
	}
	newClient, err := s.r.ChangeUsername(clientId, username)
	if err != nil {
		return core.User{}, core.ErrInternal
	}
	return newClient.ToDomain(), nil
}

func (s *UserService) Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error) {
	user, err := s.r.GetById(userId)
	if err != nil {
		return user.ToDTO(), core.ErrNotFound
	}
	if clientId == userId {
		return user.ToDTO(), core.ErrIncorrectBody
	}
	ok := !s.r.IsSubscriptionExists(clientId, userId)
	if !ok {
		return user.ToDTO(), core.ErrAlreadyExists
	}
	updatedUser, err := s.r.Subscribe(clientId, userId)
	if err != nil {
		return core.UserDTO{}, err
	}
	return updatedUser.ToDTO(), nil
}

func (s *UserService) Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error) {
	user, err := s.r.GetById(userId)
	if err != nil {
		return user.ToDTO(), core.ErrNotFound
	}
	if clientId == userId {
		return user.ToDTO(), core.ErrIncorrectBody
	}
	ok := s.r.IsSubscriptionExists(clientId, userId)
	if !ok {
		return user.ToDTO(), core.ErrAlreadyExists
	}
	updatedUser, err := s.r.Unsubscribe(clientId, userId)
	if err != nil {
		return core.UserDTO{}, err
	}
	return updatedUser.ToDTO(), nil
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{r: repository}
}
func (s *UserService) RegisterUser(id uuid.UUID, user core.User) (core.User, error) {
	if ok, err := s.validateUserFields(user); err != nil {
		return core.User{}, err
	} else if !ok {
		return core.User{}, core.ErrIncorrectBody
	}

	user.Id = id
	u, err := s.r.Register(user)
	return u.ToDomain(), err
}

func (s *UserService) validateUserFields(user core.User) (bool, error) {
	_, err := s.r.GetByUsername(user.Username)
	if err == nil {
		return false, errors.New("username already exists")
	}
	if !s.validateUsername(user.Username) {
		return false, errors.New("invalid username")
	}
	if !s.validateNickname(user.Nickname) {
		return false, errors.New("invalid nickname")
	}
	return true, nil
}

func (s *UserService) validateUsername(username string) bool {
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
func (s *UserService) validateNickname(nickname string) bool {
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
