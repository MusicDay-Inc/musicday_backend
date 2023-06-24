package core

import (
	"github.com/google/uuid"
	"time"
)

var zeroTime = time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)

type SongDAO struct {
	Id           uuid.UUID `db:"id"`
	Author       string    `db:"author"`
	Name         string    `db:"name"`
	Date         time.Time `db:"date"`
	DurationTime time.Time `db:"duration"`
	AuthorId     uuid.UUID `db:"author_id"`
}

type SongWithReviewDAO struct {
	SongDAO
	ReviewNullableDAO
	SongId   uuid.UUID `db:"song_id"`
	ReviewId uuid.UUID `db:"review_id"`
}

func (s *SongDAO) ToDomain() (res Song) {
	res.Id = s.Id
	res.Author = s.Author
	res.Name = s.Name
	res.Date = s.Date
	res.Duration = s.DurationTime.Sub(zeroTime)
	//res.DurationTime = s.DurationTime
	return
}

type AlbumDAO struct {
	Id           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Author       string    `db:"author"`
	Date         time.Time `db:"date"`
	SongAmount   int       `db:"song_amount"`
	DurationTime time.Time `db:"duration"`
	AuthorId     string    `db:"author_id"`
}
type AlbumWithReviewDAO struct {
	AlbumDAO
	ReviewNullableDAO
	AlbumId  uuid.UUID `db:"album_id"`
	ReviewId uuid.UUID `db:"review_id"`
}

func (a *AlbumDAO) ToDomain() (res Album) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.DurationTime.Sub(zeroTime)
	return
}
