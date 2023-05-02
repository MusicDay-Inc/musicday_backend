package core

import (
	"github.com/google/uuid"
	"time"
)

type SongDAO struct {
	Id       uuid.UUID     `db:"id"`
	Name     string        `db:"song_name"`
	Author   string        `db:"song_author"`
	Date     time.Time     `db:"song_date"`
	Duration time.Duration `db:"song_duration"`
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
	Name       string        `db:"album_name"`
	Author     string        `db:"album_author"`
	Date       time.Time     `db:"album_date"`
	SongAmount int           `db:"song_amount"`
	Duration   time.Duration `db:"album_duration"`
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
