package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"server/internal/core"
)

type ReleaseRepo struct {
	db *sqlx.DB
}

func (r ReleaseRepo) GetSongById(song uuid.UUID) (core.SongDAO, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReleaseRepo) GetAlbumById(album uuid.UUID) (core.AlbumDAO, error) {
	//TODO implement me
	panic("implement me")
}

func NewReleaseRepo(database *sqlx.DB) *ReleaseRepo {
	return &ReleaseRepo{db: database}
}
