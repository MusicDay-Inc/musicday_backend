package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"server/internal/core"
)

type SongRepository struct {
	db *sqlx.DB
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r SongRepository) GetById(songId uuid.UUID) (song core.SongDAO, err error) {
	q := `
	SELECT * FROM songs WHERE id = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&song, q, songId)
	if err != nil {
		return song, err
	}
	return
}
