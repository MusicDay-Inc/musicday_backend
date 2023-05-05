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
	SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.SongReviewDTO, error)
}

type Album interface {
	GetById(songId uuid.UUID) (core.Album, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error)
}

type Review interface {
	//GetById(id uuid.UUID) (core.Review, error)
	GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error)
	PostReview(review core.Review) (core.ReviewDTO, error)
}

type Service struct {
	Token
	User
	Song
	Album
	Review
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Token:  NewTokenService(repos.User),
		User:   NewUserService(repos.User),
		Song:   NewSongService(repos.Song),
		Album:  NewAlbumService(repos.Album),
		Review: NewReviewService(repos.Review, repos.Song, repos.Album),
	}
}
