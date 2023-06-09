package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"server/internal/core"
)

type User interface {
	// Create returns id of new user, and changes his id
	Create(gmail string) (uuid.UUID, error)
	Exists(gmail string) bool
	GetById(userId uuid.UUID) (user core.User, err error)
	GetByUsername(username string) (core.User, error)
	GetByGmail(gmail string) (user core.User, err error)
	Register(u core.User) (user core.User, err error)
	ChangeUsername(id uuid.UUID, username string) (user core.User, err error)
	ChangeNickname(id uuid.UUID, nickname string) (user core.User, err error)
	InstallPicture(id uuid.UUID) (user core.User, err error)
	Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error)
	SearchUsers(query string, clientId uuid.UUID, limit int, offset int) ([]core.User, error)
	ExistsWithId(id uuid.UUID) bool
	IsSubscriptionExists(clientId uuid.UUID, userId uuid.UUID) bool
	Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error)
	GetSubscribers(userId uuid.UUID, limit int, offset int) ([]core.User, error)
	GetSubscriptionsOf(userId uuid.UUID, limit int, offset int) ([]core.User, error)
	GetBio(userId uuid.UUID) (string, error)
	CreateBio(userId uuid.UUID, bio string) (string, error)
	InstallAppID(clientId uuid.UUID, playerID uuid.UUID) error
	GetPlayerID(userId uuid.UUID) (string, error)
	InitPlayerID() error
}

type Song interface {
	GetById(songId uuid.UUID) (core.Song, error)
	SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.SongWithReview, error)
}

type Album interface {
	GetById(album uuid.UUID) (core.Album, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error)
	SearchAlbumsWithReview(query string, userId uuid.UUID, limit int, offset int) ([]core.AlbumWithReview, error)
	GetContainingSong(songId uuid.UUID) (uuid.UUID, error)
}

type Review interface {
	GetById(id uuid.UUID) (core.Review, error)
	GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error)
	InsertReview(review core.Review) (core.Review, error)
	ExistsToRelease(userId uuid.UUID, releaseId uuid.UUID) (bool, error)
	ExistsFromUser(userId uuid.UUID, releaseId uuid.UUID) (bool, error)
	UpdateReview(review core.Review) (core.Review, error)
	GetSubscriptionReviews(releaseId uuid.UUID, clientId uuid.UUID) ([]core.Review, error)
	Delete(id uuid.UUID) error
	GetSongReviewsFromUser(userId uuid.UUID, limit int, offset int, param string) ([]core.Review, error)
	GetAlbumReviewsFromUser(userId uuid.UUID, limit int, offset int) ([]core.Review, error)
	GetReviewsFromUser(userId uuid.UUID, limit int, offset int, param string) ([]core.Review, error)
	GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) ([]core.Review, error)
	CountReviewsOfUser(userId uuid.UUID, isToSongs bool) (int32, error)
}

type Repository struct {
	User
	Song
	Album
	Review
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:   NewUserRepository(db),
		Song:   NewSongRepository(db),
		Album:  NewAlbumRepository(db),
		Review: NewReviewRepository(db),
	}
}
