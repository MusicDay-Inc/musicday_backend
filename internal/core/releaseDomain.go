package core

import (
	"github.com/google/uuid"
	"time"
)

type Song struct {
	Id       uuid.UUID
	Author   string
	Name     string
	Date     time.Time
	Duration time.Duration
}

// TODO add validation if will be used for search?
// or just ignore fields exept Id
func (s *Song) ToDAO() (res SongDAO) {
	res.Id = s.Id
	res.Name = s.Name
	res.Author = s.Author
	res.Date = s.Date
	res.DurationTime = zeroTime.Add(s.Duration)
	//res.DurationTime = s.DurationTime
	return
}

func (s *Song) ToPayload() (res SongPayload) {
	res.Id = s.Id
	res.Name = s.Name
	res.Author = s.Author
	res.Date = s.Date
	res.Duration = s.Duration
	return
}

type Album struct {
	Id         uuid.UUID
	Name       string
	Author     string
	Date       time.Time
	SongAmount int
	Duration   time.Duration
}

// TODO add validation if will be used for search?
// or just ignore fields exept Id
func (a *Album) ToDAO() (res AlbumDAO) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.DurationTime = zeroTime.Add(a.Duration)
	return
}

func (a *Album) ToPayload() (res AlbumPayload) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.Duration
	return
}

func (a *Album) ToFullPayload(s []Song) (res AlbumPayload) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.Duration
	res.Songs = make([]SongPayload, len(s))
	for i, song := range s {
		res.Songs[i] = song.ToPayload()
	}
	return
}
