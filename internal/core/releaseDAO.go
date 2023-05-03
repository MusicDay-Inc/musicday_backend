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

func (s *SongDAO) ToReleaseDomain() (r Release) {
	r.Id = s.Id
	r.IsAlbum = false
	r.Name = s.Name
	r.Author = s.Author
	r.Date = s.Date
	r.SongAmount = -1
	r.Duration = s.DurationTime.Sub(zeroTime)
	//r.DurationTime = s.DurationTime

	return
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

func (a *AlbumDAO) ToReleaseDomain() (r Release) {
	r.Id = a.Id
	r.IsAlbum = true
	r.Name = a.Name
	r.Author = a.Author
	r.Date = a.Date
	r.SongAmount = a.SongAmount
	r.Duration = a.DurationTime.Sub(zeroTime)
	return
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
