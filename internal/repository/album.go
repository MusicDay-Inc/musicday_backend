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

func (r AlbumRepository) SearchAlbumsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) (albums []core.AlbumWithReviewDAO, err error) {
	q := `
	SELECT albums.id AS "album_id", name, author, date, song_amount, duration, author_id, 
	       reviews.id AS "review_id", user_id, is_song_reviewed, release_id, published_at, score, review_text
	FROM albums
         LEFT JOIN reviews on (albums.id = reviews.release_id AND reviews.user_id = $1)
	WHERE name ILIKE $2 || '%'
	ORDER BY albums.name
	LIMIT $3 OFFSET $4;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&albums, q, userId, searchReq, limit, offset)
	if err != nil {
		logrus.Error(err)
		return albums, err
	}
	for i, awr := range albums {
		albums[i].AlbumDAO.Id = awr.AlbumId
		albums[i].ReviewNullableDAO.Id = awr.ReviewId
	}
	return albums, nil
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
