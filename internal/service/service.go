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
	Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error)
	ChangeUsername(clientId uuid.UUID, username string) (core.User, error)
	ChangeNickname(clientId uuid.UUID, nickname string) (core.User, error)
	SearchUsers(query string, clientId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	Exists(id uuid.UUID) bool
	GetById(id uuid.UUID) (core.UserDTO, error)
	SubscriptionExists(clientId uuid.UUID, userId uuid.UUID) bool
	Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error)
	GetSubscribers(userId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	GetSubscriptions(userId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	GetBio(userId uuid.UUID) (string, error)
	CreateBio(clientId uuid.UUID, bio string) (string, error)
	UploadAvatar(clientId uuid.UUID) (core.UserDTO, error)
	AddPlayerID(clientId uuid.UUID, playerID uuid.UUID) error
	GetPlayerID(userId uuid.UUID) (string, error)
	InitPlayerID() error
}

type Song interface {
	GetById(songId uuid.UUID) (core.Song, error)
	SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.SongWithReviewDTO, error)
}

type Album interface {
	GetById(songId uuid.UUID) (core.Album, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error)
	SearchAlbumsWithReview(query string, userId uuid.UUID, limit int, offset int) ([]core.AlbumWithReviewDTO, error)
	GetCoverId(srcId uuid.UUID) (uuid.UUID, error)
}

type Review interface {
	//GetById(id uuid.UUID) (core.Review, error)
	GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error)
	PostReview(review core.Review) (core.ReviewDTO, error)
	GetSubscriptionReviews(releaseId uuid.UUID, clientId uuid.UUID, limit int, offset int) ([]core.ReviewOfUserDTO, error)
	DeleteReviewFromUser(userId uuid.UUID, reviewId uuid.UUID) (core.ReviewDTO, error)
	GetAllUserReviews(userId uuid.UUID, limit int, offset int, param string, orderParam string) ([]core.ReviewDTO, error)
	GetSongReviewsOfUser(userId uuid.UUID, limit int, offset int, param string, orderParam string) ([]core.ReviewDTO, error)
	GetAlbumReviewsOfUser(userId uuid.UUID, limit int, offset int, param string, orderParam string) ([]core.ReviewDTO, error)
	GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) ([]core.ReviewOfUserDTO, error)
	CountSongReviewsOf(userId uuid.UUID) (int32, error)
	CountAlbumReviewsOf(userId uuid.UUID) (int32, error)
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
		Review: NewReviewService(repos.Review, repos.Song, repos.Album, repos.User),
	}
}
