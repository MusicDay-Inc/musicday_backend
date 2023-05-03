package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type SongService struct {
	r repository.Song
}

func (s SongService) GetById(songId uuid.UUID) (core.Song, error) {
	song, err := s.r.GetById(songId)
	if err != nil {
		return core.Song{}, err
	}
	return song.ToDomain(), nil
}

func NewSongService(songRepo repository.Song) *SongService {
	return &SongService{r: songRepo}
}
