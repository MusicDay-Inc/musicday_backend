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

func (s *Song) ToDTO() (res SongDTO) {
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

func (a *Album) ToDTO() (res AlbumDTO) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.Duration
	return
}

func (a *Album) ToFullDTO(s []Song) (res AlbumDTO) {
	res.Id = a.Id
	res.Name = a.Name
	res.Author = a.Author
	res.Date = a.Date
	res.SongAmount = a.SongAmount
	res.Duration = a.Duration
	res.Songs = make([]SongDTO, len(s))
	for i, song := range s {
		res.Songs[i] = song.ToDTO()
	}
	return
}

// TODO DELETE, ONLY REVIEW WILL CONTAIN SONG AND ALBUM
type Release struct {
	Id         uuid.UUID
	IsAlbum    bool
	Name       string
	Author     string
	Date       time.Time
	SongAmount int
	Duration   time.Duration
}

//func (r *Release) ToSmallDTO() (res ReleaseDTO) {
//	res.Id = r.Id
//	res.Name = r.Name
//	res.Author = r.Author
//	res.Date = r.Date
//	res.SongAmount = r.SongAmount
//	res.Duration = r.Duration
//	return
//}
//
//func (r *Release) ToFullDTO(s []Song) (res ReleaseDTO) {
//	res.Id = r.Id
//	res.IsAlbum = r.IsAlbum
//	res.Name = r.Name
//	res.Author = r.Author
//	res.Date = r.Date
//	res.SongAmount = r.SongAmount
//	res.Duration = r.Duration
//	res.Songs = make([]SongDTO, len(s))
//	for i, song := range s {
//		res.Songs[i] = song.ToDTO()
//	}
//	return
//}
