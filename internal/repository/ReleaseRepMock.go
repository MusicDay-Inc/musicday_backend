package repository

import (
	"errors"
	"server/internal/core"
)

type ReleaseItemRepo struct {
	data []core.Song
}

func NewReleaseItemRepo() *ReleaseItemRepo {
	data := []core.Song{core.NewSong(0, "De Do Do", "The Police"), core.NewSong(1, "Message in a Bottle", "The Police")}
	return &ReleaseItemRepo{data: data}
}

func (rep *ReleaseItemRepo) GetById(releaseId int) (core.Song, error) {
	//err := ((releaseId < 0 || releaseId > len(rep.data)  "error": "ok")
	//var err error
	var item core.Song
	if releaseId < 0 || releaseId >= len(rep.data) {
		err := errors.New("incorrect release id")
		return item, err
	}
	err := error(nil)
	return rep.data[releaseId], err
}

func NewReleaseRepo() *ReleaseItemRepo {
	return &ReleaseItemRepo{}
}
