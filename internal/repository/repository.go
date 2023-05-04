package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"server/internal/core"
)

//type ReleaseItem interface {
//	GetSongById(song uuid.UUID) (core.SongDAO, error)
//	GetAlbumById(album uuid.UUID) (core.AlbumDAO, error)
//}

type User interface {
	// Create returns id of new user, and changes his id
	Create(gmail string) (uuid.UUID, error)
	Exists(gmail string) bool
	GetById(userId uuid.UUID) (user core.UserDAO, err error)
	GetByUsername(username string) (core.UserDAO, error)
	GetByGmail(gmail string) (user core.UserDAO, err error)
	Register(u core.User) (user core.UserDAO, err error)
	ChangeUsername(u core.User) (user core.UserDAO, err error)
	ChangeNickname(u core.User) (user core.UserDAO, err error)
	InstallPicture(id uuid.UUID) (user core.UserDAO, err error)
}

type Song interface {
	GetById(songId uuid.UUID) (core.SongDAO, error)
}

type Album interface {
	GetById(album uuid.UUID) (core.AlbumDAO, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.SongDAO, error)
}

type Review interface {
	GetById(id uuid.UUID) (core.ReviewDAO, error)
	GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.ReviewDAO, error)
}

type Repository struct {
	//ReleaseItem
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
