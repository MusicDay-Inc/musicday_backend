package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type Token interface {
	GetJWT(gmail string) (core.JWT, error)
	ParseToken(token string) (uuid.UUID, bool, error)
	GenerateJWT(userId uuid.UUID, registered bool) (core.JWT, error)
}

type User interface {
	RegisterUser(userId uuid.UUID, user core.User) (core.User, error)
}

type Song interface {
	GetById(songId uuid.UUID) (core.Song, error)
}

type Album interface {
	GetById(songId uuid.UUID) (core.Album, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error)
}

type Service struct {
	Token
	User
	Song
	Album
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Token: NewTokenService(repos.User),
		User:  NewUserService(repos.User),
		Song:  NewSongService(repos.Song),
		Album: NewAlbumService(repos.Album),
	}
}
