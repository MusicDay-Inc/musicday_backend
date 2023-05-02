package core

import (
	"github.com/google/uuid"
	"time"
)

type SongDAO struct {
	Id       uuid.UUID     `db:"id"`
	Name     string        `db:"name"`
	Author   string        `db:"author"`
	Date     time.Time     `db:"date"`
	Duration time.Duration `db:"duration"`
}

func (s *SongDAO) ToReleaseDomain() (r Release) {
	r.Id = s.Id
	r.IsAlbum = false
	r.Name = s.Name
	r.Author = s.Author
	r.Date = s.Date
	r.SongAmount = -1
	r.Duration = s.Duration
	return
}
func (s *SongDAO) ToDomain() (res Song) {
	res.Id = s.Id
	res.Name = s.Name
	res.Author = s.Author
	res.Date = s.Date
	res.Duration = s.Duration
	return
}

type AlbumDAO struct {
	Id         uuid.UUID     `db:"id"`
	Name       string        `db:"name"`
	Author     string        `db:"author"`
	AuthorId   string        `db:"author_id"`
	Date       time.Time     `db:"date"`
	SongAmount int           `db:"song_amount"`
	Duration   time.Duration `db:"duration"`
}

func (a *AlbumDAO) ToReleaseDomain() (r Release) {
	r.Id = a.Id
	r.IsAlbum = true
	r.Name = a.Name
	r.Author = a.Author
	r.Date = a.Date
	r.SongAmount = a.SongAmount
	r.Duration = a.Duration
	return
}

func (a *AlbumDAO) ToDomain() (res Album) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.Duration
	return
}
