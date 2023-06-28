package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type SongService struct {
	r repository.Song
}

func (s SongService) SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) (res []core.SongWithReview, err error) {
	songs, err := s.r.SearchSongsWithReview(searchReq, userId, limit, offset)
	if err != nil {
		return
	}
	return songs, nil
}

func (s SongService) GetById(songId uuid.UUID) (core.Song, error) {
	song, err := s.r.GetById(songId)
	if err != nil {
		return core.Song{}, err
	}
	return song, nil
}

func NewSongService(songRepo repository.Song) *SongService {
	return &SongService{r: songRepo}
}
