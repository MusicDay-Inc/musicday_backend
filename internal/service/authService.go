package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/oauth2/v2"
	"net/http"
)

var (
	httpClient = &http.Client{}
	GoogleAuth *oauth2.Service
)

func GetGmail(idToken string) (string, error) {
	tokenInfo, err := verifyIdToken(idToken)
	if err != nil {
		return "", ErrTokenInvalid
	}
	return tokenInfo.Email, nil
}

// ErrTokenInvalid TODO вынести в другое место
var (
	ErrTokenInvalid = errors.New("token invalid")
)

func verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	tokenInfoCall := GoogleAuth.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

type Oauth2Service struct {
	service *oauth2.Service
}

func initializeAuthService() {
	serv, err := oauth2.New(httpClient)
	if err != nil {
		logrus.Error(err)
	}
	GoogleAuth = serv
}
