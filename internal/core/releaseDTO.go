package core

import (
	"github.com/google/uuid"
	"time"
)

type SongDTO struct {
	Id       uuid.UUID     `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	Author   string        `json:"author,omitempty"`
	Date     time.Time     `json:"date,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}
type SongReviewDTO struct {
	SongDTO   `json:"song,omitempty"`
	ReviewDTO `json:"review,omitempty"`
}

type AlbumReviewDTO struct {
	AlbumDTO  `json:"album,omitempty"`
	ReviewDTO `json:"review,omitempty"`
}

type AlbumDTO struct {
	Id         uuid.UUID     `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Author     string        `json:"author,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
	SongAmount int           `json:"song_amount,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Songs      []SongDTO     `json:"songs,omitempty"`
}

type ReleaseDTO struct {
	Id         uuid.UUID     `json:"id,omitempty"`
	IsAlbum    bool          `json:"is_album,omitempty"`
	Name       string        `json:"name,omitempty"`
	Author     string        `json:"author,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	SongAmount int           `json:"song_amount,omitempty"`
	Songs      []SongDTO     `json:"songs,omitempty"`
}
