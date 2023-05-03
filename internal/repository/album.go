package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"server/internal/core"
)

type AlbumRepository struct {
	db *sqlx.DB
}

func (r AlbumRepository) GetSongsFromAlbum(id uuid.UUID) ([]core.SongDAO, error) {
	q := `
	SELECT id, author, name, date, duration, author_id
	from (SELECT song_id FROM album_songs WHERE album_id = $1) sq
         JOIN songs s on id = sq.song_id;
	`
	var (
		songs []core.SongDAO
		err   error
	)
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&songs, q, id)
	if err != nil {
		logrus.Error(err)
		return songs, err
	}
	return songs, nil
}

func NewAlbumRepository(database *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: database}
}

func (r AlbumRepository) GetById(albumId uuid.UUID) (album core.AlbumDAO, err error) {
	q := `
	SELECT * FROM albums WHERE id = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&album, q, albumId)
	if err != nil {
		return album, err
	}
	return
}
