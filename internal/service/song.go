package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type SongService struct {
	r repository.Song
}

func (s SongService) SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) (res []core.SongWithReviewDTO, err error) {
	songs, err := s.r.SearchSongsWithReview(searchReq, userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.SongWithReviewDTO, len(songs))
	for i, song := range songs {
		sDomain, rDomain := song.SongDAO.ToDomain(), song.ReviewNullableDAO.ToDomain()
		res[i] = core.SongWithReviewDTO{
			SongDTO:   sDomain.ToDTO(),
			ReviewDTO: rDomain.ToEmptyDTO(),
		}
	}
	return res, nil
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
