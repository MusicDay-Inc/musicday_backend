package core

import (
	"github.com/google/uuid"
	"time"
)

type SongDTO struct {
	Id       uuid.UUID     `json:"id"`
	Name     string        `json:"name"`
	Author   string        `json:"author"`
	Date     time.Time     `json:"date"`
	Duration time.Duration `json:"duration"`
}

type AlbumDTO struct {
	Id         uuid.UUID     `json:"id"`
	Name       string        `json:"name"`
	Author     string        `json:"author"`
	Date       time.Time     `json:"date"`
	SongAmount int           `json:"song_amount"`
	Duration   time.Duration `json:"duration"`
	Songs      []SongDTO     `json:"songs,omitempty"`
}

type ReleaseDTO struct {
	Id         uuid.UUID     `json:"id"`
	IsAlbum    bool          `json:"is_album"`
	Name       string        `json:"name"`
	Author     string        `json:"author"`
	Date       time.Time     `json:"date"`
	Duration   time.Duration `json:"duration"`
	SongAmount int           `json:"song_amount"`
	Songs      []SongDTO     `json:"songs,omitempty"`
}
