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

func (r SongRepository) SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) (songs []core.SongWithReviewDAO, err error) {
	q := `
	SELECT *
	FROM songs
         LEFT JOIN reviews on (songs.id = reviews.release_id AND reviews.user_id = $1)
	WHERE name ILIKE $2 || '%'
	ORDER BY songs.name
	LIMIT $3 OFFSET $4;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&songs, q, userId, searchReq, limit, offset)
	if err != nil {
		logrus.Error(err)
		return songs, err
	}
	return songs, nil
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
