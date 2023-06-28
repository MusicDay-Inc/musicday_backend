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

func (r AlbumRepository) GetContainingSong(songId uuid.UUID) (id uuid.UUID, err error) {
	q := `
	SELECT album_id FROM album_songs WHERE song_id = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&id, q, songId)
	if err != nil {
		return id, err
	}
	return
}

func (r AlbumRepository) SearchAlbumsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.AlbumWithReview, error) {
	var (
		albums []core.AlbumWithReviewDAO
		err    error
		res    []core.AlbumWithReview
	)
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
		return res, err
	}
	res = make([]core.AlbumWithReview, len(albums))
	for i, awr := range albums {
		res[i] = awr.ToDomain()
		res[i].Album.Id = awr.AlbumId
		res[i].Review.Id = awr.ReviewId
	}
	return res, nil
}

func (r AlbumRepository) GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error) {
	var (
		songs []core.SongDAO
		err   error
	)
	q := `
	SELECT id, author, name, date, duration, author_id
	from (SELECT song_id FROM album_songs WHERE album_id = $1) sq
         JOIN songs s on id = sq.song_id;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&songs, q, id)
	if err != nil {
		logrus.Error(err)
		return []core.Song{}, err
	}
	res := make([]core.Song, len(songs))
	for i, v := range songs {
		res[i] = v.ToDomain()
	}
	return res, nil
}

func NewAlbumRepository(database *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: database}
}

func (r AlbumRepository) GetById(albumId uuid.UUID) (core.Album, error) {
	var (
		album core.AlbumDAO
		err   error
	)
	q := `
	SELECT * FROM albums WHERE id = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&album, q, albumId)
	if err != nil {
		return core.Album{}, err
	}
	return album.ToDomain(), nil
}
