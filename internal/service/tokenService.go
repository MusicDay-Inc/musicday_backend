package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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
	UserId       int  `json:"user_id"`
	IsRegistered bool `json:"is_registered"`
}

func initializeSecret() {
	signingKey = viper.GetString("signing_key")
}

func NewTokenService(userRepo repository.User) *TokenService {
	initializeSecret()
	return &TokenService{userRepo: userRepo}
}

func (s *TokenService) GetJWT(gmail string) (string, error) {
	if !s.userRepo.Exists(gmail) {
		userId, err := s.userRepo.Create(gmail)
		if err != nil {
			return "", err
		}
		return GenerateJWT(userId, false)
	}
	user, err := s.userRepo.GetByGmail(gmail)
	if err != nil {
		return "", err
	}
	return GenerateJWT(user.Id, user.IsRegistered)
}

func GenerateJWT(userId int, isRegistered bool) (string, error) {
	if isRegistered {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().AddDate(registeredTTLYears, 0, 0).Unix(),
			},
			userId,
			true,
		})
		return token.SignedString([]byte(signingKey))
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * unregisteredTTLMinutes).Unix(),
			},
			userId,
			false,
		})
		return token.SignedString([]byte(signingKey))
	}

}

func (s *TokenService) ParseToken(token string) (int, bool, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, false, err
	}

	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return 0, false, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.IsRegistered, nil
}
