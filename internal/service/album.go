package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type AlbumService struct {
	r repository.Album
}

func (s AlbumService) GetCoverId(srcId uuid.UUID) (uuid.UUID, error) {
	a, err := s.r.GetById(srcId)
	if err == nil {
		return a.Id, nil
	}
	aId, err := s.r.GetContainingSong(srcId)
	if err != nil {
		return uuid.UUID{}, err
	}
	return aId, nil
}

func (s AlbumService) SearchAlbumsWithReview(query string, userId uuid.UUID, limit int, offset int) (res []core.AlbumWithReview, err error) {
	res, err = s.r.SearchAlbumsWithReview(query, userId, limit, offset)
	if err != nil {
		return
	}
	return res, nil
}

func (s AlbumService) GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error) {
	songs, err := s.r.GetSongsFromAlbum(id)
	if err != nil {
		return []core.Song{}, err
	}
	return songs, nil
}

func NewAlbumService(r repository.Album) *AlbumService {
	return &AlbumService{r: r}
}

func (s AlbumService) GetById(albumId uuid.UUID) (core.Album, error) {
	album, err := s.r.GetById(albumId)
	if err != nil {
		return core.Album{}, err
	}
	return album, nil
}
