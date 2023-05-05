package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"server/internal/core"
	"server/internal/repository"
	"time"
)

type TokenService struct {
	userRepo repository.User
}

var (
	signingKey = ""
)

const (
	registeredTTLYears     = 1
	unregisteredTTLMinutes = 15
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId       uuid.UUID `json:"user_id"`
	IsRegistered bool      `json:"is_registered"`
}

func initializeSecret() {
	signingKey = viper.GetString("signing_key")
}

func NewTokenService(userRepo repository.User) *TokenService {
	initializeSecret()
	return &TokenService{userRepo: userRepo}
}

func (s *TokenService) GetJWT(gmail string) (core.JWT, error) {
	if !s.userRepo.Exists(gmail) {
		userId, err := s.userRepo.Create(gmail)
		if err != nil {
			return core.JWT{}, err
		}
		return s.GenerateJWT(userId, false)
	}
	user, err := s.userRepo.GetByGmail(gmail)
	if err != nil {
		return core.JWT{}, err
	}
	return s.GenerateJWT(user.Id, user.IsRegistered)
}

func (s *TokenService) GenerateJWT(userId uuid.UUID, registered bool) (core.JWT, error) {
	if registered {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().AddDate(registeredTTLYears, 0, 0).Unix(),
			},
			userId,
			true,
		})
		str, err := token.SignedString([]byte(signingKey))
		return core.JWT{Token: str, IsRegistered: registered}, err
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * unregisteredTTLMinutes).Unix(),
			},
			userId,
			false,
		})
		str, err := token.SignedString([]byte(signingKey))
		return core.JWT{Token: str, IsRegistered: registered}, err
	}

}

func (s *TokenService) ParseToken(token string) (uuid.UUID, bool, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.UUID{0}, false, err
	}

	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return uuid.UUID{0}, false, errors.New("token claims are not of type *tokenClaims")
	}
	usr, err := s.userRepo.GetById(claims.UserId)
	if claims.IsRegistered != usr.IsRegistered {
		if err != nil {
			return uuid.UUID{0}, false, err
		}
		return uuid.UUID{0}, false, errors.New("registration flags with token and user do not match")
	}
	return claims.UserId, claims.IsRegistered, nil
}
