package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type AlbumService struct {
	r repository.Album
}

func (s AlbumService) GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error) {
	songs, err := s.r.GetSongsFromAlbum(id)
	if err != nil {
		return []core.Song{}, err
	}
	res := make([]core.Song, len(songs))
	for i, song := range songs {
		res[i] = song.ToDomain()
	}
	return res, nil
}

func NewAlbumService(r repository.Album) *AlbumService {
	return &AlbumService{r: r}
}

func (s AlbumService) GetById(albumId uuid.UUID) (core.Album, error) {
	album, err := s.r.GetById(albumId)
	if err != nil {
		return core.Album{}, err
	}
	return album.ToDomain(), nil
}
